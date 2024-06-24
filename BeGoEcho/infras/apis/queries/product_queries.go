package queries

import (
	"context"
	"fmt"

	"github.com/CMDezz/KB/dto"
	"github.com/CMDezz/KB/utils/constants"
	"github.com/lib/pq"
)

func (queries *Queries) DBCreateProduct(ctx context.Context, product *dto.Product) (*dto.Product, error) {
	query := fmt.Sprintf(`INSERT INTO %s (name, short_desc, "desc", article, discount_applied) VALUES ($1, $2, $3, $4, $5) RETURNING *`,
		constants.TableProduct,
	)

	res := dto.Product{}
	err := queries.Postgres.SQLxDBContext.QueryRowxContext(ctx, query, product.Name, product.ShortDesc, product.Desc, product.Article, product.DiscountApplied).StructScan(&res)

	if err != nil {
		queries.Postgres.HandleError(err, query)
		return nil, err
	}

	return &res, nil
}

func (queries *Queries) DBGetAllProduct(ctx context.Context) (*[]dto.Product, error) {
	query := fmt.Sprintf(`SELECT * 
	FROM %v p 
	LEFT JOIN discounts d ON p.discount_applied = d.id;`,
		constants.TableProduct,
	)

	var res []dto.Product

	err := queries.Postgres.SQLxDBContext.SelectContext(ctx, &res, query)

	if err != nil {
		queries.Postgres.HandleError(err, query)
		return nil, err
	}

	return &res, nil
}

func (queries *Queries) DBGetProductById(ctx context.Context, id int64) (*dto.Product, error) {
	query := fmt.Sprintf("SELECT * FROM %v WHERE id=%d",
		constants.TableProduct, id,
	)

	var res dto.Product

	err := queries.Postgres.SQLxDBContext.GetContext(ctx, &res, query)

	if err != nil {
		queries.Postgres.HandleError(err, query)
		return nil, err
	}

	return &res, nil
}

func (queries *Queries) DBUpdateProductById(ctx context.Context, product *dto.Product) (*dto.Product, error) {
	query := fmt.Sprintf(`
        UPDATE %v
        SET name = $2, short_desc = $3, "desc" = $4, article = $5, discount_applied = $6
        WHERE id = $1
        RETURNING *
    `, constants.TableProduct)

	var res dto.Product
	err := queries.Postgres.SQLxDBContext.QueryRowxContext(ctx, query, product.Id, product.Name, product.ShortDesc, product.Desc, product.Article, product.DiscountApplied).StructScan(&res)
	if err != nil {
		queries.Postgres.HandleError(err, query)
		return nil, err
	}

	return &res, nil
}

func (queries *Queries) DBDeleteProductById(ctx context.Context, id int64) error {
	query := fmt.Sprintf("DELETE FROM %v WHERE id=%d;",
		constants.TableProduct, id,
	)

	_, err := queries.Postgres.SQLxDBContext.ExecContext(ctx, query)

	if err != nil {
		queries.Postgres.HandleError(err, query)
		return err
	}

	return nil
}
func (queries *Queries) DBDeleteVariantById(ctx context.Context, id int64) error {
	query := fmt.Sprintf("DELETE FROM %v WHERE id=%d;",
		constants.TableProductVariant, id,
	)

	_, err := queries.Postgres.SQLxDBContext.ExecContext(ctx, query)

	if err != nil {
		queries.Postgres.HandleError(err, query)
		return err
	}

	return nil
}

func (queries *Queries) DBCreateProductVariant(ctx context.Context, product *dto.ProductVariant) (*dto.ProductVariant, error) {
	query := fmt.Sprintf(`INSERT INTO %s (name, img_main, imgs_detail, qty, price,variant_on_product) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *`,
		constants.TableProductVariant,
	)

	res := dto.ProductVariant{}
	err := queries.Postgres.SQLxDBContext.QueryRowxContext(ctx, query, product.Name, product.ImgMain, pq.StringArray(product.ImgsDetail), product.Qty, product.Price, product.VariantOnProduct).StructScan(&res)

	if err != nil {
		queries.Postgres.HandleError(err, query)
		return nil, err
	}

	return &res, nil
}
func (queries *Queries) DBUpdateVariantById(ctx context.Context, product *dto.ProductVariant) (*dto.ProductVariant, error) {
	query := fmt.Sprintf(`
	UPDATE %v
	SET name = $2, img_main = $3, imgs_detail = $4, qty = $5, price = $6, variant_on_product = $7
	WHERE id = $1
	RETURNING *
`, constants.TableProductVariant)

	res := dto.ProductVariant{}
	err := queries.Postgres.SQLxDBContext.QueryRowxContext(ctx, query, product.Id, product.Name, product.ImgMain, pq.StringArray(product.ImgsDetail), product.Qty, product.Price, product.VariantOnProduct).StructScan(&res)

	if err != nil {
		queries.Postgres.HandleError(err, query)
		return nil, err
	}

	return &res, nil
}

func (queries *Queries) DBGetVariantsByProductId(ctx context.Context, id int64) (*[]dto.ProductVariant, error) {
	query := fmt.Sprintf(`SELECT * FROM %s WHERE variant_on_product=%d`,
		constants.TableProductVariant, id,
	)

	var res []dto.ProductVariant

	err := queries.Postgres.SQLxDBContext.SelectContext(ctx, &res, query)

	if err != nil {
		queries.Postgres.HandleError(err, query)
		return nil, err
	}

	return &res, nil
}
