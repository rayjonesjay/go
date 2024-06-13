package sound

import "testing"

func TestBeep(t *testing.T) {
	// Listen for the beep sound.
	//We expect this to fail on a non-Linux OS, or a non-free-desktop compatible OS
	Beep()
}
