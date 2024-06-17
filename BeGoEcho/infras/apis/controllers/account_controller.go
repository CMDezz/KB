package controllers

import (
	"context"

	"github.com/CMDezz/KB/dto"
	"github.com/CMDezz/KB/gerror"
	"github.com/CMDezz/KB/infras/middleware"
	"github.com/CMDezz/KB/infras/response"
	"github.com/CMDezz/KB/utils"
	"github.com/labstack/echo/v4"
)

func (controllers *Controllers) GetAllAccount(eCtx echo.Context) error {
	var err error
	ctx := eCtx.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	err = middleware.CheckIsStaff(eCtx)
	if err != nil {
		return controllers.StatusErrorPermission(eCtx, err)
	}
	res, err := controllers.Services.GetAllAccount(ctx)
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorRetriveData, err.Error(), utils.FuncName())
		return controllers.StatusInternalServerError(eCtx, message, resp)
	}
	return controllers.StatusOkResponse(eCtx, &res)
}

func (controllers *Controllers) CreateAccount(eCtx echo.Context) error {
	var err error
	ctx := eCtx.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	err = middleware.CheckIsStaff(eCtx)
	if err != nil {
		return controllers.StatusErrorPermission(eCtx, err)
	}
	req := new(dto.CreateAccountRequest)

	err = eCtx.Bind(&req)

	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorBindData, err.Error(), utils.FuncName())
		return controllers.StatusBadRequest(eCtx, message, resp)
	}

	res, err := controllers.Services.CreateAccount(ctx, req)
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorRetriveData, err.Error(), utils.FuncName())
		return controllers.StatusInternalServerError(eCtx, message, resp)
	}
	return controllers.StatusOkResponse(eCtx, &res)
}

func (controllers *Controllers) LoginAccount(eCtx echo.Context) error {
	ctx := eCtx.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	req := new(dto.LoginAccountRequest)

	err := eCtx.Bind(&req)

	if err != nil {
		if err != nil {
			message, resp := response.NewErrorResponse(gerror.ErrorRetriveData, err.Error(), utils.FuncName())
			return controllers.StatusInternalServerError(eCtx, message, resp)
		}
	}

	res, err := controllers.Services.LoginAccount(ctx, req, controllers.Token)
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorRetriveData, err.Error(), utils.FuncName())
		return controllers.StatusInternalServerError(eCtx, message, resp)
	}
	return controllers.StatusOkResponse(eCtx, &res)
}
