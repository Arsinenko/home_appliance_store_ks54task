package services

import (
	"HomeApplianceStore/pkg/gen"
	"context"
	"time"
)

type RoleService struct {
	Queries *gen.Queries
}

type CreateRoleDto struct {
	Name string `json:"name"`
}

func ToRoleDto(role gen.Role) RoleDto {
	return RoleDto{
		Id:        role.ID,
		Name:      role.Name,
		CreatedAt: role.CreatedAt.Time,
	}
}

type RoleDto struct {
	Id        int32     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func (rs *RoleService) CreateRole(ctx context.Context, roleName string) (RoleDto, error) {

	createdRole, err := rs.Queries.CreateRole(ctx, roleName)
	if err != nil {
		return RoleDto{}, err
	}
	response := ToRoleDto(createdRole)
	return response, nil
}

func (rs *RoleService) GetRoles(ctx context.Context) ([]RoleDto, error) {
	roles, err := rs.Queries.GetRoles(ctx)
	if err != nil {
		return nil, err
	}
	response := make([]RoleDto, len(roles))
	for i, r := range roles {
		response[i] = ToRoleDto(r)
	}
	return response, nil
}

func (rs *RoleService) GetRole(ctx context.Context, id int32) (RoleDto, error) {

	role, err := rs.Queries.GetRole(ctx, id)
	if err != nil {
		return RoleDto{}, err
	}
	response := ToRoleDto(role)
	return response, nil
}

func (rs *RoleService) UpdateRole(ctx context.Context, id int32, name string) (RoleDto, error) {

	role, err := rs.Queries.UpdateRole(ctx, gen.UpdateRoleParams{ID: id, Name: name})
	if err != nil {
		return RoleDto{}, err
	}
	response := ToRoleDto(role)
	return response, nil
}
