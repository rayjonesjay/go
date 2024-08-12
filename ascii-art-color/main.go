package main

import (
	"os"

	"ascii/colors"
	"ascii/data"
	"ascii/flags"
	"ascii/fmtx"
	"ascii/help"
	"ascii/output"
)

func main() {
	mArgs := os.Args[1:]
	if len(mArgs) == 0 {
		return
	}

	commandArgs := flags.ParseFlags(mArgs)
	if commandArgs.Text == "" {
		return
	}

	options := data.Options{}
	for _, flag := range commandArgs.Flags {
		switch flag.Name {
		case "color":
			rgbColor := colors.CheckColorModel(flag.Value)
			options.ColorFlags = append(
				options.ColorFlags, data.ColorInfo{
					Color:  rgbColor,
					Substr: flag.ExtraValue,
				},
			)
		case "align":
			if !contains(flag.Value) {
				fmtx.FatalErrorf("invalid alignment: %q\n", flag.Value)
			}
			options.Align = flag.Value
		case "output":
			options.Output = flag.Value
		default:
			fmtx.Errorf("unrecognized option %q\n", "--"+flag.Name)
			help.PrintUsage()
		}
	}

	draws := data.DrawInfo{
		Text:    commandArgs.Text,
		Style:   commandArgs.Banner,
		Options: options,
	}
	output.Draw(draws, options.Output)
}

func contains(s string) bool {
	for _, flag := range []string{"left", "right", "center", "justify"} {
		if flag == s {
			return true
		}
	}
	return false
}
