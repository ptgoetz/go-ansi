package main

import (
	"github.com/ptgoetz/go-ansi/pkg/ansi"
	"github.com/ptgoetz/go-ansi/pkg/ansi/color"
	"github.com/ptgoetz/go-ansi/pkg/ansi/component"
	"os"
)

func main() {
	c := ansi.NewConsole(os.Stdout)

	panel := component.NewPanelWithLabel("FATAL", "ACCESS DENIED!", -1, 5)

	panel.BorderColor = color.Red
	panel.Render(c)

	panel.BorderColor = color.Orange
	panel.Label = "ERROR"
	panel.Message = "I'm sorry Dave. I'm afraid I can't do that."
	panel.Render(c)

	panel.BorderColor = color.Yellow
	panel.Label = "WARNING"
	panel.Message = "The beatings will continue until morale improves."
	panel.Render(c)

	panel.BorderColor = color.CornflowerBlue
	panel.Message = "It would be cool"
	panel.Label = "INFO"
	panel.Render(c)

	panel.BorderColor = color.DarkGreen
	panel.Message = "If we killed that frog"
	panel.Label = "DEBUG"
	panel.Render(c)

	panel.BorderColor = color.DimGray
	panel.Message = "It won't croak again"
	panel.Label = "TRACE"
	panel.Render(c)
}
