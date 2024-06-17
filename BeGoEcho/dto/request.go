package dto

import "time"

type CreateDiscountRequest struct {
	Name      string `json:"name"`
	BeginAt   string `json:"begin_at"`
	ExpireAt  string `json:"expire_at"`
	IsDeleted bool   `json:"is_deleted"`
}
type UpdateDiscountRequest struct {
	Id       int64      `json:"id"`
	Name     string     `json:"name"`
	BeginAt  *time.Time `json:"begin_at,omitempty"`
	ExpireAt *time.Time `json:"expire_at,omitempty"`
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
