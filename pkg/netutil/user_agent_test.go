package netutil

import (
	"math/rand"
	"testing"

	"github.com/ejuju/cybersec/internal/testutil"
)

func TestSampleHTTPUserAgents(t *testing.T) {
	t.Run("should not be nil, of length 0, or contain empty strings", func(t *testing.T) {
		testutil.Check(t,
			testutil.CheckZeroLengthError[string, *any](SampleHTTPUserAgents),
			testutil.CheckEmptyStringInSliceError(SampleHTTPUserAgents),
		)
	})
}

func TestRandUserAgent(t *testing.T) {
	t.Run("should be able to return a random user-agent from the default sample", func(t *testing.T) {
		if RandUserAgent(nil) == "" {
			t.Fatal("user agent should not be empty")
		}
	})

	t.Run("should not return the same user-agent twice in a row", func(t *testing.T) {
		rand.Seed(0)
		if RandUserAgent(nil) == RandUserAgent(nil) {
			t.Fatal("user agent should not return the same string twice for rand.Seed(0)")
		}
	})

	t.Run("should be able to use custom user-agent list", func(t *testing.T) {
		list := []string{"a"}
		if RandUserAgent(list) != list[0] {
			t.Fatalf("user agent should be: %#v", list[0])
		}
	})
}
