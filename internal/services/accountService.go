package services

import (
	"HomeApplianceStore/pkg/gen"
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type AccountDto struct {
	Id        int32     `json:"id"`
	Login     string    `json:"login"`
	CreatedAt time.Time `json:"created_at"`
	IsAlive   bool      `json:"is_alive"`
}

func ToAccountDto(account gen.Account) AccountDto {
	return AccountDto{
		Id:        account.ID,
		Login:     account.Login,
		CreatedAt: account.CreatedAt.Time,
		IsAlive:   account.IsAlive,
	}
}

type AccountInterface interface {
	CreateAccount(ctx context.Context, login string) (AccountDto, error)
	GetAccount(ctx context.Context, id int32) (AccountDto, error)
	GetAccounts(ctx context.Context) ([]AccountDto, error)
	UpdateAccount(id int32, ctx context.Context, request UpdateAccountDto) (AccountDto, error)
	DeleteAccount(id int32, ctx context.Context) error
}

type AccountService struct {
	Queries gen.Queries
}

type CreateAccountDto struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (a AccountService) CreateAccount(ctx context.Context, request CreateAccountDto) (AccountDto, error) {

	account, err := a.Queries.CreateAccount(ctx, gen.CreateAccountParams{
		Login:     request.Login,
		Password:  request.Password,
		CreatedAt: pgtype.Timestamp{Time: time.Now()},
		IsAlive:   true,
	})
	if err != nil {
		return AccountDto{}, err
	}
	response := ToAccountDto(account)
	return response, nil

}

func (a AccountService) GetAccount(ctx context.Context, id int32) (AccountDto, error) {

	account, err := a.Queries.GetAccount(ctx, id)
	if err != nil {
		return AccountDto{}, err
	}
	response := ToAccountDto(account)
	return response, nil
}

func (a AccountService) GetAccounts(ctx context.Context) ([]AccountDto, error) {
	accounts, err := a.Queries.ListAccounts(ctx)
	if err != nil {
		return nil, err
	}
	response := make([]AccountDto, len(accounts))
	for i, account := range accounts {
		response[i] = ToAccountDto(account)
	}
	return response, nil
}

type UpdateAccountDto struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	IsAlive  bool   `json:"is_alive"`
}

func (a AccountService) UpdateAccount(ctx context.Context, id int32, request UpdateAccountDto) (AccountDto, error) {

	account, err := a.Queries.UpdateAccount(ctx, gen.UpdateAccountParams{
		ID:       id,
		Login:    request.Login,
		Password: request.Password,
		IsAlive:  request.IsAlive,
	})
	if err != nil {
		return AccountDto{}, err
	}
	response := ToAccountDto(account)
	return response, nil

}

func (a AccountService) DeleteAccount(ctx context.Context, id int32) error {
	err := a.Queries.DeleteAccount(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
