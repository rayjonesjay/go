package sound

import (
	"fmt"
	"os"
	"os/exec"
)

// Beep plays a beep sound. To do this, the target system must have the program 'paplay' installed
// and that the file: '/usr/share/sounds/freedesktop/stereo/bell.oga' exists and is a valid OGG file
func Beep() {
	// Linux has the program, paplay, that plays a given audio file of the supported file formats including .oga files
	// We can, thus, play the default bell sound, preinstalled by freedesktop
	//(an open desktop standard followed by most windowing systems, including Ubuntu), with the command:
	//paplay /usr/share/sounds/freedesktop/stereo/bell.oga
	paplay := "paplay"
	bell := "/usr/share/sounds/freedesktop/stereo/bell.oga"
	cmd := exec.Command(paplay, bell)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error: Couldn't play Beep sound. "+
			"Ensure the program '%s' is installed and that the file: '%s' exists\n", paplay, bell)
		os.Exit(1)
	}
}
