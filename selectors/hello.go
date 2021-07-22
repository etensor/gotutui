package selectors

// bubbletea to-do-list probando

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// Vamos a hacer una lista.
// Ahora el modelo:

type model struct {
	choices  []string         // items de la lista
	cursor   int              // para moverse entre los items
	selected map[int]struct{} // cuales items estan seleccionados
}

// init, update, view

var initialModel = model{
	// La lista es una lista de productos:
	choices: []string{"erbin suav", "birritas frias", "pizzitas melas", "programar TUIs"},

	// Un mapa indica bien cuales estan seleccionadas.
	selected: make(map[int]struct{}),
}

func (m model) Init() tea.Cmd {
	// return `nil` -> no + I/O
	return nil
}

// Update (tea.KeyMsg)
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// Hay una tecla presionada:
	case tea.KeyMsg:

		// Cual?
		switch msg.String() {

		//Salir del programa
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// enter, spacebar: select item
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}				
			}
		}
	}

	// return updated model, no command
	return m, nil
}

// View

func (m model) View() string {
	// header
	s := "Melo pa una fiestica, tener:\n\n"

	//iterar sobre choices
	for i, choice := range m.choices {
		// cursor en lugar i
		cursor := " " // no
		if m.cursor == i {
			cursor = ">" // cursor
		}

		//cursor act
		checked := " " // no act
		if _, ok := m.selected[i]; ok {
			checked = "x" // act, seleccionado
			// now use this to execute programs, different pages. But this would be for multi element selection...

		}

		// Renderizar fila:
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)

	}

	// footer:
	s += "\n Presionar q para salir.\n"

	return s
}

func Select_items() {
	//fmt.Println("Hello,World!")

	p := tea.NewProgram(initialModel)
	if err := p.Start(); err != nil {
		fmt.Printf("Error pana! %v", err)
		os.Exit(1)
	}
}
