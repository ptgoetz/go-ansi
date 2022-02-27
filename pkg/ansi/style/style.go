package style

import (
	"github.com/ptgoetz/go-ansi/pkg/ansi"
)

// Set represents a set of Style flags that can be toggled
type Set struct {
	none          bool
	bold          bool
	dim           bool
	italic        bool
	underline     bool
	blink         bool
	inverse       bool
	strikethrough bool
}

// styleCodes returns a slice of ansi style codes
func (s *Set) styleCodes() []int {
	var codes []int
	if s.bold {
		codes = append(codes, ansi.StyleBold)
	}
	if s.dim {
		codes = append(codes, ansi.StyleDim)
	}
	if s.italic {
		codes = append(codes, ansi.StyleItalic)
	}
	if s.underline {
		codes = append(codes, ansi.StyleUnderline)
	}
	if s.blink {
		codes = append(codes, ansi.StyleBlink)
	}
	if s.inverse {
		codes = append(codes, ansi.StyleInverse)
	}
	if s.strikethrough {
		codes = append(codes, ansi.StyleBold)
	}
	if len(codes) == 0 {
		codes = append(codes, ansi.StyleNone)
	}

	return codes
}

// Style represents an ANSI TTY format. It has a foreground color, a background color, and a set of styles
type Style struct {
	styleSet        *Set
	foregroundColor string
	backgroundColor string
}

// *** COLORS ***

// DefaultStyle returns a Style pointer with all zero values
func DefaultStyle() *Style {
	return &Style{
		styleSet: &Set{},
	}
}

// Red sets the foreground color to red
func (s *Style) Red() *Style {
	return s
}

// Green sets the foreground color to green
func (s *Style) Green() *Style {
	return s
}

// Blue sets the foreground color to blue
func (s *Style) Blue() *Style {
	return s
}

// *** RGB Colors ***

// RGB sets the foreground color to the given r, g, b, values
func (s *Style) RGB(r int, g int, b int) *Style {
	return s
}

// RGBBg sets the background color to the given r, g, b, values
func (s *Style) RGBBg(r int, g int, b int) *Style {
	return s
}

// *** Styles ***

// Bold enables ANSI bold output
func (s *Style) Bold() *Style {
	s.styleSet.bold = true
	return s
}

// Dim enables ANSI dim output
func (s *Style) Dim() *Style {
	s.styleSet.dim = true
	return s
}

// Italic enables ANSI italic output
func (s *Style) Italic() *Style {
	s.styleSet.italic = true
	return s
}

// Underline enables ANSI underline output
func (s *Style) Underline() *Style {
	s.styleSet.underline = true
	return s
}

// Blink enables ANSI blink output
func (s *Style) Blink() *Style {
	s.styleSet.blink = true
	return s
}

// Inverse enables ANSI inverse output
func (s *Style) Inverse() *Style {
	s.styleSet.inverse = true
	return s
}

// Strikethrough enables ANSI strikethrough output
func (s *Style) Strikethrough() *Style {
	s.styleSet.strikethrough = true
	return s
}
