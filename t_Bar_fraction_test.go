package percbar

import (
	"fmt"
	"strings"
	"testing"
)

// The percent label shows one decimal only when the true percentage's
// fractional part is strictly more than .5; otherwise it shows the plain
// floored integer. Spec examples: 49.2 -> 49; 49.5 -> 49; 49.75 -> 49.8;
// 49.9 -> 49.9.
func TestPercentLabelShowsFractionPastHalf(t *testing.T) {
	tests := []struct {
		name      string
		value     float64 // sector "a"; sector "b" is 1000-value, so sum is always 1000
		wantLabel string  // expected percent text for sector "a"
	}{
		{"49.2% floors, no fraction shown", 492, "49%"},
		{"49.5% exactly half, no fraction shown", 495, "49%"},
		{"49.75% rounds to one decimal", 497.5, "49.8%"},
		{"49.9% shows one decimal", 499, "49.9%"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			bar := New(map[string]float64{"a": tc.value, "b": 1000 - tc.value}).SetOptions(OptionsColorBlind)
			out := bar.String()

			want := fmt.Sprintf(" a %s (%.0f)", tc.wantLabel, tc.value)
			if !strings.Contains(out, want) {
				t.Errorf("expected output to contain %q, got:\n%s", want, out)
			}
		})
	}
}
