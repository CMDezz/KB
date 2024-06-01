package dto

type CreateDiscountRequest struct {
	Name      string `json:"name"`
	BeginAt   string `json:"begin_at"`
	ExpireAt  string `json:"expire_at"`
	IsDeleted bool   `json:"is_deleted"`
}
