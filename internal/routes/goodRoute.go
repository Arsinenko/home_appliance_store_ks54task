package routes

import (
	"HomeApplianceStore/internal/services"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// @Summary      Создать новый товар
// @Description  Создаёт новый товар
// @Tags         goods
// @Accept       json
// @Produce      json
// @Param        input   body      services.CreateGoodDto  true  "Данные товара"
// @Success      201     {object}  services.GoodDto
// @Router       /goods [post]
func CreateProductHandler(service services.GoodsService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dto services.CreateGoodDto
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		response, err := service.CreateProduct(r.Context(), dto)
		if err != nil {
			if errors.Is(err, services.ProductNotFound) {
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

// @Summary      Получить товар по id
// @Description  Возвращает товар по идентификатору
// @Tags         goods
// @Produce      json
// @Param        id   path      int  true  "ID товара"
// @Success      200  {object}  services.GoodDto
// @Failure      400  {object}  string
// @Router       /goods/{id} [get]
func GetProductHandler(service services.GoodsService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		response, err := service.GetProduct(r.Context(), int32(id))
		if err != nil {
			if errors.Is(err, services.ProductNotFound) {
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

// @Summary      Получить список товаров
// @Description  Возвращает все товары
// @Tags         goods
// @Produce      json
// @Success      200  {array}   services.GoodDto
// @Failure      400  {object}  string
// @Router       /goods [get]
func GetProductsHandler(service services.GoodsService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response, err := service.GetGoods(r.Context())
		if err != nil {
			if errors.Is(err, services.ProductNotFound) {
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

// @Summary      Обновить товар
// @Description  Обновляет данные товара
// @Tags         goods
// @Accept       json
// @Produce      json
// @Param        input   body      services.UpdateGoodDto  true  "Данные для обновления товара"
// @Success      200     {object}  services.GoodDto
// @Failure      400     {object}  string
// @Router       /goods [put]
func UpdateProductHandler(service services.GoodsService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dto services.UpdateGoodDto
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		response, err := service.UpdateGoods(r.Context(), dto)
		if err != nil {
			if errors.Is(err, services.ProductNotFound) {
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

// @Summary      Удалить товар
// @Description  Удаляет товар по идентификатору
// @Tags         goods
// @Produce      json
// @Param        id   path      int  true  "ID товара"
// @Success      204
// @Failure      400  {object}  string
// @Router       /goods/{id} [delete]
func DeleteProductHandler(service services.GoodsService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		err = service.DeleteGood(r.Context(), int32(id))
		if err != nil {
			if errors.Is(err, services.ProductNotFound) {
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

func NewGoodsRouter(service services.GoodsService) http.Handler {
	r := chi.NewRouter()

	r.Post("/", CreateProductHandler(service))
	r.Get("/{id}", GetProductHandler(service))
	r.Get("/", GetProductsHandler(service))
	r.Put("/", UpdateProductHandler(service))
	r.Delete("/{id}", DeleteProductHandler(service))

	return r
}
