package testutil

import "testing"

func TestNilPointerError(t *testing.T) {
	var nilSlice []any = nil

	t.Run("should return an error for nil slice", func(t *testing.T) {
		// if NilPointerError[string, string, []string](nilSlice) == nil {
		if CheckNilPointerError[any, *any](nilSlice) == nil {
			t.Fatal("should return a nil slice error")
		}
	})
}

func TestZeroLengthError(t *testing.T) {
	var zeroLengthSlice = []any{}

	t.Run("should return an error for slice of length 0", func(t *testing.T) {
		if CheckZeroLengthError[any, *any](zeroLengthSlice) == nil {
			t.Fatal("should return a zero-length slice error")
		}
	})
}

func TestEmptyStringInSliceError(t *testing.T) {
	var sliceWithEmptyElement = []string{"a", ""}

	t.Run("should return an error for slice with empty element", func(t *testing.T) {
		if CheckEmptyStringInSliceError(sliceWithEmptyElement) == nil {
			t.Fatal("should return an empty slice element error")
		}
	})
}
