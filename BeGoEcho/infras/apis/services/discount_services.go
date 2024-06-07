package services

import (
	"context"
	"time"

	"github.com/CMDezz/KB/dto"
	"github.com/CMDezz/KB/infras/logger"
	"github.com/CMDezz/KB/utils/constants"
)

func (services *Services) CreateDiscount(ctx context.Context, req *dto.CreateDiscountRequest) (*dto.Discount, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.TimeoutRequestDefault)
	defer cancel()

	// discount.BeginAt = time.for

	beginAt, err := time.Parse("2006-01-02T15:04:05.999Z", req.BeginAt)
	if err != nil {
		return nil, err
	}
	expireAt, err := time.Parse("2006-01-02T15:04:05.999Z", req.ExpireAt)
	if err != nil {
		return nil, err
	}

	newDiscount := &dto.Discount{
		Name:      req.Name,
		BeginAt:   beginAt,
		ExpireAt:  expireAt,
		IsDeleted: req.IsDeleted,
	}

	discount, err := services.Queries.DBCreateDiscount(ctx, newDiscount)

	if err != nil {
		logger.Error("SERVICE - CreateDiscount - Error: %v", err)
		return nil, err
	}

	return discount, nil
}

func (services *Services) GetAllDiscount(ctx context.Context) (*[]dto.Discount, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.TimeoutRequestDefault)
	defer cancel()

	res, err := services.Queries.DBGetAllDiscount(ctx)

	if err != nil {
		logger.Error("SERVICE - GetAllDiscount - Error %v", err)
		return nil, err
	}

	return res, nil
}

func (services *Services) GetDiscountById(ctx context.Context, id int64) (*dto.Discount, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.TimeoutRequestDefault)
	defer cancel()

	res, err := services.Queries.DBGetDiscountById(ctx, id)

	if err != nil {
		logger.Error("SERVICE - GetDiscountById - Error %v", err)
		return nil, err
	}
	return res, nil

}

func (services *Services) UpdateDiscountById(ctx context.Context, req *dto.UpdateDiscountRequest) (*dto.Discount, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.TimeoutRequestDefault)
	defer cancel()

	discount, err := services.Queries.DBGetDiscountById(ctx, req.Id)

	if err != nil {
		logger.Error("SERVICE - GetDiscountById - Error %v", err)
		return nil, err
	}

	var newDiscount dto.Discount

	if req.Name != constants.StringEmpty {
		newDiscount.Name = req.Name
	} else {
		newDiscount.Name = discount.Name
	}
	if req.BeginAt != nil {
		newDiscount.BeginAt = *req.BeginAt
	} else {
		newDiscount.BeginAt = discount.BeginAt
	}
	if req.ExpireAt != nil {
		newDiscount.ExpireAt = *req.ExpireAt
	} else {
		newDiscount.ExpireAt = discount.ExpireAt
	}
	newDiscount.Id = discount.Id
	newDiscount.IsDeleted = discount.IsDeleted
	newDiscount.CreatedAt = discount.CreatedAt

	res, err := services.Queries.DBUpdateDiscountById(ctx, &newDiscount)

	if err != nil {
		logger.Error("SERVICE - GetDiscountById - Error %v", err)
		return nil, err
	}
	return res, nil

}

func (services *Services) DeleteDiscountById(ctx context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(ctx, constants.TimeoutRequestDefault)
	defer cancel()

	discount, err := services.Queries.DBGetDiscountById(ctx, id)

	if err != nil {
		logger.Error("SERVICE - GetDiscountById - Error %v", err)
		return err
	}

	err = services.Queries.DBDeleteDiscountById(ctx, discount.Id)

	if err != nil {
		logger.Error("SERVICE - DeleteDiscountById - Error %v", err)
		return err
	}
	return nil

}
