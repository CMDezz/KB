package dto

import "time"

type CreateDiscountRequest struct {
	Name     string  `json:"name"`
	BeginAt  string  `json:"begin_at"`
	ExpireAt string  `json:"expire_at"`
	Type     int     `json:"type"`
	Value    float64 `json:"Value"`
}
type UpdateDiscountRequest struct {
	Id       int64      `json:"id"`
	Name     string     `json:"name"`
	BeginAt  *time.Time `json:"begin_at,omitempty"`
	ExpireAt *time.Time `json:"expire_at,omitempty"`
	Type     int        `json:"type"`
	Value    float64    `json:"Value"`
}
type CreateCategoryRequest struct {
	Name   string `json:"name"`
	Parent *int64 `json:"parent"`
}
type UpdateCategoryRequest struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Parent    *int64 `json:"parent"`
	IsDeleted *bool  `json:"is_deleted"`
}

type GetByIdRequest struct {
	Id int64 `json:"id"`
}

type CreateAccountRequest struct {
	Username   string  `json:"username"`
	Password   string  `json:"password"`
	Email      string  `json:"email"`
	FullName   *string `json:"full_name"`
	PhoneFloat *string `json:"phone_float"`
	Role       *int64  `json:"role"`
}

type LoginAccountRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateCollectionRequest struct {
	Name      string  `json:"name"`
	ShortDesc string  `json:"short_desc"`
	Desc      string  `json:"desc"`
	Article   *string `json:"article"`
}

type UpdateCollectionRequest struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	ShortDesc string `json:"short_desc"`
	Desc      string `json:"desc"`
	Article   string `json:"article"`
	IsDeleted bool   `json:"is_deleted"`
}

type CreateProductRequest struct {
	Name            string                         `json:"name"`
	ShortDesc       string                         `json:"short_desc"`
	Desc            string                         `json:"desc"`
	Article         *string                        `json:"article"`
	DiscountApplied *int64                         `json:"discount_applied"`
	Variants        *[]CreateProductVariantRequest `json:"variants"`
}

type CreateProductVariantRequest struct {
	Name       string   `json:"name"`
	ImgMain    string   `json:"img_main"`
	ImgsDetail []string `json:"imgs_detail"`
	Qty        int      `json:"qty"`
	Price      float64  `json:"price"`
}

type UpdateProductRequest struct {
	Id              int64                         `json:"id"`
	Name            string                        `json:"name"`
	ShortDesc       string                        `json:"short_desc"`
	Desc            string                        `json:"desc"`
	DiscountApplied *int64                        `json:"discount_applied"`
	Article         string                        `json:"article"`
	IsDeleted       bool                          `json:"is_deleted"`
	Variants        []UpdateProductVariantRequest `json:"variants"`
}
type UpdateProductVariantRequest struct {
	Id         *int64   `json:"id"`
	Name       string   `json:"name"`
	ImgMain    string   `json:"img_main"`
	ImgsDetail []string `json:"imgs_detail"`
	Qty        int      `json:"qty"`
	Price      float64  `json:"price"`
	IsDeleted  bool     `json:"is_deleted"`
}
