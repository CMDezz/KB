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

	CreateCategory(ctx context.Context, req *dto.CreateCategoryRequest) (*dto.Category, error)
	GetCategoryById(ctx context.Context, id int64) (*dto.Category, error)
	GetAllCategory(ctx context.Context) (*[]dto.Category, error)
	UpdateCategoryById(ctx context.Context, req *dto.UpdateCategoryRequest) (*dto.Category, error)
	DeleteCategoryById(ctx context.Context, id int64) error

	CreateCollection(ctx context.Context, req *dto.CreateCollectionRequest) (*dto.Collection, error)
	GetCollectionById(ctx context.Context, id int64) (*dto.Collection, error)
	GetAllCollection(ctx context.Context) (*[]dto.Collection, error)
	UpdateCollectionById(ctx context.Context, req *dto.UpdateCollectionRequest) (*dto.Collection, error)
	DeleteCollectionById(ctx context.Context, id int64) error

	CreateProduct(ctx context.Context, req *dto.CreateProductRequest) (*dto.ProductWithVariant, error)
	GetProductById(ctx context.Context, id int64) (*dto.Product, error)
	GetAllProduct(ctx context.Context) (*[]dto.Product, error)
	UpdateProductById(ctx context.Context, req *dto.UpdateProductRequest) (*dto.ProductWithVariant, error)
	DeleteProductById(ctx context.Context, id int64) error
}
