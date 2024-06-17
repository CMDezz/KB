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

func (controllers *Controllers) CreateCategory(eCtx echo.Context) error {
	var err error
	ctx := eCtx.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	err = middleware.CheckIsStaff(eCtx)
	if err != nil {
		return controllers.StatusErrorPermission(eCtx, err)
	}
	req := new(dto.CreateCategoryRequest)

	err = eCtx.Bind(&req)

	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorBindData, err.Error(), utils.FuncName())
		return controllers.StatusBadRequest(eCtx, message, resp)
	}

	category, err := controllers.Services.CreateCategory(ctx, req)

	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorBindData, err.Error(), utils.FuncName())
		return controllers.StatusBadRequest(eCtx, message, resp)
	}

	return controllers.StatusOkResponse(eCtx, &category)
}

func (controllers *Controllers) GetAllCategory(eCtx echo.Context) error {
	var err error
	ctx := eCtx.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	res, err := controllers.Services.GetAllCategory(ctx)
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorRetriveData, err.Error(), utils.FuncName())
		return controllers.StatusInternalServerError(eCtx, message, resp)
	}
	return controllers.StatusOkResponse(eCtx, &res)
}

func (controllers *Controllers) GetCategoryById(eCtx echo.Context) error {
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

	res, err := controllers.Services.GetCategoryById(ctx, id)

	if err != nil {
		return controllers.StatusError(eCtx, err)
	}

	return controllers.StatusOkResponse(eCtx, &res)
}

func (controllers *Controllers) UpdateCategoryById(eCtx echo.Context) error {
	var err error
	ctx := eCtx.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	err = middleware.CheckIsStaff(eCtx)
	if err != nil {
		return controllers.StatusErrorPermission(eCtx, err)
	}
	req := new(dto.UpdateCategoryRequest)

	err = eCtx.Bind(&req)

	if err != nil {
		return controllers.StatusErrorBadRequest(eCtx, err)
	}

	res, err := controllers.Services.UpdateCategoryById(ctx, req)

	if err != nil {
		return controllers.StatusError(eCtx, err)
	}

	return controllers.StatusOkResponse(eCtx, &res)
}

func (controllers *Controllers) DeleteCategoryById(eCtx echo.Context) error {
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

	err = controllers.Services.DeleteCategoryById(ctx, id)

	if err != nil {
		return controllers.StatusError(eCtx, err)
	}

	return controllers.StatusOkResponse(eCtx, nil)
}
