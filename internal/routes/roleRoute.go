package routes

import (
	"HomeApplianceStore/internal/services"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

// @Summary      Получить список ролей
// @Description  Возвращает все роли
// @Tags         roles
// @Produce      json
// @Success      200  {array}   services.RoleDto
// @Router       /roles [get]
func getRolesHandler(roleService *services.RoleService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		roles, err := roleService.GetRoles(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error getting roles"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(roles)
	}
}

// @Summary      Получить роль по id
// @Description  Возвращает роль по идентификатору
// @Tags         roles
// @Produce      json
// @Param        id   path      int  true  "ID роли"
// @Success      200  {object}  services.RoleDto
// @Failure      400  {object}  map[string]string
// @Router       /roles/{id} [get]
func getRoleByIDHandler(roleService *services.RoleService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid id"))
			return
		}
		role, err := roleService.GetRole(r.Context(), int32(id))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Role not found"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(role)
	}
}

// @Summary      Создать роль
// @Description  Создаёт новую роль
// @Tags         roles
// @Accept       json
// @Produce      json
// @Param        role  body      services.CreateRoleDto  true  "Данные для создания роли"
// @Success      201   {object}  services.RoleDto
// @Failure      400   {object}  map[string]string
// @Router       /roles [post]
func createRoleHandler(roleService *services.RoleService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req services.CreateRoleDto
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid request"))
			return
		}
		role, err := roleService.CreateRole(r.Context(), req.Name)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error creating role"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(role)
	}
}

// @Summary      Обновить роль
// @Description  Обновляет данные роли по id
// @Tags         roles
// @Accept       json
// @Produce      json
// @Param        id   path      int                        true  "ID роли"
// @Param        role body      services.CreateRoleDto      true  "Данные для обновления роли"
// @Success      200  {object}  services.RoleDto
// @Failure      400  {object}  map[string]string
// @Router       /roles/{id} [put]
func updateRoleHandler(roleService *services.RoleService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid id"))
			return
		}
		var req services.CreateRoleDto
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid request"))
			return
		}
		role, err := roleService.UpdateRole(r.Context(), int32(id), req.Name)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error updating role"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(role)
	}
}

func NewRoleRouter(roleService *services.RoleService) http.Handler {
	r := chi.NewRouter()

	r.Get("/", getRolesHandler(roleService))
	r.Get("/{id}", getRoleByIDHandler(roleService))
	r.Post("/", createRoleHandler(roleService))
	r.Put("/{id}", updateRoleHandler(roleService))

	return r
}
