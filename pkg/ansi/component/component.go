package component

import (
	"fmt"
	ansi2 "github.com/ptgoetz/go-ansi/pkg/ansi"
	"os"
	"time"
)

func demoProgressPercent() {
	w := ansi2.NewWriter(os.Stderr)
	backspaceLen := -1
	// set the prompt style and color, use default background
	w.SetStyle(ansi2.StyleUnderline)
	w.SetForeground(255/2, 255/2, 255/2)
	w.Print("Progress: ")

	// Reset style to defaults
	w.ResetStyle()

	w.SetForeground(255, 255, 255)
	for i := 0; i <= 100; i++ {
		s := fmt.Sprintf("%d", i) + "%"
		w.MoveCursor(ansi2.CursorLeft, backspaceLen)
		w.Print(s)
		backspaceLen = len(s)

		time.Sleep(time.Millisecond * 50)
	}
	w.ResetStyle()
	w.Print("\n")

	//w.SetStyle(ansi.StyleUnderline)
	w.SetStyle(ansi2.ColorGreen)
	w.Print("Dog Balls.\n")
	w.ResetStyle()
	w.Print("Cat Piss.")
	w.Print("\n")
}
