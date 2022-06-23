package testutil

import "testing"

func TestCheckNilPointerError(t *testing.T) {
	var nilSlice []any = nil
	var notNilSlice []any = []any{}

	t.Run("should return an error for nil slice", func(t *testing.T) {
		if CheckNilPointerError[any, *any](nilSlice) == nil {
			t.Fatal("should return a nil slice error")
		}
		if CheckNilPointerError[any, *any](notNilSlice) != nil {
			t.Fatal("should NOT return a nil slice error")
		}
	})
}

func TestCheckZeroLengthError(t *testing.T) {
	var zeroLengthSlice = []any{}

	t.Run("should return an error for slice of length 0", func(t *testing.T) {
		if CheckZeroLengthError[any, *any](zeroLengthSlice) == nil {
			t.Fatal("should return a zero-length slice error")
		}
	})
}

func TestCheckEmptyStringInSliceError(t *testing.T) {
	var sliceWithEmptyElement = []string{"a", ""}

	t.Run("should return an error for slice with empty element", func(t *testing.T) {
		if CheckEmptyStringInSliceError(sliceWithEmptyElement) == nil {
			t.Fatal("should return an empty slice element error")
		}
	})
}

func TestCheckNotEqualError(t *testing.T) {
	expected := 10

	t.Run("should return an error for slice with empty element", func(t *testing.T) {
		if CheckNotEqualError(1+expected, expected) == nil {
			t.Fatal("should return a not equal error")
		}
	})
}
