package colors

// import "errors" // Reset terminal color to default

const (
	Red   = "\033[31m"
	Reset = "\033[0m"
)

// Color defines a structure for RGB color representation.
// Each color component (R, G, B) is represented by an uint8,
// which corresponds to a value between 0 and 255.
type Color struct {
	// Red represents the red component of the color.
	Red uint8

	// Green represents the green component of the color.
	Green uint8

	// Blue represents the blue component of the color.
	Blue uint8
}
