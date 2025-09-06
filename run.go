package assert

import (
	"testing"
	"testing/synctest"
)

func RunTest(t *testing.T, name string, f func(t *testing.T)) bool {
	t.Helper()
	return t.Run(name, func(t *testing.T) {
		t.Helper()
		t.Parallel()
		synctest.Test(t, f)
	})
}
