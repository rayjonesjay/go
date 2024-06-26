package colors

// Contains ansi color codes
const (
	RESET   = "\033[0m"
	BLACK   = "\033[30m"
	RED     = "\033[31m"
	GREEN   = "\033[32m"
	YELLOW  = "\033[33m"
	BLUE    = "\033[34m"
	MAGENTA = "\033[35m"
	CYAN    = "\033[36m"
	WHITE   = "\033[37m"
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
