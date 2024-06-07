package controllers

import (
	"context"
	"errors"

	"github.com/CMDezz/KB/dto"
	"github.com/CMDezz/KB/gerror"
	"github.com/CMDezz/KB/infras/response"
	"github.com/CMDezz/KB/utils"
	"github.com/labstack/echo/v4"
)

func (controllers *Controllers) CreateDiscount(eCtx echo.Context) error {
	var err error
	ctx := eCtx.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	req := new(dto.CreateDiscountRequest)

	err = eCtx.Bind(&req)

	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorBindData, err.Error(), utils.FuncName())
		return controllers.StatusBadRequest(eCtx, message, resp)
	}

	discount, err := controllers.Services.CreateDiscount(ctx, req)

	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorBindData, err.Error(), utils.FuncName())
		return controllers.StatusBadRequest(eCtx, message, resp)
	}

	return controllers.StatusOkResponse(eCtx, &discount)
}

func (controllers *Controllers) GetAllDiscount(eCtx echo.Context) error {
	var err error
	ctx := eCtx.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	res, err := controllers.Services.GetAllDiscount(ctx)
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorRetriveData, err.Error(), utils.FuncName())
		return controllers.StatusInternalServerError(eCtx, message, resp)
	}
	return controllers.StatusOkResponse(eCtx, &res)
}

func (controllers *Controllers) GetDiscountById(eCtx echo.Context) error {
	var err error
	ctx := eCtx.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	id, err := utils.ToInt64(eCtx.Param("id"))
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorValidData, err.Error(), utils.FuncName())
		return controllers.StatusBadRequest(eCtx, message, resp)
	}

	res, err := controllers.Services.GetDiscountById(ctx, id)

	if err != nil {
		return controllers.StatusError(eCtx, err)
	}

	return controllers.StatusOkResponse(eCtx, &res)
}

func (controllers *Controllers) UpdateDiscountById(eCtx echo.Context) error {
	var err error
	ctx := eCtx.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	req := new(dto.UpdateDiscountRequest)

	err = eCtx.Bind(&req)

	// if utils.IsZeroTime(req.BeginAt) {
	// 	return controllers.StatusErrorBadRequest(eCtx, errors.New("beging_at field is required"))
	// }
	// if utils.IsZeroTime(req.ExpireAt) {
	// 	return controllers.StatusErrorBadRequest(eCtx, errors.New("beging_at field is required"))
	// }
	if err != nil {
		return controllers.StatusErrorBadRequest(eCtx, err)
	}

	res, err := controllers.Services.UpdateDiscountById(ctx, req)

	if err != nil {
		return controllers.StatusError(eCtx, err)
	}

	return controllers.StatusOkResponse(eCtx, &res)
}

func (controllers *Controllers) DeleteDiscountById(eCtx echo.Context) error {
	var err error
	ctx := eCtx.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	id, err := utils.ToInt64(eCtx.Param("id"))
	if err != nil {
		return controllers.StatusErrorBadRequest(eCtx, errors.New("id is not correct format"))
	}

	err = controllers.Services.DeleteDiscountById(ctx, id)

	if err != nil {
		return controllers.StatusError(eCtx, err)
	}

	return controllers.StatusOkResponse(eCtx, nil)
}
