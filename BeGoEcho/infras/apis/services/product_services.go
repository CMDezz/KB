package services

import (
	"context"

	"github.com/CMDezz/KB/dto"
	"github.com/CMDezz/KB/infras/logger"
	"github.com/CMDezz/KB/utils/constants"
)

func (services *Services) CreateProduct(ctx context.Context, req *dto.CreateProductRequest) (*dto.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.TimeoutRequestDefault)
	defer cancel()

	newProduct := &dto.Product{
		Name:      req.Name,
		ShortDesc: req.ShortDesc,
		Desc:      req.Desc,
		Article:   *req.Article,
	}

	if req.DiscountApplied != nil {
		newProduct.DiscountApplied = req.DiscountApplied
	}

	category, err := services.Queries.DBCreateProduct(ctx, newProduct)

	if err != nil {
		logger.Error("SERVICE - CreateProduct - Error: %v", err)
		return nil, err
	}

	return category, nil
}

func (services *Services) GetAllProduct(ctx context.Context) (*[]dto.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.TimeoutRequestDefault)
	defer cancel()

	res, err := services.Queries.DBGetAllProduct(ctx)

	if err != nil {
		logger.Error("SERVICE - GetAllProduct - Error %v", err)
		return nil, err
	}

	return res, nil
}

func (services *Services) GetProductById(ctx context.Context, id int64) (*dto.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.TimeoutRequestDefault)
	defer cancel()

	res, err := services.Queries.DBGetProductById(ctx, id)

	if err != nil {
		logger.Error("SERVICE - GetProductById - Error %v", err)
		return nil, err
	}
	return res, nil

}

func (services *Services) UpdateProductById(ctx context.Context, req *dto.UpdateProductRequest) (*dto.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.TimeoutRequestDefault)
	defer cancel()

	product, err := services.Queries.DBGetProductById(ctx, req.Id)

	if err != nil {
		logger.Error("SERVICE - GetProductById - Error %v", err)
		return nil, err
	}

	var newProduct dto.Product
	newProduct.Id = product.Id
	newProduct.Article = req.Article
	newProduct.ShortDesc = req.ShortDesc
	newProduct.Desc = req.Desc
	newProduct.Name = req.Name
	newProduct.DiscountApplied = req.DiscountApplied
	newProduct.CreatedAt = product.CreatedAt
	newProduct.IsDeleted = product.IsDeleted

	res, err := services.Queries.DBUpdateProductById(ctx, &newProduct)

	if err != nil {
		logger.Error("SERVICE - GetProductById - Error %v", err)
		return nil, err
	}
	return res, nil

}

func (services *Services) DeleteProductById(ctx context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(ctx, constants.TimeoutRequestDefault)
	defer cancel()

	product, err := services.Queries.DBGetProductById(ctx, id)

	if err != nil {
		logger.Error("SERVICE - GetProductById - Error %v", err)
		return err
	}

	err = services.Queries.DBDeleteProductById(ctx, product.Id)

	if err != nil {
		logger.Error("SERVICE - DeleteProductById - Error %v", err)
		return err
	}
	return nil

}
