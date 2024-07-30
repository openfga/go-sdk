package telemetry

import (
	"testing"

	"go.opentelemetry.io/otel/attribute"
)

func TestPrepareAttributes(t *testing.T) {
	m := &Metrics{}

	tests := []struct {
		name           string
		attrs          map[*Attribute]string
		expectedLength int
	}{
		{
			name: "Single attribute",
			attrs: map[*Attribute]string{
				FGAClientRequestClientID: "12345",
			},
			expectedLength: 1,
		},
		{
			name: "Multiple attributes",
			attrs: map[*Attribute]string{
				FGAClientRequestClientID: "12345",
				FGAClientRequestMethod:   "POST",
				HTTPResponseStatusCode:   "200",
				UserAgent:                "Mozilla/5.0",
			},
			expectedLength: 4,
		},
		{
			name:           "No attributes",
			attrs:          map[*Attribute]string{},
			expectedLength: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := m.PrepareAttributes(tc.attrs)

			if len(result.ToSlice()) != tc.expectedLength {
				t.Errorf("Test %s failed: expected length %d, got %d", tc.name, tc.expectedLength, len(result.ToSlice()))
			}

			for attr, val := range tc.attrs {
				if value, ok := result.Value(attribute.Key(attr.Name)); !ok || value.AsString() != val {
					t.Errorf("Test %s failed: expected attribute %s with value %s, got value %s", tc.name, attr.Name, val, value.AsString())
				}
			}
		})
	}
}
