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
