package utils

import (
	"runtime"
	"strconv"
	"time"

	"github.com/CMDezz/KB/utils/constants"
)

func FuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

// ToInt Convert string to int type
func ToInt(v string) (int, error) {
	n, err := strconv.Atoi(v)
	if err != nil {
		return -1, err
	}

	return n, nil
}

func ToInt64(v string) (int64, error) {
	n, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return -1, err
	}

	return n, nil
}

func ToTime(v string) (time.Time, error) {
	parsedTime, err := time.Parse("2006-01-02T15:04:05.999Z", v)
	if err != nil {
		return time.Time{}, err // Return the zero value of time.Time
	}
	return parsedTime, nil
}

func IsZeroTime(v *time.Time) bool {
	if v == nil {
		return true
	}

	if *v == constants.TimeEmpty {
		return true
	}

	return false

}

func EmptyStringIfNil(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}

func EmptyValueIfNil[T any](v *T) T {
	if v == nil {
		var zero T
		return zero
	}
	return *v
}

func Includes[T comparable](a []T, v T) bool {
	for _, t := range a {
		if t == v {
			return true
		}
	}
	return false
}
