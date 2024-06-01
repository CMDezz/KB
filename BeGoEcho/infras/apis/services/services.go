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

	// fmt.Println("call to this")

	return discount, nil
}
