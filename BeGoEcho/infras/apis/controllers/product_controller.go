package controllers

import (
	"context"
	"errors"
	"fmt"

	"github.com/CMDezz/KB/dto"
	"github.com/CMDezz/KB/gerror"
	"github.com/CMDezz/KB/infras/middleware"
	"github.com/CMDezz/KB/infras/response"
	"github.com/CMDezz/KB/utils"
	"github.com/labstack/echo/v4"
)

func (controllers *Controllers) CreateProduct(eCtx echo.Context) error {
	var err error
	ctx := eCtx.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	err = middleware.CheckIsStaff(eCtx)
	if err != nil {
		return controllers.StatusErrorPermission(eCtx, err)
	}
	req := new(dto.CreateProductRequest)

	err = eCtx.Bind(&req)

	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorBindData, err.Error(), utils.FuncName())
		return controllers.StatusBadRequest(eCtx, message, resp)
	}

	product, err := controllers.Services.CreateProduct(ctx, req)

	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorBindData, err.Error(), utils.FuncName())
		return controllers.StatusBadRequest(eCtx, message, resp)
	}

	return controllers.StatusOkResponse(eCtx, &product)
}

func (controllers *Controllers) GetAllProduct(eCtx echo.Context) error {
	var err error
	ctx := eCtx.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	res, err := controllers.Services.GetAllProduct(ctx)
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorRetriveData, err.Error(), utils.FuncName())
		return controllers.StatusInternalServerError(eCtx, message, resp)
	}
	return controllers.StatusOkResponse(eCtx, &res)
}

func (controllers *Controllers) GetProductById(eCtx echo.Context) error {
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

	res, err := controllers.Services.GetProductById(ctx, id)

	if err != nil {
		return controllers.StatusError(eCtx, err)
	}

	return controllers.StatusOkResponse(eCtx, &res)
}

func (controllers *Controllers) UpdateProductById(eCtx echo.Context) error {
	var err error
	ctx := eCtx.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	err = middleware.CheckIsStaff(eCtx)
	if err != nil {
		return controllers.StatusErrorPermission(eCtx, err)
	}
	req := new(dto.UpdateProductRequest)

	err = eCtx.Bind(&req)

	if err != nil {
		fmt.Println("hahah 1", err)
		return controllers.StatusErrorBadRequest(eCtx, err)
	}

	res, err := controllers.Services.UpdateProductById(ctx, req)

	if err != nil {
		return controllers.StatusError(eCtx, err)
	}

	return controllers.StatusOkResponse(eCtx, &res)
}

func (controllers *Controllers) DeleteProductById(eCtx echo.Context) error {
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

	err = controllers.Services.DeleteProductById(ctx, id)

	if err != nil {
		return controllers.StatusError(eCtx, err)
	}

	return controllers.StatusOkResponse(eCtx, nil)
}
