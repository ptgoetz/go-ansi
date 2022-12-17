package main

import (
	"github.com/ptgoetz/go-ansi/pkg/ansi"
	"github.com/ptgoetz/go-ansi/pkg/ansi/color"
	"github.com/ptgoetz/go-ansi/pkg/ansi/component"
	"os"
)

func main() {
	console := ansi.NewConsole(os.Stdout)

	panel := component.NewPanelWithLabel("ERROR", " ACCESS DENIED! ", -1, 5)

	panel.BorderColor = color.Red
	panel.Render(console)

	panel.BorderColor = color.Yellow
	panel.Label = "WARNING"
	panel.Message = "I'm sorry Dave. I'm afraid I can't do that."
	panel.Render(console)

	panel.BorderColor = color.Green
	panel.Message = "That would be cool."
	panel.Label = "INFO"
	panel.Render(console)
}
