package ansi

import (
	"fmt"
	"testing"
)

func TestAnsiFormat(t *testing.T) {
	s := FormatAndReset("Taylor", ColorRed, StyleUnderline)
	fmt.Println(s)
}
