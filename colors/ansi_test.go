package colors

import (
	"fmt"
	"testing"
)

// Testing the ansi codes
func TestAnsiColors(t *testing.T) {
	fmt.Printf("%sThis should be red%s\n", RED, RESET)
	fmt.Printf("%sThis should be black%s\n", BLACK, RESET)
	fmt.Printf("%sThis should be green%s\n", GREEN, RESET)
	fmt.Printf("%sThis should be yellow%s\n", YELLOW, RESET)
	fmt.Printf("%sThis should be blue%s\n", BLUE, RESET)
	fmt.Printf("%sThis should be magneta%s\n", MAGENTA, RESET)
	fmt.Printf("%sThis should be cyan%s\n", CYAN, RESET)
	fmt.Printf("%sThis should be white%s\n", WHITE, RESET)
}
