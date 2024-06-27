package colors

import "errors" // Reset terminal color to default

const (
	Red   = "\033[31m"
	Reset = "\033[0m"
)

// Color defines a structure for RGB color representation.
// Each color component (R, G, B) is represented by an uint8,
// which corresponds to a value between 0 and 255.
type Color struct {
	// R represents the red component of the color.
	R uint8

	// G represents the green component of the color.
	G uint8

	// B represents the blue component of the color.
	B uint8
}

// ParseColor attempts to parse the color defined in the given string.
// Returns the defined color if formatted correctly, with no error
func ParseColor(s string) (Color, error) {
	switch s {
	case "red":
		return Color{255, 0, 0}, nil
	case "green":
		return Color{0, 255, 0}, nil
	case "blue":
		return Color{0, 0, 255}, nil
	}

	return Color{}, errors.New("invalid color: " + s)
}
