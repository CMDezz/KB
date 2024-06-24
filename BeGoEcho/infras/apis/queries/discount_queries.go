package queries

import (
	"context"
	"fmt"

	"github.com/CMDezz/KB/dto"
	"github.com/CMDezz/KB/utils/constants"
)

func (queries *Queries) DBCreateDiscount(ctx context.Context, discount *dto.Discount) (*dto.Discount, error) {
	query := fmt.Sprintf("INSERT INTO %s (name, begin_at, expire_at,type ,value) VALUES ($1, $2, $3,$4,$5) RETURNING *",
		constants.TableDiscount,
	)

	res := dto.Discount{}
	err := queries.Postgres.SQLxDBContext.QueryRowxContext(ctx, query, discount.Name, discount.BeginAt, discount.ExpireAt, discount.Type, discount.Value).StructScan(&res)

	if err != nil {
		queries.Postgres.HandleError(err, query)
		return nil, err
	}

	return &res, nil
}

func (queries *Queries) DBGetAllDiscount(ctx context.Context) (*[]dto.Discount, error) {
	query := fmt.Sprintf("SELECT * FROM %v;",
		constants.TableDiscount,
	)

	var res []dto.Discount

	err := queries.Postgres.SQLxDBContext.SelectContext(ctx, &res, query)

	if err != nil {
		queries.Postgres.HandleError(err, query)
		return nil, err
	}

	return &res, nil
}

func (queries *Queries) DBGetDiscountById(ctx context.Context, id int64) (*dto.Discount, error) {
	query := fmt.Sprintf("SELECT * FROM %v WHERE id=%d",
		constants.TableDiscount, id,
	)

	var res dto.Discount

	err := queries.Postgres.SQLxDBContext.GetContext(ctx, &res, query)

	if err != nil {
		queries.Postgres.HandleError(err, query)
		return nil, err
	}

	return &res, nil
}

func (queries *Queries) DBUpdateDiscountById(ctx context.Context, discount *dto.Discount) (*dto.Discount, error) {
	query := fmt.Sprintf(`
        UPDATE %v
        SET name = $2, begin_at = $3, expire_at = $4, type = $5, value = $6
        WHERE id = $1
        RETURNING *
    `, constants.TableDiscount)

	var res dto.Discount
	err := queries.Postgres.SQLxDBContext.QueryRowxContext(ctx, query, discount.Id, discount.Name, discount.BeginAt, discount.ExpireAt, discount.Type, discount.Value).StructScan(&res)
	if err != nil {
		queries.Postgres.HandleError(err, query)
		return nil, err
	}

	return &res, nil
}

func (queries *Queries) DBDeleteDiscountById(ctx context.Context, id int64) error {
	query := fmt.Sprintf("DELETE FROM %v WHERE id=%d;",
		constants.TableDiscount, id,
	)

	_, err := queries.Postgres.SQLxDBContext.ExecContext(ctx, query)

	if err != nil {
		queries.Postgres.HandleError(err, query)
		return err
	}

	return nil
}
