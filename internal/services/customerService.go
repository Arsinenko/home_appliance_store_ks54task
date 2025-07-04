package services

import (
	"HomeApplianceStore/pkg/gen"
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	"math/big"
	"time"
)

type CustomerInt interface {
	CreateCustomer(ctx context.Context, request CreateCustomerDto) (CustomerDto, error)
	GetCustomer(ctx context.Context, id int32) (CustomerDto, error)
	GetCustomers(ctx context.Context) ([]CustomerDto, error)
	UpdateCustomer(ctx context.Context, request UpdateCustomerDto) (CustomerDto, error)
	DeleteCustomer(ctx context.Context, id int32) error
}

type CustomerService struct {
	Queries gen.Queries
}

func (c CustomerService) CreateCustomer(ctx context.Context, request CreateCustomerDto) (CustomerDto, error) {
	customer, err := c.Queries.CreateCustomer(ctx, gen.CreateCustomerParams{
		AccountID: request.AccountId,
		Balance:   pgtype.Numeric{Int: big.NewInt(int64(request.Balance))},
		CreatedAt: pgtype.Timestamp{Time: time.Now()},
		IsAlive:   true,
	})
	if err != nil {
		return CustomerDto{}, err
	}
	account, err := c.Queries.GetAccount(ctx, request.AccountId)
	if err != nil {
		return CustomerDto{}, err
	}
	accountDto := ToAccountDto(account)
	response := CustomerDto{
		Id:        customer.ID,
		Account:   accountDto,
		Balance:   customer.Balance.Int.Int64(),
		CreatedAt: customer.CreatedAt.Time,
		IsAlive:   account.IsAlive,
	}
	return response, nil
}

func (c CustomerService) GetCustomer(ctx context.Context, id int32) (CustomerDto, error) {
	customer, err := c.Queries.GetCustomer(ctx, id)
	if err != nil {
		return CustomerDto{}, err
	}
	response := CustomerDto{
		Id: customer.ID,
		Account: AccountDto{
			Id:        customer.AccountID,
			Login:     customer.AccountLogin,
			CreatedAt: customer.AccountCreatedAt.Time,
			IsAlive:   customer.AccountIsAlive,
		},
		Balance:   customer.Balance.Int.Int64(),
		CreatedAt: customer.CreatedAt.Time,
		IsAlive:   customer.IsAlive,
	}
	return response, nil
}

func (c CustomerService) GetCustomers(ctx context.Context) ([]CustomerDto, error) {
	customers, err := c.Queries.ListCustomers(ctx)
	if err != nil {
		return nil, err
	}
	response := make([]CustomerDto, len(customers))
	for i, customer := range customers {
		response[i] = CustomerDto{
			Id: customer.ID,
			Account: AccountDto{
				Id:        customer.AccountID,
				Login:     customer.AccountLogin,
				CreatedAt: customer.CreatedAt.Time,
				IsAlive:   customer.IsAlive,
			},
			Balance:   customer.Balance.Int.Int64(),
			CreatedAt: customer.CreatedAt.Time,
			IsAlive:   customer.IsAlive,
		}
	}
	return response, nil
}

func (c CustomerService) UpdateCustomer(ctx context.Context, request UpdateCustomerDto) (CustomerDto, error) {
	customer, err := c.Queries.UpdateCustomer(ctx, gen.UpdateCustomerParams{
		ID:      request.Id,
		Balance: pgtype.Numeric{Int: request.Balance},
		IsAlive: request.IsAlive,
	})
	if err != nil {
		return CustomerDto{}, err
	}
	account, err := c.Queries.GetAccount(ctx, request.AccountId)
	if err != nil {
		return CustomerDto{}, err
	}
	accountDto := ToAccountDto(account)
	response := CustomerDto{
		Id:        customer.ID,
		Account:   accountDto,
		Balance:   customer.Balance.Int.Int64(),
		CreatedAt: customer.CreatedAt.Time,
	}
	return response, nil
}

func (c CustomerService) DeleteCustomer(ctx context.Context, id int32) error {
	err := c.Queries.DeleteCustomer(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

type CustomerDto struct {
	Id        int32      `json:"id"`
	Account   AccountDto `json:"account"`
	Balance   int64      `json:"balance"`
	CreatedAt time.Time  `json:"created_at"`
	IsAlive   bool       `json:"is_alive"`
}

type CreateCustomerDto struct {
	AccountId int32
	Balance   int
}

type UpdateCustomerDto struct {
	Id        int32 `json:"id"`
	AccountId int32
	Balance   *big.Int
	IsAlive   bool
}
