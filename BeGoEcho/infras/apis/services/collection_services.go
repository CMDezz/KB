package services

import (
	"context"

	"github.com/CMDezz/KB/dto"
	"github.com/CMDezz/KB/infras/logger"
	"github.com/CMDezz/KB/utils/constants"
)

func (services *Services) CreateCollection(ctx context.Context, req *dto.CreateCollectionRequest) (*dto.Collection, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.TimeoutRequestDefault)
	defer cancel()

	newCollection := &dto.Collection{
		Name:      req.Name,
		ShortDesc: req.ShortDesc,
		Desc:      req.Desc,
		Article:   *req.Article,
	}

	category, err := services.Queries.DBCreateCollection(ctx, newCollection)

	if err != nil {
		logger.Error("SERVICE - CreateCollection - Error: %v", err)
		return nil, err
	}

	return category, nil
}

func (services *Services) GetAllCollection(ctx context.Context) (*[]dto.Collection, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.TimeoutRequestDefault)
	defer cancel()

	res, err := services.Queries.DBGetAllCollection(ctx)

	if err != nil {
		logger.Error("SERVICE - GetAllCollection - Error %v", err)
		return nil, err
	}

	return res, nil
}

func (services *Services) GetCollectionById(ctx context.Context, id int64) (*dto.Collection, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.TimeoutRequestDefault)
	defer cancel()

	res, err := services.Queries.DBGetCollectionById(ctx, id)

	if err != nil {
		logger.Error("SERVICE - GetCollectionById - Error %v", err)
		return nil, err
	}
	return res, nil

}

func (services *Services) UpdateCollectionById(ctx context.Context, req *dto.UpdateCollectionRequest) (*dto.Collection, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.TimeoutRequestDefault)
	defer cancel()

	collection, err := services.Queries.DBGetCollectionById(ctx, req.Id)

	if err != nil {
		logger.Error("SERVICE - GetCollectionById - Error %v", err)
		return nil, err
	}

	var newCollection dto.Collection
	newCollection.Id = collection.Id
	newCollection.Article = req.Article
	newCollection.ShortDesc = req.ShortDesc
	newCollection.Desc = req.Desc
	newCollection.Name = req.Name
	newCollection.CreatedAt = collection.CreatedAt
	newCollection.IsDeleted = collection.IsDeleted

	res, err := services.Queries.DBUpdateCollectionById(ctx, &newCollection)

	if err != nil {
		logger.Error("SERVICE - GetCollectionById - Error %v", err)
		return nil, err
	}
	return res, nil

}

func (services *Services) DeleteCollectionById(ctx context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(ctx, constants.TimeoutRequestDefault)
	defer cancel()

	collection, err := services.Queries.DBGetCollectionById(ctx, id)

	if err != nil {
		logger.Error("SERVICE - GetCollectionById - Error %v", err)
		return err
	}

	err = services.Queries.DBDeleteCollectionById(ctx, collection.Id)

	if err != nil {
		logger.Error("SERVICE - DeleteCollectionById - Error %v", err)
		return err
	}
	return nil

}
