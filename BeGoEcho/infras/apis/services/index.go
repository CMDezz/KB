package services

import (
	"context"

	"github.com/CMDezz/KB/dto"
	"github.com/CMDezz/KB/infras/apis/queries"
	"github.com/CMDezz/KB/infras/token"
)

type Services struct {
	Queries queries.IQueries
}

func NewServices(queries queries.IQueries, token *token.JWTTokenMaker) IServices {
	return &Services{
		Queries: queries,
	}
}

type IServices interface {
	CreateDiscount(ctx context.Context, req *dto.CreateDiscountRequest) (*dto.Discount, error)
	GetDiscountById(ctx context.Context, id int64) (*dto.Discount, error)
	GetAllDiscount(ctx context.Context) (*[]dto.Discount, error)
	UpdateDiscountById(ctx context.Context, req *dto.UpdateDiscountRequest) (*dto.Discount, error)
	DeleteDiscountById(ctx context.Context, id int64) error

	GetAllAccount(ctx context.Context) (*[]dto.Account, error)
	CreateAccount(ctx context.Context, req *dto.CreateAccountRequest) (*dto.Account, error)

	LoginAccount(ctx context.Context, req *dto.LoginAccountRequest, tokenMaker token.JWTTokenMaker) (string, error)
}
