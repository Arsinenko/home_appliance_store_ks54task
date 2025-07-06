package routes

import (
	"HomeApplianceStore/internal/services"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

// @Summary      Создать магазин
// @Description  Создаёт новый магазин
// @Tags         stores
// @Accept       json
// @Produce      json
// @Param        store  body      services.CreateStoreDto  true  "Данные для создания магазина"
// @Success      201    {object}  services.StoreDto
// @Failure      400    {object}  string
// @Router       /stores [post]
func createStoreHandler(service services.StoreService) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var storeDto services.CreateStoreDto
		err := json.NewDecoder(r.Body).Decode(&storeDto)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()
		response, err := service.Create(r.Context(), storeDto)
		if err != nil {
			if errors.Is(err, services.StoreNotFound) {
				w.WriteHeader(http.StatusNotFound)
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
	})
}

// @Summary      Получить магазин по id
// @Description  Возвращает магазин по идентификатору
// @Tags         stores
// @Produce      json
// @Param        id   path      int  true  "ID магазина"
// @Success      200  {object}  services.StoreDto
// @Failure      400  {object}  string
// @Router       /stores/{id} [get]
func GetStoreHandler(service services.StoreService) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		response, err := service.GetStore(r.Context(), int32(id))
		if err != nil {
			if errors.Is(err, services.StoreNotFound) {
				w.WriteHeader(http.StatusNotFound)
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
	})
}

// @Summary      Получить список магазинов
// @Description  Возвращает все магазины
// @Tags         stores
// @Produce      json
// @Success      200  {array}   services.StoreDto
// @Failure      400  {object}  string
// @Router       /stores [get]
func GetStoresHandler(service services.StoreService) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, err := service.GetStores(r.Context())
		if err != nil {
			if errors.Is(err, services.StoreNotFound) {
				w.WriteHeader(http.StatusNotFound)
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
	})
}

// @Summary      Обновить магазин
// @Description  Обновляет данные магазина по id
// @Tags         stores
// @Accept       json
// @Produce      json
// @Param        id      path      int                    true  "ID магазина"
// @Param        store   body      services.UpdateStoreDto true  "Данные для обновления магазина"
// @Success      200     {object}  services.StoreDto
// @Failure      400     {object}  string
// @Router       /stores/{id} [put]
func UpdateStoreHandler(service services.StoreService) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var request services.UpdateStoreDto
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()
		response, err := service.UpdateStore(r.Context(), request)
		if err != nil {
			if errors.Is(err, services.StoreNotFound) {
				w.WriteHeader(http.StatusNotFound)
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
	})
}

// @Summary      Удалить магазин
// @Description  Удаляет магазин по id
// @Tags         stores
// @Produce      json
// @Param        id   path      int  true  "ID магазина"
// @Success      204
// @Failure      400  {object}  string
// @Router       /stores/{id} [delete]
func DeleteStoreHandler(service services.StoreService) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		err = service.DeleteStore(r.Context(), int32(id))
		if err != nil {
			if errors.Is(err, services.StoreNotFound) {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(err.Error()))
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})
}

func NewStoreRouter(service services.StoreService) http.Handler {
	r := chi.NewRouter()
	r.Post("/", createStoreHandler(service))
	r.Get("/{id}", GetStoreHandler(service))
	r.Get("/", GetStoresHandler(service))
	r.Put("/{id}", UpdateStoreHandler(service))
	r.Delete("/{id}", DeleteStoreHandler(service))

	return r
}
