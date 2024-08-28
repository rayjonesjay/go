package colors

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"ascii/fmtx"
)

// UpperCase for exporting
var (
	ColorsMap = map[string]Color{
		"red":       {Red: 255, Green: 0, Blue: 0},
		"green":     {Red: 0, Green: 255, Blue: 0},
		"blue":      {Red: 0, Green: 0, Blue: 255},
		"yellow":    {Red: 255, Green: 255, Blue: 0},
		"cyan":      {Red: 0, Green: 255, Blue: 255},
		"magenta":   {Red: 255, Green: 0, Blue: 255},
		"black":     {Red: 0, Green: 0, Blue: 0},
		"white":     {Red: 255, Green: 255, Blue: 255},
		"gray":      {Red: 128, Green: 128, Blue: 128},
		"orange":    {Red: 255, Green: 165, Blue: 0},
		"purple":    {Red: 128, Green: 0, Blue: 128},
		"brown":     {Red: 165, Green: 42, Blue: 42},
		"pink":      {Red: 255, Green: 192, Blue: 203},
		"lightblue": {Red: 173, Green: 216, Blue: 230},
		"navy":      {Red: 0, Green: 0, Blue: 128},
		"lime":      {Red: 0, Green: 255, Blue: 0},
		"olive":     {Red: 128, Green: 128, Blue: 0},
		"teal":      {Red: 0, Green: 128, Blue: 128},
		"maroon":    {Red: 128, Green: 0, Blue: 0},
		"violet":    {Red: 238, Green: 130, Blue: 238},
		"gold":      {Red: 255, Green: 215, Blue: 0},
		"silver":    {Red: 192, Green: 192, Blue: 192},
		"indigo":    {Red: 75, Green: 0, Blue: 130},
		"coral":     {Red: 255, Green: 127, Blue: 80},
	}
)

// ParseColor attempts to parse the color defined in the given string.
// Returns the defined color if formatted correctly, with no error
// --color=red, --color=#ff0000, --color=rgb(255, 0, 0) or --color=hsl(0, 100%, 50%)
func CheckColorModel(colorModel string) Color {
	colorModel = strings.ToLower(colorModel)

	color, ok := ColorsMap[colorModel]
	if ok {
		return color
	}

	color, ok, err := CalculateRGBFormat(colorModel)
	if err == nil {
		return color
	} else if ok {
		fmtx.FatalErrorf("invalid RGB color: %q\n%v\n", colorModel, err)
	}

	color, err = CalculateHexFormat(colorModel)
	if err == nil {
		return color
	}

	fmtx.FatalErrorf("invalid color model: %q\n", colorModel)
	return Color{}
}

// Calculate the RGB format
func CalculateRGBFormat(s string) (Color, bool, error) {
	// match 3 digits in a bracket
	pattern := `^\s*rgb\s*\(\s*(\d{1,3})\s*\,\s*(\d{1,3})\s*\,\s*(\d{1,3})\s*\)\s*$`
	re := regexp.MustCompile(pattern)

	matches := re.FindStringSubmatch(s)
	if matches != nil {
		r := matches[1]
		rInt, _ := strconv.Atoi(r)

		g := matches[2]
		gInt, _ := strconv.Atoi(g)

		b := matches[3]
		bInt, _ := strconv.Atoi(b)

		err := checkRGBRange(rInt, gInt, bInt)
		if err != nil {
			return Color{}, true, err
		}
		return Color{Red: uint8(rInt), Green: uint8(gInt), Blue: uint8(bInt)}, true, nil
	}
	return Color{}, false, errors.New("")
}

func checkRGBRange(r, g, b int) error {
	err := checkNamedColorRange(r, "Red")
	if err != nil {
		return err
	}
	err = checkNamedColorRange(g, "Green")
	if err != nil {
		return err
	}
	err = checkNamedColorRange(b, "Blue")
	if err != nil {
		return err
	}

	return nil
}

func checkNamedColorRange(code int, name string) error {
	if !(code >= 0 && code <= 255) {
		return fmt.Errorf("%s color value out of range: %d", name, code)
	}
	return nil
}

// --color=#ff0000
func CalculateHexFormat(s string) (Color, error) {
	pattern := `^#([a-fA-F0-9]{2}|[a-fA-F0-9]{4}|[a-fA-F0-9]{6})$`
	re := regexp.MustCompile(pattern)
	if re.MatchString(s) {
		s = strings.TrimPrefix(s, "#")
		var red, green, blue string = "00", "00", "00"
		if len(s) >= 2 {
			red = s[:2]
		}
		if len(s) >= 4 {
			green = s[2:4]
		}

		if len(s) >= 6 {
			blue = s[4:]
		}

		blueInt, err := hexToInt(blue)
		if err != nil {
			return Color{}, err
		}
		greenInt, err := hexToInt(green)
		if err != nil {
			return Color{}, err
		}
		redInt, err := hexToInt(red)
		if err != nil {
			return Color{}, err
		}

		return Color{Red: redInt, Green: greenInt, Blue: blueInt}, nil
	}

	return Color{}, errors.New("")
}

func hexToInt(s string) (uint8, error) {
	decimal, err := strconv.ParseInt(s, 16, 32)
	if err != nil {
		return 0, err
	}
	return uint8(decimal), nil
}
