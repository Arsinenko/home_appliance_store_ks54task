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
