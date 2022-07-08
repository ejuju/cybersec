package netutil

import (
	"testing"

	"github.com/ejuju/cybersec/internal/testutil"
)

func TestAllowedNetworkNames(t *testing.T) {
	t.Parallel()
	t.Run("should not be empty or nil", func(t *testing.T) {
		testutil.FailOnError(
			t,
			testutil.CheckZeroLengthError[struct{}, string](AllowedNetworkNames),
		)
	})
}
