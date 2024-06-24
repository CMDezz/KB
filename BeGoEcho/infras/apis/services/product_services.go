package services

import (
	"context"

	"github.com/CMDezz/KB/dto"
	"github.com/CMDezz/KB/infras/logger"
	"github.com/CMDezz/KB/utils"
	"github.com/CMDezz/KB/utils/constants"
)

func (services *Services) CreateProduct(ctx context.Context, req *dto.CreateProductRequest) (*dto.ProductWithVariant, error) {
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

	product, err := services.Queries.DBCreateProduct(ctx, newProduct)
	if err != nil {
		logger.Error("SERVICE - CreateProduct - Error: %v", err)
		return nil, err
	}
	res := dto.ProductWithVariant{
		Product: product,
	}

	var sliceOfVariants []dto.ProductVariant
	if req.Variants != nil {
		for _, v := range *req.Variants {
			newVariant := &dto.ProductVariant{
				Name:             v.Name,
				ImgMain:          v.ImgMain,
				ImgsDetail:       v.ImgsDetail,
				Qty:              v.Qty,
				Price:            v.Price,
				VariantOnProduct: product.Id,
			}
			_, err := services.Queries.DBCreateProductVariant(ctx, newVariant)
			if err != nil {
				logger.Error("SERVICE - CreateProduct - Error: %v", err)
				return nil, err
			}
			sliceOfVariants = append(sliceOfVariants, *newVariant)
		}
	}

	res.Variants = &sliceOfVariants

	return &res, nil
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

func (services *Services) UpdateProductById(ctx context.Context, req *dto.UpdateProductRequest) (*dto.ProductWithVariant, error) {
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

	result, err := services.Queries.DBUpdateProductById(ctx, &newProduct)

	if err != nil {
		logger.Error("SERVICE - GetProductById - Error %v", err)
		return nil, err
	}

	response := dto.ProductWithVariant{
		Product: result,
	}

	var sliceOfVariants []dto.ProductVariant

	oldVariants, err := services.Queries.DBGetVariantsByProductId(ctx, product.Id)
	if err != nil {
		logger.Error("SERVICE - GetProductById - Error %v", err)
		return nil, err
	}

	var tempStoringId []int64

	for _, v := range req.Variants {
		newVariant := &dto.ProductVariant{
			Name:             v.Name,
			ImgMain:          v.ImgMain,
			ImgsDetail:       v.ImgsDetail,
			Qty:              v.Qty,
			Price:            v.Price,
			VariantOnProduct: product.Id,
		}
		if v.Id != nil { //update old variant
			newVariant.Id = *v.Id
			_, err := services.Queries.DBUpdateVariantById(ctx, newVariant)
			if err != nil {
				logger.Error("SERVICE - CreateProduct - Error: %v", err)
				return nil, err
			}
			tempStoringId = append(tempStoringId, *v.Id)
		} else { // create new variant
			_, err := services.Queries.DBCreateProductVariant(ctx, newVariant)
			if err != nil {
				logger.Error("SERVICE - CreateProduct - Error: %v", err)
				return nil, err
			}
		}
		sliceOfVariants = append(sliceOfVariants, *newVariant)
	}

	//clean unupdate variants
	for _, oldV := range *oldVariants {
		if !utils.Includes(tempStoringId, oldV.Id) {
			err := services.Queries.DBDeleteVariantById(ctx, oldV.Id)
			if err != nil {
				logger.Error("SERVICE - CreateProduct - Error: %v", err)
				return nil, err
			}
		}
	}

	response.Variants = &sliceOfVariants

	return &response, nil

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
