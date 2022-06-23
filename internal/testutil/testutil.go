package testutil

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

func Check(t *testing.T, errs ...error) {
	for _, err := range errs {
		if err != nil {
			t.Fatal(err)
		}
	}
}

// Accepts pointers to any type, and all slices and maps.
func CheckNilPointerError[T any, K comparable, PtrT *T | []T | map[K]T](ptr PtrT) error {
	if ptr == nil {
		return errors.New("pointer is nil")
	}
	return nil
}

func CheckZeroLengthError[T any, K comparable, LenT []T | string | map[K]T](t LenT) error {
	if len(t) == 0 {
		return errors.New("length is zero")
	}
	return nil
}

func CheckEmptyStringInSliceError(input []string) error {
	for i, str := range input {
		if str == "" {
			return errors.New("slice contains empty string at index " + strconv.Itoa(i))
		}
	}
	return nil
}

func CheckNotEqualError[T comparable](want T, got T) error {
	if want != got {
		return errors.New(fmt.Sprintf("want: %#v, but got: %#v", want, got))
	}
	return nil
}
