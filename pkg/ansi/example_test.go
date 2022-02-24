package ansi

import (
	"fmt"
	"os"
)

func ExampleFormat() {
	s := FormatAndReset("fg=default, bg=default", ColorDefault, ColorBgDefault)
	fmt.Fprintln(os.Stderr, s)
	// Output:
}

func ExampleRGBForeground() {

	s := RGBForeground("0, 255, 255", true, 0, 255, 255)
	fmt.Fprintln(os.Stderr, s)

	s = RGBForeground("255, 0, 255", true, 255, 0, 255)
	fmt.Fprintln(os.Stderr, s)

	s = RGBForeground("255, 255, 0", true, 255, 255, 0)
	fmt.Fprintln(os.Stderr, s)

	s = RGBForeground("0, 0, 255", true, 0, 0, 255)
	fmt.Fprintln(os.Stderr, s)

	s = RGBForeground("0, 255, 0", true, 0, 255, 0)
	fmt.Fprintln(os.Stderr, s)

	s = RGBForeground("255, 0, 0", true, 255, 0, 0)
	fmt.Fprintln(os.Stderr, s)

	s = RGBForeground("0, 0, 0", true, 0, 0, 0)
	fmt.Fprintln(os.Stderr, s)

	s = RGBForeground("255, 255,255", true, 255, 255, 255)
	fmt.Fprintln(os.Stderr, s)

	// Output:
}

func ExampleRGBBackground() {

	s := RGBBackground("0, 255, 255", true, 0, 255, 255)
	fmt.Fprintln(os.Stderr, s)

	s = RGBBackground("255, 0, 255", true, 255, 0, 255)
	fmt.Fprintln(os.Stderr, s)

	s = RGBBackground("255, 255, 0", true, 255, 255, 0)
	fmt.Fprintln(os.Stderr, s)

	s = RGBBackground("0, 0, 255", true, 0, 0, 255)
	fmt.Fprintln(os.Stderr, s)

	s = RGBBackground("0, 255, 0", true, 0, 255, 0)
	fmt.Fprintln(os.Stderr, s)

	s = RGBBackground("255, 0, 0", true, 255, 0, 0)
	fmt.Fprintln(os.Stderr, s)

	s = RGBBackground("0, 0, 0", true, 0, 0, 0)
	fmt.Fprintln(os.Stderr, s)

	s = RGBBackground("255, 255,255", true, 255, 255, 255)
	fmt.Fprintln(os.Stderr, s)

	// Output:
}

func ExampleFormatAndReset() {
	// Explicitly write to os.Stderr, so no output is expected.
	s := FormatAndReset("fg=default, bg=default", ColorDefault, ColorBgDefault)
	fmt.Fprintln(os.Stderr, s)

	s = FormatAndReset("fg=yellow, bg=default", ColorYellow, ColorBgDefault)
	fmt.Fprintln(os.Stderr, s)

	s = FormatAndReset("fg=red, bg=default", ColorRed, ColorBgDefault)
	fmt.Fprintln(os.Stderr, s)

	s = FormatAndReset("fg=blue, bg=BrightWhite", ColorBlue, ColorBgBrightWhite)
	fmt.Fprintln(os.Stderr, s)

	s = FormatAndReset("fg=blue, bg=Black", ColorBlue, ColorBgBlack)
	fmt.Fprintln(os.Stderr, s)

	s = FormatAndReset("fg=blue, bg=BrightBlack", ColorGreen, ColorBgBrightBlack)
	fmt.Fprintln(os.Stderr, s)

	s = FormatAndReset("fg=BrightGreen, bg=BrightRed", ColorBrightGreen, ColorBgBrightRed)
	fmt.Fprintln(os.Stderr, s)

	// Output:
}

func ExampleMoveCursor() {
	fmt.Fprint(os.Stderr, "Expecting 'foo_baz': ")
	fmt.Fprint(os.Stderr, "foo_bar")
	fmt.Fprint(os.Stderr, MoveCursor(CursorLeft, 3))
	fmt.Fprintln(os.Stderr, "baz")

	// Output:
}
