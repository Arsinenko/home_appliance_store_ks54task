package routes

import (
	"HomeApplianceStore/internal/services"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

// @Summary      Создать клиента
// @Description  Создаёт нового клиента
// @Tags         customers
// @Accept       json
// @Produce      json
// @Param        Customer  body      services.CreateCustomerDto  true  "Данные для создания клиента"
// @Success      200      {object}  services.CustomerDto
// @Failure      400      {object}  map[string]string
// @Router       /customers [post]
func createCustomerHandler(service services.CustomerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var createCustomerDto services.CreateCustomerDto
		err := json.NewDecoder(r.Body).Decode(&createCustomerDto)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()
		customer, err := service.CreateCustomer(r.Context(), createCustomerDto)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}
}
func getCustomerHandler(service services.CustomerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		customer, err := service.GetCustomer(r.Context(), int32(id))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)

	}
}
func getCustomersHandler(service services.CustomerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		customers, err := service.GetCustomers(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
func updateCustomerHandler(service services.CustomerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var updateCustomerDto services.UpdateCustomerDto
		err := json.NewDecoder(r.Body).Decode(&updateCustomerDto)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()
		customer, err := service.UpdateCustomer(r.Context(), updateCustomerDto)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}
}
func deleteCustomerHandler(service services.CustomerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		err = service.DeleteCustomer(r.Context(), int32(id))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

func NewCustomerRouter(service services.CustomerService) http.Handler {
	r := chi.NewRouter()

	r.Post("/", createCustomerHandler(service))
	r.Get("/{id}", getCustomerHandler(service))
	r.Get("/", getCustomersHandler(service))
	r.Put("/{id}", updateCustomerHandler(service))
	r.Delete("/{id}", deleteCustomerHandler(service))

	return r

}
