package main

import (
	"fmt"
	"github.com/ptgoetz/go-ansi/pkg/ansi"
	"os"
	"time"
)

func main() {
	demoProgressPercent()
}

func demoProgressPercent() {
	w := ansi.NewConsole(os.Stderr)
	backspaceLen := -1
	// set the prompt style and color, use default background
	//w.SetStyle(ansi.StyleBold)
	w.SetForeground(255, 0, 0)
	w.Print("Progress: ")

	// Reset style to defaults
	w.ResetStyle()

	w.SetForeground(255, 255, 255)
	for i := 0; i <= 100; i++ {
		s := fmt.Sprintf("%d", i) + "%"
		w.MoveCursor(ansi.CursorLeft, backspaceLen)
		w.Print(s)
		backspaceLen = len(s)

		time.Sleep(time.Millisecond * 50)
	}
	w.ResetStyle()
	w.Print("\n")

	//w.SetStyle(ansi.StyleUnderline)
	w.SetStyle(ansi.ColorGreen)
	w.Print("Foo.\n")
	w.ResetStyle()
	w.Print("Bar.")
	w.Print("\n")
}
