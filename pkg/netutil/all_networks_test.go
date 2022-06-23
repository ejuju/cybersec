package netutil

import (
	"testing"

	"github.com/ejuju/cybersec/internal/testutil"
)

func TestAllNetworks(t *testing.T) {
	t.Parallel()
	t.Run("should not be empty", func(t *testing.T) {
		testutil.Check(t, testutil.CheckZeroLengthError[struct{}, string](AllNetworks))
	})
}
