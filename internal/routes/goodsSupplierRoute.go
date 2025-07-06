package routes

import (
	"HomeApplianceStore/internal/services"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// @Summary      Создать связь между товаром и поставщиком
// @Description  Создаёт новую связь между товаром и поставщиком
// @Tags         goods-supplier
// @Accept       json
// @Produce      json
// @Param        request  body      services.CreateGoodsSupplierDto  true  "Данные для создания связи"
// @Success      201      {object}  services.CreateGoodsSupplierDto
// @Failure      400      {object}  string
// @Failure      500      {object}  string
// @Router       /goods-supplier [post]
func CreateGoodSupplierHandler(serivce services.GoodsSupplierService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dto services.CreateGoodsSupplierDto
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		err := serivce.CreateGoodsSupplier(r.Context(), dto)
		if err != nil {
			if errors.Is(err, services.GoodsSupplierLinkNotFoundError) {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
			}
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

// @Summary      Получить товары по поставщику
// @Description  Возвращает товары, связанные с поставщиком
// @Tags         goods-supplier
// @Produce      json
// @Param        supplier_id  path      int  true  "ID поставщика"
// @Success      200          {array}   services.GoodDto
// @Failure      400          {object}  string
// @Failure      500          {object}  string
// @Router       /goods-supplier/by_supplier_id/{id} [get]
func GetGoodsBySupplierHandler(service services.GoodsSupplierService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		supplierId, err := strconv.Atoi(chi.URLParam(r, "supplier_id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		response, err := service.GetGoodsBySupplier(r.Context(), int32(supplierId))
		if err != nil {
			if errors.Is(err, services.GoodsSupplierLinkNotFoundError) {
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

// @Summary      Получить поставщиков по товару
// @Description  Возвращает поставщиков, связанных с товаром
// @Tags         goods-supplier
// @Produce      json
// @Param        good_id  path      int  true  "ID товара"
// @Success      200      {array}   services.SupplierDto
// @Failure      400      {object}  string
// @Failure      500      {object}  string
// @Router       /goods-supplier/by_good_id/{id} [get]
func GetSuppliersByGoodHandler(service services.GoodsSupplierService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		goodId, err := strconv.Atoi(chi.URLParam(r, "good_id"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		response, err := service.GetSuppliersByGoodId(r.Context(), int32(goodId))
		if err != nil {
			if errors.Is(err, services.GoodsSupplierLinkNotFoundError) {
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

// @Summary      Удалить связь между товаром и поставщиком
// @Description  Удаляет связь между товаром и поставщиком по идентификатору
// @Tags         goods-supplier
// @Produce      json
// @Param        id  path      int  true  "ID связи"
// @Success      204
// @Failure      400      {object}  string
// @Failure      500      {object}  string
// @Router       /goods-supplier/{id} [delete]
func DeleteGoodsSupplierHandler(service services.GoodsSupplierService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}
		err = service.DeleteGoodsSupplier(r.Context(), int32(id))
		if err != nil {
			if errors.Is(err, services.GoodsSupplierLinkNotFoundError) {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusNoContent)
		w.Header().Set("Content-Type", "application/json")
	}
}
func NewGoodsSupplierRouter(service services.GoodsSupplierService) http.Handler {
	r := chi.NewRouter()

	r.Post("/", CreateGoodSupplierHandler(service))
	r.Get("/by_supplier_id/{id}", GetGoodsBySupplierHandler(service))
	r.Get("/by_good_id/{id}", GetGoodsBySupplierHandler(service))
	r.Delete("/", DeleteGoodsSupplierHandler(service))

	return r
}
