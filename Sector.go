package percbar

import (
	"github.com/fatih/color"
)

// sector for percentage bar
type sector struct {
	label string
	count float64

	// calculated
	percents int
	char     rune
	color    *color.Color
}

func newSector(label string, count float64) *sector {
	return &sector{
		label: label,
		count: count,
	}
}
