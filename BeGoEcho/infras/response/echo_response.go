package response

import (
	"database/sql"
	"net/http"

	"github.com/CMDezz/KB/gerror"
	"github.com/CMDezz/KB/utils"
	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
}

func (baseResponse *BaseResponse) StatusOkResponse(context echo.Context, data any) error {
	_, response := NewResponse(data, "Successful")
	return context.JSON(http.StatusOK, response)
}

func (baseReponse *BaseResponse) StatusErrorResponse(context echo.Context, errorCode int, mess string, data ErrorResponse) error {
	return context.JSON(errorCode, ErrorResponse{
		Message:   mess,
		ErrorCode: data.ErrorCode,
		Exception: data.Exception,
	})
}

func (baseReponse *BaseResponse) StatusBadRequest(context echo.Context, mess string, err ErrorResponse) error {
	return baseReponse.StatusErrorResponse(context, http.StatusBadRequest, mess, err)
}

func (baseReponse *BaseResponse) StatusInternalServerError(context echo.Context, mess string, err ErrorResponse) error {
	return baseReponse.StatusErrorResponse(context, http.StatusInternalServerError, mess, err)
}

func (baseReponse *BaseResponse) StatusNotFound(context echo.Context, mess string, err ErrorResponse) error {
	return baseReponse.StatusErrorResponse(context, http.StatusNotFound, mess, err)
}

func (baseReponse *BaseResponse) StatusError(context echo.Context, err error) error {
	if err == sql.ErrNoRows {
		message, resp := NewErrorResponse(gerror.ErrorNotFound, err.Error(), utils.FuncName())
		return baseReponse.StatusErrorResponse(context, http.StatusNotFound, message, resp)
	}

	message, resp := NewErrorResponse(gerror.ErrorInternal, err.Error(), utils.FuncName())
	return baseReponse.StatusErrorResponse(context, http.StatusInternalServerError, message, resp)
}
func (baseReponse *BaseResponse) StatusErrorBadRequest(context echo.Context, err error) error {
	message, resp := NewErrorResponse(gerror.ErrorBadRequest, err.Error(), utils.FuncName())

	return baseReponse.StatusErrorResponse(context, http.StatusBadRequest, message, resp)
}
func (baseReponse *BaseResponse) StatusErrorPermission(context echo.Context, err error) error {
	message, resp := NewErrorResponse(gerror.ErrorPermission, err.Error(), utils.FuncName())

	return baseReponse.StatusErrorResponse(context, http.StatusUnauthorized, message, resp)
}
