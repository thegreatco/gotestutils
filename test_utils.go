package gotestutils

import (
	"os"
	"testing"
)

type test struct {
	requiresRoot      bool
	requiredBoardType *BoardType
	requiresSbc       bool
}

func Test() *test {
	return &test{}
}

func (t *test) RequiresRoot() *test {
	t.requiresRoot = true
	return t
}

func (t *test) RequiresBoardType(boardType BoardType) *test {
	t.requiredBoardType = &boardType
	return t
}

func (t *test) RequiresSbc() *test {
	t.requiresSbc = true
	return t
}

func (t *test) ShouldSkip(test *testing.T) {
	if t.requiresSbc {
		boardType, err := GetBoardType()
		if err != nil {
			test.Error(err)
		}
		if boardType == BoardTypeUnknown {
			test.Skipf("Test requires physical SBC")
		}
	}

	if t.requiredBoardType != nil && !IsBoardType(*t.requiredBoardType) {
		test.Skipf("Test requires board type %v", t.requiredBoardType)
	}
	if t.requiresRoot && !IsRoot() {
		test.Skip("Test requires root")
	}
}

func IsRoot() bool {
	return os.Geteuid() == 0
}
