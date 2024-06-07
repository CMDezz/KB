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

type GetByIdRequest struct {
	Id int64 `json:"id"`
}
