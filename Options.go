package percbar

// Options presets
var (
	OptionsDefault = &Options{
		Chars:             "▐",
		Colors:            "green, blue, magenta, cyan, white, black, higreen, hiblue, himagenta, hicyan, hiwhite, hiblack",
		Header:            "{SUM} in total",
		HaveFooter:        true,
		AllowFooterColors: true,
	}

	OptionsColorBlind = &Options{
		Chars:             "█▐░▚▒▓▣▢",
		Header:            "{SUM} in total",
		HaveFooter:        true,
		AllowFooterColors: false,
	}

	OptionsColorBlind2 = &Options{
		Chars:             "|#=@+%*$=>",
		Header:            "{SUM} in total",
		HaveFooter:        true,
		AllowFooterColors: false,
	}

	OptionsRGB = &Options{
		Chars:             "RGB",
		Colors:            "red, green, blue",
		Header:            "{SUM} in total",
		HaveFooter:        true,
		AllowFooterColors: true,
	}

	OptionsBgColors = &Options{
		Chars:             "▐",
		Colors:            "green:green, blue:blue, magenta:magenta, cyan:cyan, white:white, black:black, higreen:higreen, hiblue:hiblue, himagenta:himagenta, hicyan:hicyan, hiwhite:hiwhite, hiblack",
		Header:            "{SUM} in total",
		HaveFooter:        true,
		AllowFooterColors: true,
	}

	OptionsDescending = &Options{
		Chars:             "▇▆▅▄▃▂▁",
		Colors:            "red, green, blue, ",
		Header:            "{SUM} in total",
		HaveFooter:        true,
		AllowFooterColors: true,
	}
)

// Options for bar display/logic
type Options struct {
	Chars             string
	Colors            string
	Header            string
	HaveFooter        bool
	AllowFooterColors bool
}
