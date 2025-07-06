package services

import (
	"HomeApplianceStore/pkg/gen"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type CreateGoodsSupplierDto struct {
	ProductId  int32 `json:"product_id"`
	SupplierId int32 `json:"supplier_id"`
}

type GoodsSupplierInterface interface {
	CreateGoodsSupplier(ctx context.Context, dto CreateGoodsSupplierDto) error
	GetGoodsBySupplier(ctx context.Context, supplierId int32) ([]GoodDto, error)
	GetSuppliersByGoodId(ctx context.Context, id int32) ([]SupplierDto, error)
	DeleteGoodsSupplier(ctx context.Context, id int32) error
}

type GoodsSupplierService struct {
	Queries gen.Queries
}

func (g GoodsSupplierService) CreateGoodsSupplier(ctx context.Context, dto CreateGoodsSupplierDto) error {
	_, err := g.Queries.CreateGoodsSupplier(ctx, gen.CreateGoodsSupplierParams{
		GoodID:     dto.ProductId,
		SupplierID: dto.SupplierId,
		CreatedAt:  pgtype.Timestamp{Time: time.Now()},
		IsAlive:    true,
	})
	if err != nil {
		return err
	}
	return nil
}

var GoodsSupplierLinkNotFoundError = errors.New("goods supplier not found")

func (g GoodsSupplierService) GetGoodsBySupplier(ctx context.Context, supplierId int32) ([]GoodDto, error) {
	goods, err := g.Queries.ListGoodsBySupplier(ctx, supplierId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, GoodsSupplierLinkNotFoundError
		}
		return nil, err
	}
	response := make([]GoodDto, len(goods))
	for i, good := range goods {
		response[i] = ToProductDto(good)
	}
	return response, nil
}

func (g GoodsSupplierService) GetSuppliersByGoodId(ctx context.Context, id int32) ([]SupplierDto, error) {
	suppliers, err := g.Queries.ListSuppliersByGood(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, GoodsSupplierLinkNotFoundError
		}
		return nil, err
	}
	response := make([]SupplierDto, len(suppliers))
	for i, supplier := range suppliers {
		response[i] = SupplierDto{
			Id: supplier.ID,
			AccountDto: AccountDto{
				Id:        supplier.AccountID,
				Login:     supplier.AccountLogin,
				CreatedAt: supplier.CreatedAt.Time,
				IsAlive:   supplier.AccountIsAlive,
			},
			CreatedAt: supplier.CreatedAt.Time,
			IsAlive:   supplier.IsAlive,
		}
	}
	return response, nil

}

func (g GoodsSupplierService) DeleteGoodsSupplier(ctx context.Context, id int32) error {
	err := g.Queries.DeleteGoodsSupplier(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return GoodsSupplierLinkNotFoundError
		}
		return err
	}
	return nil
}
