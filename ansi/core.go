package ansi

import (
	"fmt"
	"io"
	"strings"
)

type Writer struct {
	ioWriter io.Writer
}

func NewWriter(writer io.Writer) *Writer {
	return &Writer{
		ioWriter: writer,
	}
}

func (w *Writer) SetStyle(ansiCodes ...int) {
	codeStr := intArrayToDelimitedString(ansiCodes, ";")
	fmt.Fprint(w.ioWriter, fmt.Sprintf("%s%sm", Start, codeStr))
}

func (w *Writer) SetForeground(r int, g int, b int) {
	fmt.Fprint(w.ioWriter, RGBForeground("", false, r, g, b))
}

func (w *Writer) SetBackground(r int, g int, b int) {
	fmt.Fprint(w.ioWriter, RGBBackground("", false, r, g, b))
}

func (w *Writer) ResetStyle() {
	fmt.Fprint(w.ioWriter, Reset)
}

func (w *Writer) MoveCursor(direction CursorDirection, places int) {
	fmt.Fprint(w.ioWriter, MoveCursor(direction, places))
}

func (w *Writer) Print(s string) {
	fmt.Fprint(w.ioWriter, s)
}

func intArrayToDelimitedString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

func Format(s string, reset bool, codes ...int) string {
	codeStr := intArrayToDelimitedString(codes, ";")
	if reset {
		return fmt.Sprintf("%s%sm%s%s", Start, codeStr, s, Reset)
	} else {
		return fmt.Sprintf("%s%sm%s", Start, codeStr, s)
	}
}

func FormatAndReset(s string, codes ...int) string {
	return Format(s, true, codes...)
}

func MoveCursor(direction CursorDirection, places int) string {
	retval := fmt.Sprintf("%s%d%s", Start, places, direction)
	return retval
}

func RGBForeground(s string, reset bool, r int, g int, b int) string {
	return Format(s, reset, ColorRGB, 2, r, g, b)
}

func RGBBackground(s string, reset bool, r int, g int, b int) string {
	return Format(s, reset, ColorBgRGB, 2, r, g, b)
}
