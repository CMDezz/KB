package response

import (
	"fmt"
	"net/http"

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

func (baseReponse *BaseResponse) Testz(context echo.Context, mess string) string {
	// _, response := NewErrorResponse(http.StatusNotFound, mess, exception)
	// return baseReponse.StatusErrorResponse(context, response)
	return fmt.Sprintf(mess)
}
