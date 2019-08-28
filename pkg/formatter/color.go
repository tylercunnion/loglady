package formatter

import "github.com/logrusorgru/aurora"

func colorValue(value string) aurora.Color {
	switch value {
	case "black":
		return aurora.BlackFg
	case "red":
		return aurora.RedFg
	case "green":
		return aurora.GreenFg
	case "yellow":
		return aurora.YellowFg
	case "blue":
		return aurora.BlueFg
	case "magenta":
		return aurora.MagentaFg
	case "cyan":
		return aurora.CyanFg
	case "white":
		return aurora.WhiteFg
	case "brblack":
		return aurora.BlackFg | aurora.BrightFg
	case "brred":
		return aurora.RedFg | aurora.BrightFg
	case "brgreen":
		return aurora.GreenFg | aurora.BrightFg
	case "bryellow":
		return aurora.YellowFg | aurora.BrightFg
	case "brblue":
		return aurora.BlueFg | aurora.BrightFg
	case "brmagenta":
		return aurora.MagentaFg | aurora.BrightFg
	case "brcyan":
		return aurora.CyanFg | aurora.BrightFg
	case "brwhite":
		return aurora.WhiteFg | aurora.BrightFg
	}
	return 0
}
