package main

import (
	"fmt"

	. "github.com/achannarasappa/term-grid"
	"github.com/charmbracelet/lipgloss"
	"github.com/etensor/gotutui/countdown"
)

var style = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#101D07")).
	Background(lipgloss.Color("#ADC0D8")).
	PaddingTop(2).
	PaddingLeft(4)

var stilacho = lipgloss.NewStyle().
	Italic(true).
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#101D07")).
	PaddingTop(2).
	PaddingLeft(4)

func main() {

	out := Render(
		Grid{
			GutterVertical:   2,
			GutterHorizontal: 3,
			Rows: []Row{
				{
					Width: 50,
					Cells: []Cell{
						{Width: 20, Text: stilacho.Render("Cuenti"), Overflow: WrapWord, Align: Right},
						{Width: 20, Text: style.Render("Chimba"), Overflow: WrapWord},
					},
				},
			},
		})

	fmt.Print(out + "\n\n")
	//fmt.Println(stilacho.Render("Cuenti"))
	fmt.Println(style.Render("Hello there."))
	countdown.Main_TUI()
	countdown.Alive()
	mainss()
	countdown.Countdown()
}
