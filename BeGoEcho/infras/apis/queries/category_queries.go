package queries

import (
	"context"
	"fmt"

	"github.com/CMDezz/KB/dto"
	"github.com/CMDezz/KB/utils/constants"
)

func (queries *Queries) DBCreateCategory(ctx context.Context, category *dto.Category) (*dto.Category, error) {

	res := dto.Category{}
	fmt.Println(category.Parent)
	var err error
	if category.Parent != nil && *category.Parent == 0 {
		query := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING *",
			constants.TableCategory,
		)

		err = queries.Postgres.SQLxDBContext.QueryRowxContext(ctx, query, category.Name).StructScan(&res)

	} else {

		query := fmt.Sprintf("INSERT INTO %s (name, parent) VALUES ($1, $2) RETURNING *",
			constants.TableCategory,
		)

		err = queries.Postgres.SQLxDBContext.QueryRowxContext(ctx, query, category.Name, category.Parent).StructScan(&res)

	}
	if err != nil {
		queries.Postgres.HandleError(err, "")
		return nil, err
	}

	return &res, nil
}

func (queries *Queries) DBGetAllCategory(ctx context.Context) (*[]dto.Category, error) {
	query := fmt.Sprintf("SELECT * FROM %v",
		constants.TableCategory,
	)

	var res []dto.Category

	err := queries.Postgres.SQLxDBContext.SelectContext(ctx, &res, query)

	if err != nil {
		queries.Postgres.HandleError(err, query)
		return nil, err
	}

	return &res, nil
}

func (queries *Queries) DBGetCategoryById(ctx context.Context, id int64) (*dto.Category, error) {
	query := fmt.Sprintf("SELECT * FROM %v WHERE id=%d",
		constants.TableCategory, id,
	)

	var res dto.Category

	err := queries.Postgres.SQLxDBContext.GetContext(ctx, &res, query)

	if err != nil {
		queries.Postgres.HandleError(err, query)
		return nil, err
	}

	return &res, nil
}

func (queries *Queries) DBUpdateCategoryById(ctx context.Context, category *dto.Category) (*dto.Category, error) {
	query := fmt.Sprintf(`
        UPDATE %v
        SET name = $2, parent = $3, is_deleted = $4
        WHERE id = $1
        RETURNING *
    `, constants.TableCategory)

	var res dto.Category
	err := queries.Postgres.SQLxDBContext.QueryRowxContext(ctx, query, category.Id, category.Name, category.Parent, category.IsDeleted).StructScan(&res)
	if err != nil {
		queries.Postgres.HandleError(err, query)
		return nil, err
	}

	return &res, nil
}

func (queries *Queries) DBDeleteCategoryById(ctx context.Context, id int64) error {
	query := fmt.Sprintf("DELETE FROM %v WHERE id=%d;",
		constants.TableCategory, id,
	)

	_, err := queries.Postgres.SQLxDBContext.ExecContext(ctx, query)

	if err != nil {
		queries.Postgres.HandleError(err, query)
		return err
	}

	return nil
}
