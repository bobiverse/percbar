package percbar

import (
	"strings"

	"github.com/fatih/color"
)

func trimlow(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	return s
}

// (c) https://github.com/fatih/color/blob/main/color.go
var colormap = map[string]color.Attribute{
	"black":     color.FgBlack,
	"red":       color.FgRed,
	"green":     color.FgGreen,
	"yellow":    color.FgYellow,
	"blue":      color.FgBlue,
	"magenta":   color.FgMagenta,
	"cyan":      color.FgCyan,
	"white":     color.FgWhite,
	"hiblack":   color.FgHiBlack,
	"hired":     color.FgHiRed,
	"higreen":   color.FgHiGreen,
	"hiyellow":  color.FgHiYellow,
	"hiblue":    color.FgHiBlue,
	"himagenta": color.FgHiMagenta,
	"hicyan":    color.FgHiCyan,
	"hiwhite":   color.FgHiWhite,

	"bgblack":     color.BgBlack,
	"bgred":       color.BgRed,
	"bggreen":     color.BgGreen,
	"bgyellow":    color.BgYellow,
	"bgblue":      color.BgBlue,
	"bgmagenta":   color.BgMagenta,
	"bgcyan":      color.BgCyan,
	"bgwhite":     color.BgWhite,
	"bghiblack":   color.BgHiBlack,
	"bghired":     color.BgHiRed,
	"bghigreen":   color.BgHiGreen,
	"bghiyellow":  color.BgHiYellow,
	"bghiblue":    color.BgHiBlue,
	"bghimagenta": color.BgHiMagenta,
	"bghicyan":    color.BgHiCyan,
	"bghiwhite":   color.BgHiWhite,

	"+bold":         color.Bold,
	"+faint":        color.Faint,
	"+italic":       color.Italic,
	"+underline":    color.Underline,
	"+blinkslow":    color.BlinkSlow,
	"+blinkrapid":   color.BlinkRapid,
	"+reversevideo": color.ReverseVideo,
	"+concealed":    color.Concealed,
	"+crossedout":   color.CrossedOut,
}

func textToColors(s string) []*color.Color {
	var colors []*color.Color

	arr := strings.Split(s, ",")
	for _, c := range arr {
		colors = append(colors, asColor(c))
	}

	return colors
}

func asColor(s string) *color.Color {

	s = trimlow(s)
	arr := strings.Split(s, ":")
	for i, s := range arr {
		arr[i] = trimlow(s)
	}

	var c = color.New()

	// foreground color
	if attr, has := colormap[arr[0]]; has {
		c = c.Add(attr)
	}

	// background color or color attribute
	if len(arr) >= 2 {
		for _, attr := range arr[1:] {
			// background color
			if attr, has := colormap["bg"+attr]; has {
				c = c.Add(attr)
			}
			// attribute
			if attr, has := colormap["+"+attr]; has {
				c = c.Add(attr)
			}
		}
	}

	return c
}
