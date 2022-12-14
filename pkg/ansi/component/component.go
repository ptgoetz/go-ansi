package component

import (
	"fmt"
	ansi2 "github.com/ptgoetz/go-ansi/pkg/ansi"
	"os"
	"strings"
	"time"
)

func DemoProgressPercent() {
	w := ansi2.NewConsole(os.Stderr)
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

func DemoPanel() {
	//w := ansi2.NewConsole(os.Stderr)
	fmt.Println(CORNER_UPPER_LEFT + HORIZONTAL_LINE + "LABEL" + strings.Repeat(HORIZONTAL_LINE, 5) + CORNER_UPPER_RIGHT)
	fmt.Println(VERTICAL_LINE + "  foo      " + VERTICAL_LINE)
	fmt.Println(CORNER_LOWER_LEFT + strings.Repeat(HORIZONTAL_LINE, 11) + CORNER_LOWER_RIGHT)
}
