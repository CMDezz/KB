package constants

import "time"

const (
	MinLengthSecretKey   int           = 6
	ExpiresTokenDuration time.Duration = time.Duration(1) * time.Hour
)

const (
	ENUM_PER_USER  int64 = 0
	ENUM_PER_STAFF int64 = 1
	ENUM_PER_ADMIN int64 = 2
)
