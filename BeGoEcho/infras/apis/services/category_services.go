package services

import (
	"context"

	"github.com/CMDezz/KB/dto"
	"github.com/CMDezz/KB/infras/logger"
	"github.com/CMDezz/KB/utils/constants"
)

func (services *Services) CreateCategory(ctx context.Context, req *dto.CreateCategoryRequest) (*dto.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.TimeoutRequestDefault)
	defer cancel()

	newCategory := &dto.Category{
		Name: req.Name,
	}

	if req.Parent != nil {
		newCategory.Parent = req.Parent
	}

	category, err := services.Queries.DBCreateCategory(ctx, newCategory)

	if err != nil {
		logger.Error("SERVICE - CreateCategory - Error: %v", err)
		return nil, err
	}

	return category, nil
}

func (services *Services) GetAllCategory(ctx context.Context) (*[]dto.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.TimeoutRequestDefault)
	defer cancel()

	res, err := services.Queries.DBGetAllCategory(ctx)

	if err != nil {
		logger.Error("SERVICE - GetAllCategory - Error %v", err)
		return nil, err
	}

	return res, nil
}

func (services *Services) GetCategoryById(ctx context.Context, id int64) (*dto.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.TimeoutRequestDefault)
	defer cancel()

	res, err := services.Queries.DBGetCategoryById(ctx, id)

	if err != nil {
		logger.Error("SERVICE - GetCategoryById - Error %v", err)
		return nil, err
	}
	return res, nil

}

func (services *Services) UpdateCategoryById(ctx context.Context, req *dto.UpdateCategoryRequest) (*dto.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.TimeoutRequestDefault)
	defer cancel()

	category, err := services.Queries.DBGetCategoryById(ctx, req.Id)

	if err != nil {
		logger.Error("SERVICE - GetCategoryById - Error %v", err)
		return nil, err
	}

	newCategory := dto.Category{
		Name:      req.Name,
		IsDeleted: category.IsDeleted,
	}

	if req.Parent != nil {
		parentCategory, err := services.Queries.DBGetCategoryById(ctx, *req.Parent)
		if err != nil {
			logger.Error("SERVICE - GetCategoryById - Error %v", err)
			return nil, err
		}

		newCategory.Parent = &parentCategory.Id

	}
	if req.IsDeleted != nil {
		newCategory.IsDeleted = *req.IsDeleted
	}

	res, err := services.Queries.DBUpdateCategoryById(ctx, &newCategory)

	if err != nil {
		logger.Error("SERVICE - GetCategoryById - Error %v", err)
		return nil, err
	}
	return res, nil

}

func (services *Services) DeleteCategoryById(ctx context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(ctx, constants.TimeoutRequestDefault)
	defer cancel()

	category, err := services.Queries.DBGetCategoryById(ctx, id)

	if err != nil {
		logger.Error("SERVICE - GetCategoryById - Error %v", err)
		return err
	}

	err = services.Queries.DBDeleteCategoryById(ctx, category.Id)

	if err != nil {
		logger.Error("SERVICE - DeleteCategoryById - Error %v", err)
		return err
	}
	return nil

}
