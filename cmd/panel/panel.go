package main

import (
	"github.com/ptgoetz/go-ansi/pkg/ansi"
	"github.com/ptgoetz/go-ansi/pkg/ansi/component"
	"os"
)

func main() {
	//component.DemoPanel
	console := ansi.NewConsole(os.Stdout)

	panel := component.NewPanelWithLabel("ERROR", "ACCESS DENIED, FARTKNOCKER!", -2, 10)

	panel.Render(console)
}
