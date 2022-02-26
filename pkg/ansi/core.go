package ansi

import (
	"fmt"
	"io"
	"strings"
)

// Writer defines low-level methods for ANSI formatting to a TTY represented by an io.Writer
type Writer struct {
	ioWriter io.Writer
}

// NewWriter return a pointer to an Writer given an io.Writer reference
func NewWriter(writer io.Writer) *Writer {
	return &Writer{
		ioWriter: writer,
	}
}

// SetStyle sets the TTY's style to the given set of codes.
func (w *Writer) SetStyle(ansiCodes ...int) {
	codeStr := intArrayToDelimitedString(ansiCodes, ";")
	fmt.Fprint(w.ioWriter, fmt.Sprintf("%s%sm", Start, codeStr))
}

// SetForeground sets the TTY foreground color given RGB values
func (w *Writer) SetForeground(r int, g int, b int) {
	fmt.Fprint(w.ioWriter, RGBForeground("", false, r, g, b))
}

// SetBackground sets the TTY foreground color given RGB values
func (w *Writer) SetBackground(r int, g int, b int) {
	fmt.Fprint(w.ioWriter, RGBBackground("", false, r, g, b))
}

// ResetStyle returns the TTY styles to their defaults.
func (w *Writer) ResetStyle() {
	fmt.Fprint(w.ioWriter, Reset)
}

// MoveCursor moves the TTY cursor `places` in the given CursorDirection
func (w *Writer) MoveCursor(direction CursorDirection, places int) {
	fmt.Fprint(w.ioWriter, MoveCursor(direction, places))
}

// Print writes a string to the underlying io.Writer
func (w *Writer) Print(s string) {
	fmt.Fprint(w.ioWriter, s)
}

func intArrayToDelimitedString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

// Format returns an ANSI-formatted string representing the given codes. If reset is true,
// the returned string will be terminated with the ANSI Reset code.
func Format(s string, reset bool, codes ...int) string {
	codeStr := intArrayToDelimitedString(codes, ";")
	if reset {
		return fmt.Sprintf("%s%sm%s%s", Start, codeStr, s, Reset)
	}
	return fmt.Sprintf("%s%sm%s", Start, codeStr, s)
}

// FormatAndReset returns an ANSI-formatted string representing the given codes. The returned string
// will be terminated with the ANSI Reset code.
func FormatAndReset(s string, codes ...int) string {
	return Format(s, true, codes...)
}

// MoveCursor moves the TTY cursor `places` points in the given `direction`
func MoveCursor(direction CursorDirection, places int) string {
	retval := fmt.Sprintf("%s%d%s", Start, places, direction)
	return retval
}

// RGBForeground sets the ANSI foreground color to the given r, g, b values and returns
// the formatted value of `s` as a string.
func RGBForeground(s string, reset bool, r int, g int, b int) string {
	return Format(s, reset, ColorRGB, 2, r, g, b)
}

// RGBBackground sets the ANSI background color to the given r, g, b values and returns
// the formatted value of `s` as a string.
func RGBBackground(s string, reset bool, r int, g int, b int) string {
	return Format(s, reset, ColorBgRGB, 2, r, g, b)
}
