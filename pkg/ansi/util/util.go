package util

import (
	"golang.org/x/sys/unix"
	"syscall"
)

// IsTTY detects if the current context is a TTY by calling GetTTYDimensions.
// If an error is returned, or either rows or cols is zero, it's assumed that it's not a TTY.
func IsTTY() bool {
	rows, cols, err := GetTTYDimensions()
	return err == nil && rows != 0 && cols != 0
}

// GetTTYDimensions returns the number of rows, and columns of the current TTY, or an error if the current
// context is not a TTY.
func GetTTYDimensions() (rows uint16, cols uint16, err error) {
	ws, err := unix.IoctlGetWinsize(syscall.Stdout, unix.TIOCGWINSZ)
	if err != nil {
		return 0, 0, err
	}
	return ws.Row, ws.Col, nil
}
