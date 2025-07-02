package routes

import (
	"HomeApplianceStore/internal/services"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

// @Summary      Получить список сотрудников
// @Description  Возвращает всех сотрудников
// @Tags         employees
// @Produce      json
// @Success      200  {array}   services.EmployeeDto
// @Router       /employees [get]
func getEmployeesHandler(service services.EmployeeService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		employees, err := service.GetEmployees(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(employees)
	}
}

// @Summary      Получить сотрудника по id
// @Description  Возвращает сотрудника по идентификатору
// @Tags         employees
// @Produce      json
// @Param        id   path      int  true  "ID сотрудника"
// @Success      200  {object}  services.EmployeeDto
// @Failure      400  {object}  map[string]string
// @Router       /employees/{id} [get]
func getEmployeeByIDHandler(service services.EmployeeService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 32)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		employee, err := service.GetEmployee(r.Context(), int32(id))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(employee)
	}
}

// @Summary      Создать сотрудника
// @Description  Создаёт нового сотрудника
// @Tags         employees
// @Accept       json
// @Produce      json
// @Param        employee  body      services.CreateEmployeeRequest  true  "Данные для создания сотрудника"
// @Success      200       {object}  services.EmployeeDto
// @Failure      400       {object}  map[string]string
// @Router       /employees [post]
func createEmployeeHandler(service services.EmployeeService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request services.CreateEmployeeRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		employee, err := service.CreateEmployee(r.Context(), request)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(employee)
	}
}

// @Summary      Обновить сотрудника
// @Description  Обновляет данные сотрудника по id
// @Tags         employees
// @Accept       json
// @Produce      json
// @Param        id       path      int                              true  "ID сотрудника"
// @Param        employee body      services.UpdateEmployeeDto       true  "Данные для обновления сотрудника"
// @Success      200      {object}  services.EmployeeDto
// @Failure      400      {object}  map[string]string
// @Router       /employees/{id} [put]
func updateEmployeeHandler(service services.EmployeeService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 32)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		var request services.UpdateEmployeeDto
		err = json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		employee, err := service.UpdateEmployee(r.Context(), int32(id), request)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(employee)
	}
}

// @Summary      Удалить сотрудника
// @Description  Удаляет сотрудника по id
// @Tags         employees
// @Produce      json
// @Param        id   path      int  true  "ID сотрудника"
// @Success      204  {object}  nil
// @Failure      400  {object}  map[string]string
// @Router       /employees/{id} [delete]
func deleteEmployeeHandler(service services.EmployeeService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 32)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		err = service.DeleteEmployee(r.Context(), int32(id))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

func NewEmployeeRouter(service services.EmployeeService) http.Handler {
	r := chi.NewRouter()

	r.Get("/", getEmployeesHandler(service))
	r.Get("/{id}", getEmployeeByIDHandler(service))
	r.Post("/", createEmployeeHandler(service))
	r.Put("/{id}", updateEmployeeHandler(service))
	r.Delete("/{id}", deleteEmployeeHandler(service))

	return r
}
