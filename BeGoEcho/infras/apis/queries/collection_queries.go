package queries

import (
	"context"
	"fmt"

	"github.com/CMDezz/KB/dto"
	"github.com/CMDezz/KB/utils/constants"
)

func (queries *Queries) DBCreateCollection(ctx context.Context, collection *dto.Collection) (*dto.Collection, error) {
	query := fmt.Sprintf(`INSERT INTO %s (name, short_desc, "desc", article) VALUES ($1, $2, $3, $4) RETURNING *`,
		constants.TableCollection,
	)

	res := dto.Collection{}
	err := queries.Postgres.SQLxDBContext.QueryRowxContext(ctx, query, collection.Name, collection.ShortDesc, collection.Desc, collection.Article).StructScan(&res)

	if err != nil {
		queries.Postgres.HandleError(err, query)
		return nil, err
	}

	return &res, nil
}

func (queries *Queries) DBGetAllCollection(ctx context.Context) (*[]dto.Collection, error) {
	query := fmt.Sprintf("SELECT * FROM %v",
		constants.TableCollection,
	)

	var res []dto.Collection

	err := queries.Postgres.SQLxDBContext.SelectContext(ctx, &res, query)

	if err != nil {
		queries.Postgres.HandleError(err, query)
		return nil, err
	}

	return &res, nil
}

func (queries *Queries) DBGetCollectionById(ctx context.Context, id int64) (*dto.Collection, error) {
	query := fmt.Sprintf("SELECT * FROM %v WHERE id=%d",
		constants.TableCollection, id,
	)

	var res dto.Collection

	err := queries.Postgres.SQLxDBContext.GetContext(ctx, &res, query)

	if err != nil {
		queries.Postgres.HandleError(err, query)
		return nil, err
	}

	return &res, nil
}

func (queries *Queries) DBUpdateCollectionById(ctx context.Context, collection *dto.Collection) (*dto.Collection, error) {
	query := fmt.Sprintf(`
        UPDATE %v
        SET name = $2, short_desc = $3, "desc" = $4, article = $5
        WHERE id = $1
        RETURNING *
    `, constants.TableCollection)

	var res dto.Collection
	err := queries.Postgres.SQLxDBContext.QueryRowxContext(ctx, query, collection.Id, collection.Name, collection.ShortDesc, collection.Desc, collection.Article).StructScan(&res)
	if err != nil {
		queries.Postgres.HandleError(err, query)
		return nil, err
	}

	return &res, nil
}

func (queries *Queries) DBDeleteCollectionById(ctx context.Context, id int64) error {
	query := fmt.Sprintf("DELETE FROM %v WHERE id=%d;",
		constants.TableCollection, id,
	)

	_, err := queries.Postgres.SQLxDBContext.ExecContext(ctx, query)

	if err != nil {
		queries.Postgres.HandleError(err, query)
		return err
	}

	return nil
}
