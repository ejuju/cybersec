package netutil

import (
	"testing"

	"github.com/ejuju/cybersec/internal/testutil"
)

func TestCommonPorts(t *testing.T) {
	t.Parallel()
	t.Run("should not be empty", func(t *testing.T) {
		testutil.FailOnError(
			t,
			testutil.CheckZeroLengthError[string, int](CommonPorts),
		)
	})
}
