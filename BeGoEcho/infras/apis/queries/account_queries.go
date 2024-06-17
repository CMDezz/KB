package queries

import (
	"context"
	"fmt"

	"github.com/CMDezz/KB/dto"
	"github.com/CMDezz/KB/utils/constants"
)

func (queries *Queries) DBGetAllAccount(ctx context.Context) (*[]dto.Account, error) {
	query := fmt.Sprintf("SELECT * FROM %v",
		constants.TableAccount,
	)

	var res []dto.Account

	err := queries.Postgres.SQLxDBContext.SelectContext(ctx, &res, query)

	if err != nil {
		queries.Postgres.HandleError(err, query)
		return nil, err
	}

	return &res, nil
}

func (queries *Queries) DBCreateAccount(ctx context.Context, account *dto.Account) (*dto.Account, error) {
	query := fmt.Sprintf("INSERT INTO %s (username, hased_password,email,full_name,phone_float,role,is_verified,is_deleted,created_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING *",
		constants.TableAccount,
	)

	res := dto.Account{}

	err := queries.Postgres.SQLxDBContext.QueryRowxContext(ctx, query, account.Username, account.HasedPassword, account.Email, account.FullName, account.PhoneFloat, account.Role, account.IsVerified, account.IsDeleted, account.CreatedAt).StructScan(&res)

	if err != nil {
		queries.Postgres.HandleError(err, query)
		return nil, err
	}

	return &res, nil
}

func (queries *Queries) DBGetAccountByUsername(ctx context.Context, username string) (*dto.Account, error) {
	query := fmt.Sprintf("SELECT * FROM %v WHERE username=$1",
		constants.TableAccount,
	)

	var res = dto.Account{}

	err := queries.Postgres.SQLxDBContext.GetContext(ctx, &res, query, username)

	if err != nil {
		queries.Postgres.HandleError(err, query)
		return nil, err
	}

	return &res, nil
}
