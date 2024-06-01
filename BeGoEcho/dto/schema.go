package dto

import "time"

type Discount struct {
	Id        int64     `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	BeginAt   time.Time `json:"begin_at" db:"begin_at"`
	ExpireAt  time.Time `json:"expire_at" db:"expire_at"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	IsDeleted bool      `json:"is_deleted" db:"is_deleted"`
}
