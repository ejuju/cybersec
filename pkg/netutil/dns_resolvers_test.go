package netutil

import (
	"testing"

	"github.com/ejuju/cybersec/internal/testutil"
)

func TestSampleDNSResolverAddresses(t *testing.T) {
	t.Parallel()

	t.Run("should not be empty or contain empty addresses", func(t *testing.T) {
		testutil.FailOnError(t,
			testutil.CheckZeroLengthError[string, *any](SampleDNSResolverAddresses),
			testutil.CheckEmptyStringInSliceError(SampleDNSResolverAddresses),
		)
	})
}
