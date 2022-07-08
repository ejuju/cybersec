package testutil

import "testing"

func TestCheckNotEqualError(t *testing.T) {
	t.Parallel()

	expected := 10

	t.Run("should return an error for slice with empty element", func(t *testing.T) {
		if CheckNotEqualError(1+expected, expected) == nil {
			t.Fatal("should return a not equal error")
		}
		if CheckNotEqualError(expected, expected) != nil {
			t.Fatal("should NOT return a not equal error")
		}
	})
}

func TestCheckNilPointerError(t *testing.T) {
	t.Parallel()

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
	t.Parallel()

	var zeroLengthSlice = []any{}
	var validSlice = []int{0}

	t.Run("should return an error for slice of length 0", func(t *testing.T) {
		if CheckZeroLengthError[any, *any](zeroLengthSlice) == nil {
			t.Fatal("should return a zero-length slice error")
		}
		if CheckZeroLengthError[int, string](validSlice) != nil {
			t.Fatal("should return a zero-length slice error")
		}
	})
}

func TestCheckEmptyStringInSliceError(t *testing.T) {
	t.Parallel()

	var sliceWithEmptyElement = []string{"a", ""}

	t.Run("should return an error for slice with empty element", func(t *testing.T) {
		if CheckEmptyStringInSliceError(sliceWithEmptyElement) == nil {
			t.Fatal("should return an empty slice element error")
		}
	})
}
