package services

import (
	"HomeApplianceStore/pkg/gen"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"math/big"
)

type GoodDto struct {
	Id       int32  `json:"id"`
	Article  string `json:"article"`
	Price    int64  `json:"price"`
	Name     string `json:"name"`
	Quantity int32  `json:"quantity"`
	IsAlive  bool   `json:"is_alive"`
}

type CreateGoodDto struct {
	Article  string `json:"article"`
	Price    int32  `json:"price"`
	Name     string `json:"name"`
	Quantity int32  `json:"quantity"`
}

type UpdateGoodDto struct {
	Id       int32  `json:"id"`
	Article  string `json:"article"`
	Price    int64  `json:"price"`
	Name     string `json:"name"`
	Quantity int32  `json:"quantity"`
	IsAlive  bool   `json:"is_alive"`
}

type GoodsInterface interface {
	CreateProduct(ctx context.Context, dto CreateGoodDto) (GoodDto, error)
	GetProduct(ctx context.Context, id int32) (GoodDto, error)
	GetGoods(ctx context.Context) ([]GoodDto, error)
	UpdateGoods(ctx context.Context, dto UpdateGoodDto) (GoodDto, error)
	DeleteGood(ctx context.Context, id int32) error
}

type GoodsService struct {
	Queries gen.Queries
}

var ProductNotFound = errors.New("Product not found")

func (g GoodsService) CreateProduct(ctx context.Context, dto CreateGoodDto) (GoodDto, error) {
	product, err := g.Queries.CreateGood(ctx, gen.CreateGoodParams{
		Article:  dto.Article,
		Price:    pgtype.Numeric{Int: big.NewInt(int64(dto.Price))},
		Name:     dto.Name,
		Quantity: dto.Quantity,
		IsAlive:  true,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return GoodDto{}, ProductNotFound
		}
		return GoodDto{}, err
	}
	response := ToProductDto(product)
	return response, nil
}

func ToProductDto(product gen.Good) GoodDto {
	response := GoodDto{
		Id:       product.ID,
		Article:  product.Article,
		Price:    product.Price.Int.Int64(),
		Name:     product.Name,
		Quantity: product.Quantity,
		IsAlive:  product.IsAlive,
	}
	return response
}

func (g GoodsService) GetProduct(ctx context.Context, id int32) (GoodDto, error) {
	product, err := g.Queries.GetGood(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return GoodDto{}, ProductNotFound
		}
		return GoodDto{}, err
	}
	response := ToProductDto(product)
	return response, nil
}

func (g GoodsService) GetGoods(ctx context.Context) ([]GoodDto, error) {
	products, err := g.Queries.ListGoods(ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ProductNotFound
		}
		return nil, err
	}
	response := make([]GoodDto, len(products))
	for i, product := range products {
		response[i] = ToProductDto(product)
	}
	return response, nil
}

func (g GoodsService) UpdateGoods(ctx context.Context, dto UpdateGoodDto) (GoodDto, error) {
	product, err := g.Queries.UpdateGood(ctx, gen.UpdateGoodParams{
		ID:       dto.Id,
		Article:  dto.Article,
		Price:    pgtype.Numeric{Int: big.NewInt(dto.Price)},
		Name:     dto.Name,
		Quantity: dto.Quantity,
		IsAlive:  dto.IsAlive,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return GoodDto{}, ProductNotFound
		}
		return GoodDto{}, err
	}
	response := ToProductDto(product)
	return response, nil
}

func (g GoodsService) DeleteGood(ctx context.Context, id int32) error {
	err := g.Queries.DeleteGood(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ProductNotFound
		}
		return err
	}
	return nil
}
