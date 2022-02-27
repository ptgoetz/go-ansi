package util

import (
	"golang.org/x/sys/unix"
	"os"
	"syscall"
)

func IsTerminal() bool {
	fileInfo, _ := os.Stdout.Stat()
	return (fileInfo.Mode() & os.ModeCharDevice) != 0
}

// IsAnsiTerminal detects if the current context is a TTY by calling GetTerminalDimensions.
// If an error is returned, or either rows or cols is zero, it's assumed that it's not a TTY.
func IsAnsiTerminal() bool {
	rows, cols, err := GetTerminalDimensions()
	return err == nil && rows != 0 && cols != 0
}

// GetTerminalDimensions returns the number of rows, and columns of the current TTY, or an error if the current
// context is not a TTY.
func GetTerminalDimensions() (rows uint16, cols uint16, err error) {
	ws, err := unix.IoctlGetWinsize(syscall.Stdout, unix.TIOCGWINSZ)
	if err != nil {
		return 0, 0, err
	}
	return ws.Row, ws.Col, nil
}
