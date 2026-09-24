package internalutils

import (
	"testing"
)

func TestIsWellFormedUlidString(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		isValid bool
	}{
		{
			name:    "valid_ulid",
			input:   "01GRC27AM72M4SGK4VBHF3DY0F",
			isValid: true,
		},
		{
			name:    "invalid_symbols",
			input:   "01GRC27AM72M4S$^4VBHF3DY0F",
			isValid: false,
		},
		{
			name:    "not_start_in_0-7",
			input:   "A1GRC27AM72M4SGK4VBHF3DY0F",
			isValid: false,
		},
		{
			name:    "too_long",
			input:   "01GRC27AM72M4SGK4VBHF3DY0FA",
			isValid: false,
		},
		{
			name:    "too_short",
			input:   "01GRC27AM72M4SGK4VBHF3DY0",
			isValid: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			isValid := IsWellFormedUlidString(test.input)
			if isValid != test.isValid {
				t.Errorf("Expect %s to be valid %v actual %v", test.input, test.isValid, isValid)
			}
		})
	}

}

func BenchmarkIsWellFormedUlidString(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		IsWellFormedUlidString("01GRC27AM72M4SGK4VBHF3DY0F")
	}
}
