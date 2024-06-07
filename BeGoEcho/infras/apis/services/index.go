package services

import (
	"context"

	"github.com/CMDezz/KB/dto"
	"github.com/CMDezz/KB/infras/apis/queries"
)

type Services struct {
	Queries queries.IQueries
}

func NewServices(queries queries.IQueries) IServices {
	return &Services{
		Queries: queries,
	}
}

type IServices interface {
	CreateDiscount(ctx context.Context, req *dto.CreateDiscountRequest) (*dto.Discount, error)
	GetAllDiscount(ctx context.Context) (*[]dto.Discount, error)
	GetDiscountById(ctx context.Context, id int64) (*dto.Discount, error)
	UpdateDiscountById(ctx context.Context, req *dto.UpdateDiscountRequest) (*dto.Discount, error)
	DeleteDiscountById(ctx context.Context, id int64) error
}
