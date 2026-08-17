package percbar

import "testing"

// bar.chars is rebuilt from options.Chars on every String() call and must
// not grow unbounded across repeated calls on the same *Bar.
func TestCharsDoNotGrowAcrossCalls(t *testing.T) {
	bar := New(map[string]int{"a": 10, "b": 10, "c": 10, "d": 10, "e": 10})

	bar.String()
	firstLen := len(bar.chars)

	for i := 0; i < 10; i++ {
		bar.String()
	}

	if len(bar.chars) != firstLen {
		t.Errorf("bar.chars grew from %d to %d entries after repeated String() calls; expected it to stay constant", firstLen, len(bar.chars))
	}
}
