package services

import (
	"HomeApplianceStore/pkg/gen"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type StoreInterface interface {
	Create(ctx context.Context, dto CreateStoreDto) (StoreDto, error)
	GetStore(ctx context.Context, id int32) (StoreDto, error)
	GetStores(ctx context.Context) ([]StoreDto, error)
	UpdateStore(ctx context.Context, dto UpdateStoreDto) (StoreDto, error)
	DeleteStore(ctx context.Context, id int32) error
}
type StoreService struct {
	Queries gen.Queries
}
type StoreDto struct {
	Id        int32     `json:"id"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsAlive   bool      `json:"is_alive"`
}

type CreateStoreDto struct {
	Address string `json:"address"`
}

type UpdateStoreDto struct {
	Id      int32  `json:"id"`
	Address string `json:"address"`
	IsAlive bool   `json:"is_alive"`
}

func ToStoreDto(store gen.Store) StoreDto {
	response := StoreDto{
		Id:        store.ID,
		Address:   store.Address,
		CreatedAt: store.CreatedAt.Time,
		UpdatedAt: store.UpdatedAt.Time,
		IsAlive:   store.IsAlive,
	}
	return response
}

var StoreNotFound = errors.New("Store not found")

func (s StoreService) Create(ctx context.Context, dto CreateStoreDto) (StoreDto, error) {
	store, err := s.Queries.CreateStore(ctx, gen.CreateStoreParams{
		Address:   dto.Address,
		CreatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
		IsAlive:   true,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return StoreDto{}, StoreNotFound
		}
		return StoreDto{}, err
	}
	response := ToStoreDto(store)
	return response, nil

}

func (s StoreService) GetStore(ctx context.Context, id int32) (StoreDto, error) {
	store, err := s.Queries.GetStore(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return StoreDto{}, StoreNotFound
		}
		return StoreDto{}, err
	}
	response := ToStoreDto(store)
	return response, nil
}

func (s StoreService) GetStores(ctx context.Context) ([]StoreDto, error) {
	stores, err := s.Queries.GetStores(ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []StoreDto{}, StoreNotFound
		}
		return []StoreDto{}, err
	}
	response := make([]StoreDto, len(stores))
	for i, store := range stores {
		response[i] = ToStoreDto(store)
	}
	return response, nil
}

func (s StoreService) UpdateStore(ctx context.Context, dto UpdateStoreDto) (StoreDto, error) {
	store, err := s.Queries.UpdateStore(ctx, gen.UpdateStoreParams{
		ID:      dto.Id,
		Address: dto.Address,
		IsAlive: dto.IsAlive,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return StoreDto{}, StoreNotFound
		}
		return StoreDto{}, err
	}
	response := ToStoreDto(store)
	return response, nil
}

func (s StoreService) DeleteStore(ctx context.Context, id int32) error {
	err := s.Queries.DeleteStore(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return StoreNotFound
		}
		return err
	}
	return nil
}
