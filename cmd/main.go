package main

import (
	"HomeApplianceStore/internal/routes"
	"HomeApplianceStore/internal/services"
	"HomeApplianceStore/pkg"
	"HomeApplianceStore/pkg/gen"
	"fmt"
	"log"
	"net/http"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// @title           Home Appliance Store API
// @version         1.0
// @description     API для управления магазином бытовой техники.
// @termsOfService  http://swagger.io/terms/
//
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
//
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
//
// @host      localhost:8080
// @BasePath  /
//
// @securityDefinitions.basic  BasicAuth

func main() {
	db := pkg.DatabaseInit()
	queries := gen.New(db)

	accountService := services.AccountService{Queries: *queries}
	employeeService := services.EmployeeService{Queries: *queries}
	roleService := &services.RoleService{Queries: queries}
	customerService := services.CustomerService{Queries: *queries}
	goodsService := services.GoodsService{Queries: *queries}
	storeService := services.StoreService{Queries: *queries}
	supplierService := services.SupplierService{Queries: *queries}
	goodsSupplierService := services.GoodsSupplierService{Queries: *queries}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/reference", func(w http.ResponseWriter, r *http.Request) {
		htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
			// SpecURL: "https://generator3.swagger.io/openapi.json",// allow external URL or local path file
			SpecURL: "./docs/swagger.json",
			CustomOptions: scalar.CustomOptions{
				PageTitle: "Simple API",
			},
			DarkMode: true,
		})

		if err != nil {
			fmt.Printf("%v", err)
		}

		fmt.Fprintln(w, htmlContent)
	})

	r.Mount("/accounts", routes.NewAccountRouter(accountService))
	r.Mount("/employees", routes.NewEmployeeRouter(employeeService))
	r.Mount("/roles", routes.NewRoleRouter(roleService))
	r.Mount("/customers", routes.NewCustomerRouter(customerService))
	r.Mount("/goods", routes.NewGoodsRouter(goodsService))
	r.Mount("/stores", routes.NewStoreRouter(storeService))
	r.Mount("/suppliers", routes.NewSupplierRouter(supplierService))
	r.Mount("/goods-suppliers", routes.NewGoodsSupplierRouter(goodsSupplierService))

	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", r)
}
