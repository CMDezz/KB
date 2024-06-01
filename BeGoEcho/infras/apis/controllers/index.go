package controllers

import (
	"github.com/CMDezz/KB/infras/apis/services"
	"github.com/CMDezz/KB/infras/response"
)

type Controllers struct {
	response.BaseResponse
	Services services.IServices
}
