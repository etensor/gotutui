package selector

// A simple example that shows how to retrieve a value from a Bubble Tea
// program after the Bubble Tea has exited.
//
// Thanks to Treilik for this one.

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

var choices = []string{"e","pi","phi"}
var message = " e π φ digit visualization:\n\t select number:\n"


type model struct {
	cursor int
	choice chan string // channel listening until closed or sended
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			close(m.choice) // If we're quitting just close the channel.
			return m, tea.Quit

		case "enter":
			// Send the choice on the channel and exit.
			m.choice <- choices[m.cursor]
			return m, tea.Quit

		case "down", "j":
			m.cursor++
			if m.cursor >= len(choices) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(choices) - 1
			}
		}

	}

	return m, nil
}

func (m model) View() string {
	s := strings.Builder{}
	s.WriteString(message)

	for i := 0; i < len(choices); i++ {
		if m.cursor == i {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}
		s.WriteString(choices[i])
		s.WriteString("\n")
	}
	s.WriteString("\n(press q to quit, enter to confirm)\n")

	return s.String()
}

func Select_cmode() string {
	// This is where we'll listen for the choice the user makes in the Bubble
	// Tea program.
	result := make(chan string, 1)

	choices[0] = "distances"
	choices[1] = "blocks"
	choices[2] = "sequentially"
	message    = " Acceder a cifras de que forma?\n\n"

	// Pass the channel to the initialize function so our Bubble Tea program
	// can send the final choice along when the time comes.
	p := tea.NewProgram(model{cursor: 0, choice: result})
	if err := p.Start(); err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	// Print out the final choice.
	if r := <-result; r != "" {
		fmt.Printf("\n---\nYou have chosen: %s!\n", r)
		return r
	}
	return ""
}

func Select_number() string {
	result := make(chan string, 1)

	p:= tea.NewProgram(model{cursor : 0, choice: result})
	if err := p.Start(); err != nil {
		fmt.Println("Error! : ",err)
	}

	if r := <- result; r != ""{
		fmt.Printf("\v  %s will be analyzed.", r)
		return r
	}
	return ""



}



