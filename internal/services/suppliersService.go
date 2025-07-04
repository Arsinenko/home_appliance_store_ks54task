package services

import (
	"HomeApplianceStore/pkg/gen"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type CreateSupplierDto struct {
	AccountId int32 `json:"account_id"`
}

type SupplierDto struct {
	Id         int32      `json:"id"`
	AccountDto AccountDto `json:"account"`
	CreatedAt  time.Time  `json:"created_at"`
	IsAlive    bool       `json:"is_alive"`
}

type UpdateSupplierDto struct {
	Id      int32 `json:"id"`
	IsAlive bool  `json:"is_alive"`
}

type SupplierInterface interface {
	CreateSupplier(ctx context.Context, dto CreateSupplierDto) (SupplierDto, error)
	GetSupplier(ctx context.Context, id int32) (SupplierDto, error)
	GetSuppliers(ctx context.Context) ([]SupplierDto, error)
	UpdateSupplier(ctx context.Context, dto UpdateSupplierDto) (SupplierDto, error)
	DeleteSupplier(ctx context.Context, id int32) error
}

type SupplierService struct {
	Queries gen.Queries
}

var SupplierNotFoundError = errors.New("supplier not found")

func (s SupplierService) CreateSupplier(ctx context.Context, dto CreateSupplierDto) (SupplierDto, error) {
	supplier, err := s.Queries.CreateSupplier(ctx, gen.CreateSupplierParams{
		AccountID: dto.AccountId,
		CreatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
		IsAlive:   true,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return SupplierDto{}, SupplierNotFoundError
		}
		return SupplierDto{}, err
	}
	account, err := s.Queries.GetAccount(ctx, supplier.AccountID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return SupplierDto{}, AccountNotFoundError
		}
		return SupplierDto{}, err
	}
	response := SupplierDto{
		Id: supplier.ID,
		AccountDto: AccountDto{
			Id:        supplier.AccountID,
			Login:     account.Login,
			CreatedAt: account.CreatedAt.Time,
			IsAlive:   account.IsAlive,
		},
		CreatedAt: account.CreatedAt.Time,
		IsAlive:   account.IsAlive,
	}
	return response, nil
}

func (s SupplierService) GetSupplier(ctx context.Context, id int32) (SupplierDto, error) {
	supplier, err := s.Queries.GetSupplier(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return SupplierDto{}, SupplierNotFoundError
		}
		return SupplierDto{}, err
	}
	response := SupplierDto{
		Id: supplier.ID,
		AccountDto: AccountDto{
			Id:        supplier.AccountID,
			Login:     supplier.AccountLogin,
			CreatedAt: supplier.AccountCreatedAt.Time,
			IsAlive:   supplier.AccountIsAlive,
		},
		CreatedAt: supplier.CreatedAt.Time,
		IsAlive:   supplier.IsAlive,
	}
	return response, nil
}

func (s SupplierService) GetSuppliers(ctx context.Context) ([]SupplierDto, error) {
	suppliers, err := s.Queries.ListSuppliers(ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []SupplierDto{}, SupplierNotFoundError
		}
		return []SupplierDto{}, err
	}
	response := make([]SupplierDto, len(suppliers))
	for i, supplier := range suppliers {
		response[i] = SupplierDto{
			Id: supplier.ID,
			AccountDto: AccountDto{
				Id:        supplier.AccountID,
				Login:     supplier.AccountLogin,
				CreatedAt: supplier.AccountCreatedAt.Time,
				IsAlive:   supplier.AccountIsAlive,
			},
			CreatedAt: supplier.CreatedAt.Time,
			IsAlive:   supplier.IsAlive,
		}
	}
	return response, nil
}

func (s SupplierService) UpdateSupplier(ctx context.Context, dto UpdateSupplierDto) (SupplierDto, error) {
	supplier, err := s.Queries.UpdateSupplier(ctx, gen.UpdateSupplierParams{
		ID:      dto.Id,
		IsAlive: dto.IsAlive,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return SupplierDto{}, SupplierNotFoundError
		}
		return SupplierDto{}, err
	}
	account, err := s.Queries.GetAccount(ctx, supplier.AccountID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return SupplierDto{}, AccountNotFoundError
		}
		return SupplierDto{}, err
	}

	response := SupplierDto{
		Id: supplier.ID,
		AccountDto: AccountDto{
			Id:        supplier.AccountID,
			Login:     account.Login,
			CreatedAt: account.CreatedAt.Time,
			IsAlive:   account.IsAlive,
		},
		CreatedAt: account.CreatedAt.Time,
		IsAlive:   account.IsAlive,
	}
	return response, nil
}

func (s SupplierService) DeleteSupplier(ctx context.Context, id int32) error {
	err := s.Queries.DeleteSupplier(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return SupplierNotFoundError
		}
		return err
	}
	return nil
}
