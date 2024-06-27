package terminal

import (
	"os"
	"syscall"
	"unsafe"
)

// GetTerminalSize returns the size of the terminal window, i.e, the
// number of columns and rows of text that can be written to the terminal.
// Implementation Detail:
// The Syscall function is invoked with the following parameters:
// The ioctl system call number (syscall.SYS_IOCTL). `IOCTL` stands for "input/output control"
// The file descriptor of the standard output (os.Stdout.Fd()),
// indicating that the ioctl operation is being performed on the terminal associated with standard output.
// The request code (syscall.TIOCGWINSZ), specifying that we want to get the terminal window size.
// TIOCGWINSZ stands for "Terminal Input/Output Control Get Window Size"
// A pointer to a WinSize struct (uintptr(unsafe.Pointer(ws))),
// which will be filled with the terminal's window size information.
func GetTerminalSize() (int, int, error) {
	type WinSize struct {
		Row    uint16
		Col    uint16
		XPixel uint16
		YPixel uint16
	}

	ws := &WinSize{}
	retCode, _, errno := syscall.Syscall(
		syscall.SYS_IOCTL,
		os.Stdout.Fd(),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)),
	)

	if int(retCode) == -1 {
		return 0, 0, errno
	}
	return int(ws.Col), int(ws.Row), nil
}

// IsTerminal checks if stdout is a terminal
// Implementation Detail:
// Same as that of GetTerminalSize, but this time, we use the syscall.Syscall6 function in Go,
// which is a low-level interface for making system calls that require up to six arguments.
// This time we use the request code (syscall.TCGETS) to get terminal attributes (Terminal Control GET State).
// We also pass a pointer to a Termios struct where the terminal attributes will be stored.
func IsTerminal() bool {
	var termios syscall.Termios
	_, _, err := syscall.Syscall6(
		syscall.SYS_IOCTL,
		os.Stdout.Fd(),
		uintptr(syscall.TCGETS),
		uintptr(unsafe.Pointer(&termios)),
		0,
		0,
		0,
	)
	return err == 0
}
