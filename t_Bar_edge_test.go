package percbar

import (
	"strings"
	"testing"
)

// An empty input map means zero sectors and bar.sum == 0; String() must not panic.
func TestEmptyMapDoesNotPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("String() panicked on empty map: %v", r)
		}
	}()

	bar := New(map[string]int{})

	if len(bar.sectors) != 0 {
		t.Errorf("expected 0 sectors, got %d", len(bar.sectors))
	}
	if bar.sum != 0 {
		t.Errorf("expected sum 0, got %v", bar.sum)
	}
	_ = bar.String()
}

// HaveFooter: false must produce no footer line below the bar.
func TestHaveFooterFalseOmitsFooter(t *testing.T) {
	opt := *OptionsColorBlind
	opt.HaveFooter = false

	bar := New(map[string]int{"a": 1}).SetOptions(&opt)
	out := bar.String()

	if strings.Contains(out, "a 100%") {
		t.Errorf("expected no footer line, got:\n%s", out)
	}
}

// Header: "" must produce no header line above the bar.
func TestEmptyHeaderOmitsHeaderLine(t *testing.T) {
	opt := *OptionsColorBlind
	opt.Header = ""

	bar := New(map[string]int{"a": 1}).SetOptions(&opt)
	out := bar.String()

	if strings.Contains(out, "in total") {
		t.Errorf("expected no header line, got:\n%s", out)
	}
}

// When a single Chars rune must serve more sectors than there are runes, it
// wraps around via i%len(chars) instead of panicking or running out of chars.
func TestCharsWrapAroundForManySectors(t *testing.T) {
	opt := *OptionsColorBlind
	opt.Chars = "AB"

	bar := New(map[string]int{"a": 10, "b": 10, "c": 10, "d": 10, "e": 10}).SetOptions(&opt)

	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("String() panicked with fewer chars than sectors: %v", r)
		}
	}()
	_ = bar.String()
}
