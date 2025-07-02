package services

import (
	"HomeApplianceStore/pkg/gen"
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type EmployeeInterface interface {
	CreateEmployee(ctx context.Context, request CreateEmployeeRequest) (EmployeeDto, error)
	GetEmployee(ctx context.Context, id int32) (EmployeeDto, error)
	GetEmployees(ctx context.Context) ([]EmployeeDto, error)
	UpdateEmployee(ctx context.Context, id int32, request UpdateEmployeeDto) (EmployeeDto, error)
	DeleteEmployee(id int32, ctx context.Context) error
}

type EmployeeService struct {
	Queries gen.Queries
}

func (e EmployeeService) CreateEmployee(ctx context.Context, request CreateEmployeeRequest) (EmployeeDto, error) {
	createdEmployee, err := e.Queries.CreateEmployee(ctx, gen.CreateEmployeeParams{
		AccountID: request.AccountId,
		CreatedAt: pgtype.Timestamp{Time: time.Now()},
		RoleID:    request.RoleId,
		IsAlive:   true,
	})
	if err != nil {
		return EmployeeDto{}, err
	}
	employeeRow, err := e.Queries.GetEmployee(ctx, createdEmployee.ID)
	if err != nil {
		return EmployeeDto{}, err
	}
	response := ToEmployeeDtoAny(employeeRow)
	return response, nil
}

func (e EmployeeService) GetEmployee(ctx context.Context, id int32) (EmployeeDto, error) {
	employeeRow, err := e.Queries.GetEmployee(ctx, id)
	if err != nil {
		return EmployeeDto{}, err
	}
	response := ToEmployeeDtoAny(employeeRow)
	return response, nil
}

func (e EmployeeService) GetEmployees(ctx context.Context) ([]EmployeeDto, error) {
	employeesRow, err := e.Queries.ListEmployees(ctx)
	if err != nil {
		return nil, err
	}
	response := make([]EmployeeDto, len(employeesRow))
	for i, row := range employeesRow {
		response[i] = ToEmployeeDtoAny(row)
	}
	return response, nil
}

type UpdateEmployeeDto struct {
	AccountId int32
	RoleId    int32
	IsAlive   bool
}

func (e EmployeeService) UpdateEmployee(ctx context.Context, id int32, request UpdateEmployeeDto) (EmployeeDto, error) {
	_, err := e.Queries.UpdateEmployee(ctx, gen.UpdateEmployeeParams{
		ID:      id,
		RoleID:  request.RoleId,
		IsAlive: request.IsAlive,
	})
	if err != nil {
		return EmployeeDto{}, err
	}
	employeeRow, err := e.Queries.GetEmployee(ctx, id)
	if err != nil {
		return EmployeeDto{}, err
	}
	response := ToEmployeeDtoAny(employeeRow)
	return response, nil

}

func (e EmployeeService) DeleteEmployee(ctx context.Context, id int32) error {
	err := e.Queries.DeleteEmployee(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

type EmployeeDto struct {
	Id        int32      `json:"id"`
	Account   AccountDto `json:"account"`
	Role      RoleDto    `json:"role"`
	CreatedAt time.Time  `json:"created_at"`
	IsAlive   bool       `json:"is_alive"`
}

func ToEmployeeDtoAny(row any) EmployeeDto {
	switch r := row.(type) {
	case gen.GetEmployeeRow:
		return convertRow(
			r.ID,
			r.AccountID,
			r.AccountLogin,
			r.AccountCreatedAt.Time,
			r.AccountIsAlive,
			r.RoleID,
			r.RoleName,
			r.RoleCreatedAt.Time,
			r.CreatedAt.Time,
			r.IsAlive,
		)
	case gen.ListEmployeesRow:
		return convertRow(
			r.ID,
			r.AccountID,
			r.AccountLogin,
			r.AccountCreatedAt.Time,
			r.AccountIsAlive,
			r.RoleID,
			r.RoleName,
			r.RoleCreatedAt.Time,
			r.CreatedAt.Time,
			r.IsAlive,
		)
	default:
		panic("unsupported type")
	}
}

func convertRow(id, accountID int32, accountLogin string, accountCreatedAt time.Time, accountIsAlive bool,
	roleID int32, roleName string, roleCreatedAt time.Time,
	createdAt time.Time, isAlive bool) EmployeeDto {
	return EmployeeDto{
		Id: id,
		Account: AccountDto{
			Id:        accountID,
			Login:     accountLogin,
			CreatedAt: accountCreatedAt,
			IsAlive:   accountIsAlive,
		},
		Role: RoleDto{
			Id:        roleID,
			Name:      roleName,
			CreatedAt: roleCreatedAt,
		},
		CreatedAt: createdAt,
		IsAlive:   isAlive,
	}
}

type CreateEmployeeRequest struct {
	AccountId int32 `json:"account_id"`
	RoleId    int32 `json:"role_id"`
}
