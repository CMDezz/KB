package dto

import (
	"time"

	"github.com/lib/pq"
)

type Discount struct {
	Id        int64     `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	BeginAt   time.Time `json:"begin_at" db:"begin_at"`
	ExpireAt  time.Time `json:"expire_at" db:"expire_at"`
	Type      int       `json:"type" db:"type"`
	Value     float64   `json:"value" db:"value"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	IsDeleted bool      `json:"is_deleted" db:"is_deleted"`
}
type Category struct {
	Id        int64     `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Parent    *int64    `json:"parent" db:"parent"`
	IsDeleted bool      `json:"is_deleted" db:"is_deleted"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Account struct {
	Id            int64     `json:"id" db:"id"`
	Username      string    `json:"username" db:"username"`
	HasedPassword string    `json:"hased_password" db:"hased_password"`
	Email         string    `json:"email" db:"email"`
	IsVerified    bool      `json:"is_verified" db:"is_verified"`
	FullName      string    `json:"full_name" db:"full_name"`
	PhoneFloat    string    `json:"phone_float" db:"phone_float"`
	Role          int64     `json:"role" db:"role"`
	IsDeleted     bool      `json:"is_deleted" db:"is_deleted"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
}

type Collection struct {
	Id        int64     `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	ShortDesc string    `json:"short_desc" db:"short_desc"`
	Desc      string    `json:"desc" db:"desc"`
	Article   string    `json:"article" db:"article"`
	IsDeleted bool      `json:"is_deleted" db:"is_deleted"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Product struct {
	Id              int64     `json:"id" db:"id"`
	Name            string    `json:"name" db:"name"`
	ShortDesc       string    `json:"short_desc" db:"short_desc"`
	Desc            string    `json:"desc" db:"desc"`
	DiscountApplied *int64    `json:"discount_applied" db:"discount_applied"`
	Article         string    `json:"article" db:"article"`
	IsDeleted       bool      `json:"is_deleted" db:"is_deleted"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
}
type ProductVariant struct {
	Id               int64    `json:"id" db:"id"`
	Name             string   `json:"name" db:"name"`
	ImgMain          string   `json:"img_main" db:"img_main"`
	ImgsDetail       pq.StringArray `json:"imgs_detail" db:"imgs_detail"`
	Qty              int      `json:"qty" db:"qty"`
	Price            float64  `json:"price" db:"price"`
	VariantOnProduct int64    `json:"variant_on_product" db:"variant_on_product"`
	IsDeleted        bool     `json:"is_deleted" db:"is_deleted"`
}

type ProductWithVariant struct {
	*Product
	Variants *[]ProductVariant
}
