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

func (sect *sector) isProcessed() bool {
	return sect.char != 0 // && sect.percents > 0
}

func newSector(label string, count float64) *sector {
	return &sector{
		label: label,
		count: count,
	}
}
