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
	CreateDiscount(ctx context.Context, discount *dto.CreateDiscountRequest) (*dto.Discount, error)
}
