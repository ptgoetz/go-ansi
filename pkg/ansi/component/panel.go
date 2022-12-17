package component

import (
	"github.com/ptgoetz/go-ansi/pkg/ansi"
	"github.com/ptgoetz/go-ansi/pkg/ansi/color"
	"log"
	"strings"
)

const (
	CORNER_UPPER_LEFT  = "\u256d"
	CORNER_UPPER_RIGHT = "\u256e"
	CORNER_LOWER_LEFT  = "\u2570"
	CORNER_LOWER_RIGHT = "\u256f"
	HORIZONTAL_LINE    = "\u2500"
	VERTICAL_LINE      = "\u2502"
)

const (
	FIT_CONSOLE_WIDTH = -2
	FIT_MESSAGE_WIDTH = -1

	FIT_CONSOLE_HEIGHT = -1
)

type Panel struct {
	CornerUpperLeftString  string
	CornerUpperRightString string
	CornerLowerLeftString  string
	CornerLowerRightString string
	VerticalLine           string
	HorizontalLine         string
	Width                  int
	Height                 int
	BorderColor            *color.Color
	Label                  string
	Message                string
}

func NewPanel(message string, width int, height int) *Panel {
	return NewPanelWithLabel("", message, width, height)
}

func NewPanelWithLabel(label string, message string, width int, height int) *Panel {
	return &Panel{
		CornerUpperLeftString:  CORNER_UPPER_LEFT,
		CornerUpperRightString: CORNER_UPPER_RIGHT,
		CornerLowerLeftString:  CORNER_LOWER_LEFT,
		CornerLowerRightString: CORNER_LOWER_RIGHT,
		VerticalLine:           VERTICAL_LINE,
		HorizontalLine:         HORIZONTAL_LINE,
		Width:                  width,
		Height:                 height,
		BorderColor:            color.YellowGreen,
		Label:                  label,
		Message:                message,
	}
}

func (p Panel) Render(c *ansi.Console) {
	// TODO check for fit height/fit width and react accordingly
	upperLeftOffset := 2  // corner char + horizontal line char
	upperRightOffset := 1 // corner char
	lowerLeftOffset := 1  // corner char
	lowerRightOffset := 1 // corner char
	var panelWidth = p.Width
	// figure out if we're filling the width of the console or fitting message width

	if panelWidth == FIT_CONSOLE_WIDTH {
		var err error
		_, panelWidth, err = ansi.GetConoleDimension()
		panelWidth--
		if err != nil {
			log.Fatalln(err)
		}
	} else if panelWidth == FIT_MESSAGE_WIDTH {
		panelWidth = len(p.Message) + 2
	}

	labelWidth := len(p.Label)
	upperLineWidth := panelWidth - labelWidth - upperLeftOffset - upperRightOffset
	upperLine := strings.Repeat(p.HorizontalLine, upperLineWidth)
	lowerLineWidth := panelWidth - lowerLeftOffset - lowerRightOffset
	lowerLine := strings.Repeat(p.HorizontalLine, lowerLineWidth)
	// render top line
	c.SetFgColor(*p.BorderColor)
	c.Print(p.CornerUpperLeftString)
	c.Print(p.HorizontalLine)
	c.Print(p.Label)
	c.Print(upperLine)
	c.Print(p.CornerUpperRightString)
	c.Print("\n")
	// print content
	// TODO deal with text wrap, etc. We're assuming one line with no line breaks
	// pad width
	remainderWidth := panelWidth - len(p.Message) - 2 // 2 character
	remainderPad := strings.Repeat(" ", remainderWidth)
	c.Print(p.VerticalLine)
	c.ResetStyle()
	c.Print(p.Message + remainderPad)
	c.SetFgColor(*p.BorderColor)
	c.Print(p.VerticalLine)
	c.Print("\n")

	// render bottom line
	c.Print(p.CornerLowerLeftString + lowerLine + p.CornerLowerRightString)
	c.Print("\n")
	c.ResetStyle()
}
