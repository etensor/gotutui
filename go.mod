module github.com/etensor/gotutui

go 1.16

require (
	github.com/achannarasappa/term-grid v0.2.4 // indirect
	github.com/charmbracelet/bubbletea v0.14.1
	github.com/charmbracelet/lipgloss v0.3.0
)

// replace github.com/charmbracelet/lipgloss => ../../Documentos/BubbleTUIgo/lipgloss
replace github.com/etensor/gotutui => ./countdown
