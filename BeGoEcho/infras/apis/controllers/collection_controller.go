package controllers

import (
	"context"
	"errors"

	"github.com/CMDezz/KB/dto"
	"github.com/CMDezz/KB/gerror"
	"github.com/CMDezz/KB/infras/middleware"
	"github.com/CMDezz/KB/infras/response"
	"github.com/CMDezz/KB/utils"
	"github.com/labstack/echo/v4"
)

func (controllers *Controllers) CreateCollection(eCtx echo.Context) error {
	var err error
	ctx := eCtx.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	err = middleware.CheckIsStaff(eCtx)
	if err != nil {
		return controllers.StatusErrorPermission(eCtx, err)
	}
	req := new(dto.CreateCollectionRequest)

	err = eCtx.Bind(&req)

	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorBindData, err.Error(), utils.FuncName())
		return controllers.StatusBadRequest(eCtx, message, resp)
	}

	collection, err := controllers.Services.CreateCollection(ctx, req)

	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorBindData, err.Error(), utils.FuncName())
		return controllers.StatusBadRequest(eCtx, message, resp)
	}

	return controllers.StatusOkResponse(eCtx, &collection)
}

func (controllers *Controllers) GetAllCollection(eCtx echo.Context) error {
	var err error
	ctx := eCtx.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	res, err := controllers.Services.GetAllCollection(ctx)
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorRetriveData, err.Error(), utils.FuncName())
		return controllers.StatusInternalServerError(eCtx, message, resp)
	}
	return controllers.StatusOkResponse(eCtx, &res)
}

func (controllers *Controllers) GetCollectionById(eCtx echo.Context) error {
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

	res, err := controllers.Services.GetCollectionById(ctx, id)

	if err != nil {
		return controllers.StatusError(eCtx, err)
	}

	return controllers.StatusOkResponse(eCtx, &res)
}

func (controllers *Controllers) UpdateCollectionById(eCtx echo.Context) error {
	var err error
	ctx := eCtx.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	err = middleware.CheckIsStaff(eCtx)
	if err != nil {
		return controllers.StatusErrorPermission(eCtx, err)
	}
	req := new(dto.UpdateCollectionRequest)

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

	res, err := controllers.Services.UpdateCollectionById(ctx, req)

	if err != nil {
		return controllers.StatusError(eCtx, err)
	}

	return controllers.StatusOkResponse(eCtx, &res)
}

func (controllers *Controllers) DeleteCollectionById(eCtx echo.Context) error {
	var err error
	ctx := eCtx.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = middleware.CheckIsStaff(eCtx)
	if err != nil {
		return controllers.StatusErrorPermission(eCtx, err)
	}

	id, err := utils.ToInt64(eCtx.Param("id"))
	if err != nil {
		return controllers.StatusErrorBadRequest(eCtx, errors.New("id is not correct format"))
	}

	err = controllers.Services.DeleteCollectionById(ctx, id)

	if err != nil {
		return controllers.StatusError(eCtx, err)
	}

	return controllers.StatusOkResponse(eCtx, nil)
}
