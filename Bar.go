package percbar

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/fatih/color"
)

// Bar is filled with sectors
type Bar struct {
	sectors sectors
	options Options

	// processed
	chars  []rune
	colors []*color.Color
	sum    float64

	// cached
	cache string
}

// New bar as exported func
func New(values map[string]float64) *Bar {
	bar := &Bar{
		options: *OptionsDefault,
	}

	for label, count := range values {
		bar.sectors = append(bar.sectors, newSector(label, count))
		bar.sum += count
	}

	sort.Sort(bar.sectors)

	return bar
}

// SetOptions - chain func to apply options
func (bar *Bar) SetOptions(opt *Options) *Bar {
	bar.options = *opt
	bar.cache = ""
	return bar
}

// func (bar *Bar) useBlock(i int) *CharBlock {
// 	return bar.options.Blocks[i%len(bar.options.Blocks)]
// }

// String to Stringer interface
// Example: fmt.Println(bar)
func (bar *Bar) String() string {
	if bar.cache != "" {
		return bar.cache
	}

	if bar.options.Chars == "" {
		bar.options.Chars = string([]rune(OptionsDefault.Chars)[0])
	}

	// chars to runes
	for _, c := range []rune(bar.options.Chars) {
		bar.chars = append(bar.chars, c)
	}

	bar.colors = textToColors(bar.options.Colors)

	var s string
	var footer string

	// loop and make this bar
	for i, sector := range bar.sectors {

		sector.char = bar.chars[i%len(bar.chars)]
		sector.color = bar.colors[i%len(bar.colors)]
		sector.percents = int(math.Floor((sector.count / bar.sum) * 100.0))

		sectOut := strings.Repeat(string(sector.char), sector.percents)

		// if there is no colors mode (or same color) and same block used
		// use separator to know where sectors begin/end.
		useSeparator := i > 0 && sector.percents > 0
		useSeparator = useSeparator && sector.color == bar.sectors[i-1].color
		useSeparator = useSeparator && sector.char == bar.sectors[i-1].char
		if useSeparator {
			sectOut = " " + strings.Repeat(string(sector.char), sector.percents-1)
		}

		// [MyLabel 34% (234)]
		labelParts := []string{
			" " + sector.label,                   // MyLabel
			fmt.Sprintf("%d%%", sector.percents), //  34%
			fmt.Sprintf("(%.0f)", sector.count),  // (234)
		}

		// choose longest possible for display
		sectFooter := ""
		for i := len(labelParts); i >= 0; i-- {
			if label := strings.Join(labelParts[:i], " "); len(label) < sector.percents {
				sectFooter = label + strings.Repeat(" ", sector.percents-len(label))
				break
			}
		}

		if sector.color != nil {
			sectOut = sector.color.Sprintf("%s", sectOut)
		}

		if bar.options.AllowFooterColors {
			sectFooter = sector.color.Sprintf("%s", sectFooter)
		}

		s += sectOut

		if bar.options.HaveFooter {
			footer += sectFooter
		}
	}

	// prepend if necessary
	if header := bar.options.Header; header != "" {
		ssum := fmt.Sprintf("%.2f", bar.sum)
		ssum = strings.TrimSuffix(ssum, ".00")
		header = strings.ReplaceAll(header, "{SUM}", ssum)
		s = fmt.Sprintf(" %s\n%s\n", header, s)
	}

	if footer != "" {
		s += footer + "\n"
	}

	// ready for use
	return s
}
