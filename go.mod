module github.com/etensor/gotutui

go 1.16

require (
	github.com/charmbracelet/bubbletea v0.14.1
	github.com/charmbracelet/lipgloss v0.3.0
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	golang.org/x/term v0.0.0-20210615171337-6886f2dfbf5b // indirect
)

// replace github.com/charmbracelet/lipgloss => ../../Documentos/BubbleTUIgo/lipgloss
replace github.com/etensor/gotutui => ./countdown
