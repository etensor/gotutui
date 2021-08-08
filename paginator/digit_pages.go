package paginator

// A simple program demonstrating the paginator component from the Bubbles
// component library.

// A modified simple program displaying elegantly the digit distribution of magical numbers

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	//"regexp"

	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/lipgloss"
	"github.com/rivo/tview"
	//"github.com/gdamore/tcell/v2"

	tea "github.com/charmbracelet/bubbletea"
)

// pwd of pi_viz is not paginator, but root.
type model struct {
	items     [10]string
	paginator paginator.Model
	mode      string
}

func newModel(pimode string, num_chosen string) model {
	//physicalWidth, _, _ := term.GetSize(int(os.Stdout.Fd()))
	var n_cyphers int

	if pimode == "distances" {n_cyphers = 250
	}else if pimode == "blocks" {n_cyphers = 1200}

	var items [10]string
	for i := 0; i < 10; i++ {

		//cmd := exec.Command("zsh", "./paginator/pi_viz.sh", "500", fmt.Sprintf("%d", i))
		cmd := exec.Command("zsh", "./paginator/cypher_viz.sh",num_chosen,fmt.Sprintf("%d", n_cyphers), fmt.Sprintf("%d", i))

		out, err := cmd.CombinedOutput()

		if err == nil {

			var triangle_items []string
			if pimode == "distances" || pimode == "" {
				triangle_items = strings.SplitAfter(string(out), fmt.Sprintf("K%d", i))

				var b strings.Builder
				for j := 0; j < len(triangle_items); j++ {
					b.WriteString(triangle_items[j] + "\n   ")
				}
				items[i] = b.String() 

			} else if pimode == "blocks" { 
				items[i] = string(out)
			}
		}
		// if it fails, check fmt.Printf("%+q",string(out)) result and fix.

	}

	p := paginator.NewModel()
	p.Type = paginator.Dots
	p.PerPage = 1
	p.ActiveDot = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF00")).Render("•")
	p.InactiveDot = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "250", Dark: "238"}).Render("•")
	p.SetTotalPages(len(items))

	return model{
		paginator: p,
		items:     items,
		mode:      pimode,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	}
	m.paginator, cmd = m.paginator.Update(msg)
	return m, cmd
}

func (m model) View() string {
	var b, a strings.Builder
	a.WriteString("\n    e π φ  digits distribution\n\n")
	start, end := m.paginator.GetSliceBounds(len(m.items))

	switch m.mode {
	case "distances":
		{
			for _, item := range m.items[start:end] {
				a.WriteString("   " + item)
			}
			b.WriteString("  " + m.paginator.View() + "\n\n  h/l ←/→ page • q: quit\n")
		}
	case "blocks":
		{
			b.Reset()
			for _, item := range m.items[start:end] {
				defer fmt.Printf("\033c\n e π φ Digits Distribution\n\n%s\n\t\t\t  %s\n\t\t    h/l ←/→ page • q: quit\n", item, m.paginator.View())
			}
		}
	}

	return a.String() + b.String()
}

func window_viz(num_chosen string) {
	app := tview.NewApplication()
	textView := tview.NewTextView().	
		SetScrollable(true).
		SetChangedFunc(func() {
			app.Draw()
		})
	
	textView.SetText("Epa").SetBorder(true).SetTitle("Ni idea")

	closure := Code(textView,20,20,"a brrrrrr")

	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(closure,0,1,true)

	app.SetRoot(layout, true).Run()
	go func() {
		cmd := exec.Command("./epiphi", num_chosen, "0", "150", "20", "80")

		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		//cmd.Stdin = os.Stdin
	
		err := cmd.Run()
	
		if err != nil {
			fmt.Println("\t cmd failed.")
		}
		
	}()

		
}

// The width of the code window.
const codeWidth = 56

// Code returns a primitive which displays the given primitive (with the given
// size) on the left side and its source code on the right side.
func Code(p tview.Primitive, width, height int, code string) tview.Primitive {
        // Set up code view.
        codeView := tview.NewTextView().
                SetWrap(false).
                SetDynamicColors(true)
        codeView.SetBorderPadding(1, 1, 2, 0)
        fmt.Fprint(codeView, code)

        return tview.NewFlex().
                AddItem(Center(width, height, p), 0, 1, true).
                AddItem(codeView, codeWidth, 1, false)
}


// Center returns a new primitive which shows the provided primitive in its
// center, given the provided primitive's size.
func Center(width, height int, p tview.Primitive) tview.Primitive {
	return tview.NewFlex().
			AddItem(nil, 0, 1, false).
			AddItem(tview.NewFlex().
					SetDirection(tview.FlexRow).
					AddItem(nil, 0, 1, false).
					AddItem(p, height, 1, true).
					AddItem(nil, 0, 1, false), width, 1, true).
			AddItem(nil, 0, 1, false)
}


func Pager(mode string, num_chosen string) {
	if mode == "sequentially" {
		window_viz(num_chosen)

	}
	p := tea.NewProgram(newModel( mode, num_chosen ))
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
	//q := tea.NewProgram(newModel (mode, num_chosen))
	//if err := q.Start(); err != nil {
	//	log.Fatal(err)
	//}

}
