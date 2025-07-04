package routes

import (
	"HomeApplianceStore/internal/services"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.co
	"net/http"
	"strconv"
)

// Новый тип хендлера, возвращающий ошибку
type HandlerWithError func(w http.ResponseWriter, r *http.Request) error

// Middleware для централизованной обработки ошибок
func ErrorHandler(h HandlerWithError) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) || errors.Is(err, services.AccountNotFoundError) {
				http.Error(w, "not found", http.StatusNotFound)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// Вспомогательная функция для регистрации роутов с ErrorHandler
func RegisterWithErrorHandler(r chi.Router, method, pattern string, handler HandlerWithError) {
	r.MethodFunc(method, pattern, ErrorHandler(handler))
}

// @Summary      Получить список аккаунтов
// @Description  Возвращает все аккаунты
// @Tags         accounts
// @Produce      json
// @Success      200  {array}   services.AccountDto
// @Router       /accounts [get]
func getAccountsHandler(accountService services.AccountService) HandlerWithError {
	return func(w http.ResponseWriter, r *http.Request) error {
		accounts, err := accountService.GetAccounts(r.Context())
		if err != nil {
			return err
		}
		w.Header().Set("Content-Type", "application/json")
		return json.NewEncoder(w).Encode(accounts)
	}
}

// @Summary      Получить аккаунт по id
// @Description  Возвращает аккаунт по идентификатору
// @Tags         accounts
// @Produce      json
// @Param        id   path      int  true  "ID аккаунта"
// @Success      200  {object}  services.AccountDto
// @Failure      400  {object}  map[string]string
// @Router       /accounts/{id} [get]
func getAccountByIDHandler(accountService services.AccountService) HandlerWithError {
	return func(w http.ResponseWriter, r *http.Request) error {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return nil
		}
		account, err := accountService.GetAccount(r.Context(), int32(id))
		if err != nil {
			return err
		}
		w.Header().Set("Content-Type", "application/json")
		return json.NewEncoder(w).Encode(account)
	}
}

// @Summary      Создать аккаунт
// @Description  Создаёт новый аккаунт
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        account  body      services.CreateAccountDto  true  "Данные для создания аккаунта"
// @Success      200      {object}  services.AccountDto
// @Failure      400      {object}  map[string]string
// @Router       /accounts [post]
func createAccountHandler(accountService services.AccountService) HandlerWithError {
	return func(w http.ResponseWriter, r *http.Request) error {
		var requestDto services.CreateAccountDto
		err := json.NewDecoder(r.Body).Decode(&requestDto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return nil
		}
		account, err := accountService.CreateAccount(r.Context(), requestDto)
		if err != nil {
			return err
		}
		w.Header().Set("Content-Type", "application/json")
		return json.NewEncoder(w).Encode(account)
	}
}

// @Summary      Обновить аккаунт
// @Description  Обновляет данные аккаунта по id
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id      path      int                        true  "ID аккаунта"
// @Param        account body      services.UpdateAccountDto  true  "Данные для обновления аккаунта"
// @Success      200     {object}  services.AccountDto
// @Failure      400     {object}  map[string]string
// @Router       /accounts/{id} [put]
func updateAccountHandler(accountService services.AccountService) HandlerWithError {
	return func(w http.ResponseWriter, r *http.Request) error {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return nil
		}
		var requestDto services.UpdateAccountDto
		err = json.NewDecoder(r.Body).Decode(&requestDto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return nil
		}
		defer r.Body.Close()
		account, err := accountService.UpdateAccount(r.Context(), int32(id), requestDto)
		if err != nil {
			return err
		}
		w.Header().Set("Content-Type", "application/json")
		return json.NewEncoder(w).Encode(account)
	}
}

// @Summary      Удалить аккаунт
// @Description  Удаляет аккаунт по id
// @Tags         accounts
// @Produce      json
// @Param        id   path      int  true  "ID аккаунта"
// @Success      204  {object}  nil
// @Failure      400  {object}  map[string]string
// @Router       /accounts/{id} [delete]
func deleteAccountHandler(accountService services.AccountService) HandlerWithError {
	return func(w http.ResponseWriter, r *http.Request) error {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return nil
		}
		err = accountService.DeleteAccount(r.Context(), int32(id))
		if err != nil {
			return err
		}
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}

func NewAccountRouter(accountService services.AccountService) http.Handler {
	r := chi.NewRouter()

	RegisterWithErrorHandler(r, "GET", "/", getAccountsHandler(accountService))
	RegisterWithErrorHandler(r, "GET", "/{id}", getAccountByIDHandler(accountService))
	RegisterWithErrorHandler(r, "POST", "/", createAccountHandler(accountService))
	RegisterWithErrorHandler(r, "PUT", "/{id}", updateAccountHandler(accountService))
	RegisterWithErrorHandler(r, "DELETE", "/{id}", deleteAccountHandler(accountService))

	return r
}
