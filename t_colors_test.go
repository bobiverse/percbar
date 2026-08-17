package percbar

import "testing"

// asColor must not panic on empty/unknown input and must still return a
// usable, non-nil *color.Color.
func TestAsColorUnknownAndEmpty(t *testing.T) {
	tests := []string{"", "notacolor", "  ", "RED"}

	for _, s := range tests {
		t.Run(s, func(t *testing.T) {
			c := asColor(s)
			if c == nil {
				t.Fatalf("asColor(%q) returned nil", s)
			}
		})
	}
}

// asColor combines a foreground color with a background color and/or
// attribute given after the ":" separator.
func TestAsColorForegroundBackgroundAttribute(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"fg only", "red"},
		{"fg + bg", "green:bggreen"},
		{"fg + attribute", "red:+bold"},
		{"fg + bg + attribute", "cyan:bgcyan:+underline"},
		{"case and space insensitive", " Red : +BOLD "},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			c := asColor(tc.input)
			if c == nil {
				t.Fatalf("asColor(%q) returned nil", tc.input)
			}
		})
	}
}

// textToColors must return one *color.Color per comma-separated entry, in order.
func TestTextToColorsCount(t *testing.T) {
	tests := []struct {
		name string
		text string
		want int
	}{
		{"single", "red", 1},
		{"three colors", "red, green, blue", 3},
		{"empty string still yields one entry", "", 1},
		{"trailing comma yields extra entry", "red, green, ", 3},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			colors := textToColors(tc.text)
			if len(colors) != tc.want {
				t.Errorf("expected %d colors, got %d", tc.want, len(colors))
			}
			for i, c := range colors {
				if c == nil {
					t.Errorf("color at index %d is nil", i)
				}
			}
		})
	}
}
