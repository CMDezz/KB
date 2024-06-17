package controllers

import (
	"github.com/CMDezz/KB/infras/apis/services"
	"github.com/CMDezz/KB/infras/response"
	"github.com/CMDezz/KB/infras/token"
)

type Controllers struct {
	response.BaseResponse
	Services services.IServices
	Token    token.JWTTokenMaker
}
