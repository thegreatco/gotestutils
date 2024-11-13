package gotestutils

import (
	"os"
	"testing"
)

type test struct {
	requiresRoot      bool
	requiresBoardType BoardType
}

func Test() *test {
	return &test{}
}

func (t *test) RequiresRoot() *test {
	t.requiresRoot = true
	return t
}

func (t *test) RequiresBoardType(boardType BoardType) *test {
	t.requiresBoardType = boardType
	return t
}

func (t *test) ShouldSkip(test *testing.T) {
	if !IsBoardType(t.requiresBoardType) {
		test.Skipf("Test requires board type %v", t.requiresBoardType)
	}
	if t.requiresRoot && !IsRoot() {
		test.Skip("Test requires root")
	}
}

func IsRoot() bool {
	return os.Geteuid() == 0
}
