package netutil

import (
	"math/rand"
	"testing"
)

func TestRandUserAgent(t *testing.T) {
	rand.Seed(0)

	t.Run("should be able to return a random user-agent from the default sample", func(t *testing.T) {
		if RandUserAgent(nil) == "" {
			t.Fatal("user agent should not be empty")
		}
	})

	t.Run("should not return the same user-agent twice in a row", func(t *testing.T) {
		if RandUserAgent(nil) == RandUserAgent(nil) {
			t.Fatal("user agent should not return the same string twice for rand.Seed(0)")
		}
	})

	t.Run("should be able to use custom user-agent list", func(t *testing.T) {
		list := []string{"a", "b"} // relies on rand.seed 0
		if RandUserAgent(list) != list[0] {
			t.Fatalf("user agent should be: %#v", list[0])
		}
		if RandUserAgent(list) != list[1] {
			t.Fatalf("user agent should be: %#v", list[1])
		}
	})
}
