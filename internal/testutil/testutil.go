package testutil

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

func FailOnError(t *testing.T, errs ...error) {
	for _, err := range errs {
		if err != nil {
			t.Fatal(err)
		}
	}
}

// CheckNotEqualError returns an error if two comparable elements are not equal to each other
func CheckNotEqualError[T comparable](actual T, expected T) error {
	if actual != expected {
		return errors.New(fmt.Sprintf("got unexpected value %#v, but wanted %#v", actual, expected))
	}
	return nil
}

// CheckNilPointerError returns an error if a pointer to any type, or a slice or map is nil
func CheckNilPointerError[T any, K comparable, PtrT *T | []T | map[K]T](ptr PtrT) error {
	if ptr == nil {
		return errors.New("pointer is nil")
	}
	return nil
}

// CheckZeroLengthError returns an error if the length of the slice is zero.
// NB: A nil slice will also return an error
func CheckZeroLengthError[T any, K comparable, LenT []T | string | map[K]T](t LenT) error {
	if len(t) == 0 {
		return errors.New("length is zero")
	}
	return nil
}

// CheckEmptyStringInSliceError returns an error if the slice contains an empty string.
// NB: no error is returned if the slice is nil or empty.
func CheckEmptyStringInSliceError(input []string) error {
	for i, str := range input {
		if str == "" {
			return errors.New("slice contains empty string at index " + strconv.Itoa(i))
		}
	}
	return nil
}
