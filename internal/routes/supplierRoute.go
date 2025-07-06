package routes

import (
	"HomeApplianceStore/internal/services"
	"context"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

// @Summary      Создать поставщика
// @Description  Создаёт нового поставщика
// @Tags         suppliers
// @Accept       json
// @Produce      json
// @Param        supplier  body      services.CreateSupplierDto  true  "Данные для создания поставщика"
// @Success      201      {object}  services.SupplierDto
// @Failure      400      {object}  string
// @Router       /suppliers [post]
func createSupplierHandler(service services.SupplierService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dto services.CreateSupplierDto
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		response, err := service.CreateSupplier(context.Background(), dto)
		if err != nil {
			if errors.Is(err, services.SupplierNotFoundError) {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

// @Summary      Получить поставщика по id
// @Description  Возвращает поставщика по идентификатору
// @Tags         suppliers
// @Produce      json
// @Param        id   path      int  true  "ID поставщика"
// @Success      200  {object}  services.SupplierDto
// @Failure      400  {object}  string
// @Router       /suppliers/{id} [get]
func getSupplierHandler(service services.SupplierService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		response, err := service.GetSupplier(context.Background(), int32(id))
		if err != nil {
			if errors.Is(err, services.SupplierNotFoundError) {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

// @Summary      Получить список поставщиков
// @Description  Возвращает всех поставщиков
// @Tags         suppliers
// @Produce      json
// @Success      200  {array}   services.SupplierDto
// @Failure      400  {object}  string
// @Router       /suppliers [get]
func GetSuppliersHandler(service services.SupplierService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response, err := service.GetSuppliers(context.Background())
		if err != nil {
			if errors.Is(err, services.SupplierNotFoundError) {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

// @Summary      Обновить поставщика
// @Description  Обновляет данные поставщика по id
// @Tags         suppliers
// @Accept       json
// @Produce      json
// @Param        id      path      int                        true  "ID поставщика"
// @Param        supplier body      services.UpdateSupplierDto  true  "Данные для обновления поставщика"
// @Success      200     {object}  services.SupplierDto
// @Failure      400     {object}  string
// @Router       /suppliers/{id} [put]
func UpdateSupplierHandler(service services.SupplierService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dto services.UpdateSupplierDto
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		response, err := service.UpdateSupplier(context.Background(), dto)
		if err != nil {
			if errors.Is(err, services.SupplierNotFoundError) {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

// @Summary      Удалить поставщика
// @Description  Удаляет поставщика по id
// @Tags         suppliers
// @Produce      json
// @Param        id   path      int  true  "ID поставщика"
// @Success      204
// @Failure      400  {object}  string
// @Router       /suppliers/{id} [delete]
func DeleteSupplierHandler(service services.SupplierService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		err = service.DeleteSupplier(context.Background(), int32(id))
		if err != nil {
			if errors.Is(err, services.SupplierNotFoundError) {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

func NewSupplierRouter(service services.SupplierService) http.Handler {
	router := chi.NewRouter()

	router.Get("/", GetSuppliersHandler(service))
	router.Post("/", createSupplierHandler(service))
	router.Get("/{id}", getSupplierHandler(service))
	router.Put(`/{id}`, UpdateSupplierHandler(service))
	router.Delete("/{id}", DeleteSupplierHandler(service))

	return router
}
