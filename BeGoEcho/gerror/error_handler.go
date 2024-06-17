package gerror

import (
	"github.com/lib/pq"
)

func IsDuplicateError(err error) (bool, string) {
	if pgErr, ok := err.(*pq.Error); ok {
		// PostgreSQL specific error codes for duplicate key or unique violation
		if pgErr.Code == "23505" { // Unique violation error code
			return true, pgErr.Detail
		}
	}
	return false, ""
}
