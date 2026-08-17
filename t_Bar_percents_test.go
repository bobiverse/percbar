package percbar

import (
	"strings"
	"testing"
)

// A negative value should not drag the total (and other sectors'
// percentages) below what makes sense for a progress bar. Spec: bar.sum is
// the sum of magnitudes (|count|), so -50 and 100 -> sum 150, and no
// sector's percent should ever fall outside [0, 100].
func TestNegativeValueSectorPercents(t *testing.T) {
	bar := New(map[string]int{"a": -50, "b": 100})

	if bar.sum != 150 {
		t.Errorf("expected bar.sum = 150 (sum of magnitudes), got %v", bar.sum)
	}

	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("String() panicked: %v", r)
		}
	}()

	out := bar.String()

	if !strings.Contains(out, "150 in total") {
		t.Errorf("expected header to show 150 in total, got: %q", out)
	}

	for _, s := range bar.sectors {
		if s.percents < 0 || s.percents > 100 {
			t.Errorf("sector %q percents out of [0,100] range: %d", s.label, s.percents)
		}
	}
}

// An all-zero input makes bar.sum == 0, so percent calculations must avoid NaN/Inf and String() must not panic.
func TestZeroSumDoesNotPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("String() panicked on zero-sum input: %v", r)
		}
	}()

	bar := New(map[string]int{"a": 0, "b": 0})
	_ = bar.String()
}
