package routes

import (
	"HomeApplianceStore/internal/services"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

// @Summary      Получить список аккаунтов
// @Description  Возвращает все аккаунты
// @Tags         accounts
// @Produce      json
// @Success      200  {array}   services.AccountDto
// @Router       /accounts [get]
func getAccountsHandler(accountService services.AccountService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accounts, err := accountService.GetAccounts(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(accounts)
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
func getAccountByIDHandler(accountService services.AccountService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		account, err := accountService.GetAccount(r.Context(), int32(id))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(account)
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
func createAccountHandler(accountService services.AccountService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestDto services.CreateAccountDto
		err := json.NewDecoder(r.Body).Decode(&requestDto)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		account, err := accountService.CreateAccount(r.Context(), requestDto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(account)
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
func updateAccountHandler(accountService services.AccountService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		var requestDto services.UpdateAccountDto
		err = json.NewDecoder(r.Body).Decode(&requestDto)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		account, err := accountService.UpdateAccount(r.Context(), int32(id), requestDto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(account)
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
func deleteAccountHandler(accountService services.AccountService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		err = accountService.DeleteAccount(r.Context(), int32(id))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusNoContent)
		w.Header().Set("Content-Type", "application/json")
	}
}

func NewAccountRouter(accountService services.AccountService) http.Handler {
	r := chi.NewRouter()

	r.Get("/", getAccountsHandler(accountService))
	r.Get("/{id}", getAccountByIDHandler(accountService))
	r.Post("/", createAccountHandler(accountService))
	r.Put("/{id}", updateAccountHandler(accountService))
	r.Delete("/{id}", deleteAccountHandler(accountService))

	return r
}
