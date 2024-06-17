package services

import (
	"context"
	"fmt"
	"time"

	"github.com/CMDezz/KB/dto"
	"github.com/CMDezz/KB/infras/logger"
	"github.com/CMDezz/KB/infras/token"
	"github.com/CMDezz/KB/utils"
	"github.com/CMDezz/KB/utils/constants"
)

func (services *Services) GetAllAccount(ctx context.Context) (*[]dto.Account, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.TimeoutRequestDefault)
	defer cancel()

	res, err := services.Queries.DBGetAllAccount(ctx)

	if err != nil {
		logger.Error("SERVICE - GetAllAccount - Error %v", err)
		return nil, err
	}

	return res, nil
}

func (services *Services) CreateAccount(ctx context.Context, req *dto.CreateAccountRequest) (*dto.Account, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.TimeoutServerDefault)
	defer cancel()

	//TODO : KIỂM TRA QUYỀN

	//TODO: hased password
	hased_password, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	newAccount := &dto.Account{
		Username:      req.Username,
		Email:         req.Email,
		HasedPassword: hased_password,
	}

	newAccount.Role = utils.EmptyValueIfNil(req.Role)
	newAccount.PhoneFloat = utils.EmptyValueIfNil(req.PhoneFloat)
	newAccount.FullName = utils.EmptyValueIfNil(req.FullName)
	newAccount.CreatedAt = time.Now()
	newAccount.IsDeleted = false
	newAccount.IsVerified = false

	fmt.Println(newAccount)

	//tạo account
	res, err := services.Queries.DBCreateAccount(ctx, newAccount)

	if err != nil {
		logger.Error("SERVICE - GetAllAccount - Error %v", err)
		return nil, err
	}

	return res, nil

}

func (services *Services) LoginAccount(ctx context.Context, req *dto.LoginAccountRequest, tokenMaker token.JWTTokenMaker) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.TimeoutRequestDefault)
	defer cancel()

	account, err := services.Queries.DBGetAccountByUsername(ctx, req.Username)

	if err != nil {
		logger.Error("SERVICE - LoginAccount - Error %v", err)
		return "", err
	}

	err = utils.CheckPassword(req.Password, account.HasedPassword)
	if err != nil {
		logger.Error("SERVICE - LoginAccount - Error %v", err)
		return "", err
	}

	//create token

	jwtToken, _, err := tokenMaker.NewToken(account.Username, account.Role, constants.ExpiresTokenDuration)

	if err != nil {
		logger.Error("SERVICE - LoginAccount - Error %v", err)
		return "", err
	}

	return jwtToken, nil
}
