package main

import (
	"fmt"
	"github.com/ptgoetz/go-ansi/pkg/ansi/util"
)

func main() {
	fmt.Println("IsTerminal: ", util.IsTerminal())
	if util.IsAnsiTerminal() {
		rows, cols, _ := util.GetTerminalDimensions()
		fmt.Printf("TTY detected. Rows: %d, Columns: %d\n", rows, cols)
	} else {
		fmt.Println("Not a TTY. Formatting and cursor navigation disabled.")
	}
}
