package queries

import (
	"context"
	"fmt"

	"github.com/CMDezz/KB/dto"
	"github.com/CMDezz/KB/utils/constants"
)

func (queries *Queries) DBCreateDiscount(ctx context.Context, discount *dto.Discount) (*dto.Discount, error) {
	fmt.Println("--- ", discount)
	query := fmt.Sprintf("INSERT INTO %s (name, begin_at, expire_at,is_deleted) VALUES ($1, $2, $3, $4) RETURNING *",
		constants.TableDiscount,
	)

	res := dto.Discount{}
	err := queries.Postgres.SQLxDBContext.QueryRowxContext(ctx, query, discount.Name, discount.BeginAt, discount.ExpireAt, discount.IsDeleted).StructScan(&res)

	fmt.Println(res)

	if err != nil {
		queries.Postgres.HandleError(err, query)
		return nil, err
	}

	return &res, nil
}
