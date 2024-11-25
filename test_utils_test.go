package gotestutils

import (
	"testing"

	"github.com/thegreatco/sbcidentify/raspberrypi"
)

func TestShouldSkip(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() *test
		expected bool
		msg      string
	}{
		{
			name: "Requires SBC but no SBC present",
			setup: func() *test {
				return Test().RequiresSbc()
			},
			expected: true,
			msg:      "Test requires physical SBC",
		},
		{
			name: "Requires specific board type but different board type present",
			setup: func() *test {
				return Test().RequiresBoardType(raspberrypi.RaspberryPi3B)
			},
			expected: true,
			msg:      "Test requires board type",
		},
		{
			name: "Requires root but not running as root",
			setup: func() *test {
				return Test().RequiresRoot()
			},
			expected: true,
			msg:      "Test requires root",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockTest := &testing.T{}
			tt.setup().ShouldSkip(mockTest)
			if mockTest.Skipped() != tt.expected {
				t.Errorf("expected skipped to be %v, got %v", tt.expected, mockTest.Skipped())
			}
		})
	}
}
