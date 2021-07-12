package countdown

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

var style = lipgloss.NewStyle().
    Bold(true).
    Foreground(lipgloss.Color("#FAFAFA")).
    Background(lipgloss.Color("#7D56F4")).
    Blink(true).
    Underline(true).
    PaddingTop(2).
    PaddingLeft(4).
    MarginLeft(4).
    Width(20)

func Alive() {
	fmt.Println(style.Render("Hello World"))
}

