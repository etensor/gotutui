package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/etensor/gotutui/countdown"
)

var style = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#101D07")).
	Background(lipgloss.Color("#ADC0D8")).
	PaddingTop(2).
	PaddingLeft(4).
	Width(22)

func main() {
	fmt.Println(style.Render("Hello there."))
	mainss()
	countdown.Countdown()
}
