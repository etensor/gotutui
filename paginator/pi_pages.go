package paginator

// A simple program demonstrating the paginator component from the Bubbles
// component library.

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

func newModel(pimode string) model {
	//physicalWidth, _, _ := term.GetSize(int(os.Stdout.Fd()))

	// f := func(c rune) bool {
	// 	return unicode.IsSpace(c) || c == '.'
	// }

	// s := "We. are.. them. 314152718"
	// fmt.Printf("Fields are: %q\n", strings.FieldsFunc(s, f))		//// erases . and splits by " "

	var items [10]string
	for i := 0; i < 10; i++ {
		//text := fmt.Sprintf("Item %d", i)
		//command := fmt.Sprintf("zsh ../seqpi/pi_viz.sh 100 %d", i)
		//out, err := exec.Command("echo", "hello world").Output()
		cmd := exec.Command("zsh", "./paginator/pi_viz.sh", "500", fmt.Sprintf("%d", i)) // finally.
		//cmd := exec.Command("pwd") //"\"$pwd\"")

		/// Does not execute command. but defines it.
		/// Cmd.Run
		/*
			cmd_iny := &exec.Cmd{
				Path:   "zsh ../seqpi/pi_viz.sh",
				Args:   []string{"../seqpi/pi_viz.sh", "100", fmt.Sprintf("-c %d", i)},
				Stdout: os.Stdout,
				Stderr: os.Stdout,
			}
		*/

		out, err := cmd.CombinedOutput()
		//cmd.Wait()
		if err == nil {

			//exp := fmt.Sprintf("%d(?=\\033[0;31m)",i) 				/// go regexp doesnt support lookarounds, perl syntax
			//exp := fmt.Sprintf(`\033[01;31m`+"%d"+`\033[0m`,i) 						//	///	/	/	 after i, back to normal.
			//r := regexp.MustCompile(exp)										//// problem was solved, but it was 01;31, not 0;31 tone
			//str := string(out)
			//exp := fmt.Sprintf(`\033[01;31m`+"%d",i)
			//triangle_items := strings.SplitAfter(string(out),fmt.Sprint(i))
			//fmt.Printf("%+q", string(out)+"\n")                              /// This prints everything raw, see final char, split later.
			// #9 : \x1b[01;31m\x1b[K9\x1b[m\x1b[K75665\x1b[01;31m\x1b[K9

			var triangle_items []string
			if pimode == "distances" || pimode == "" {
				triangle_items = strings.SplitAfter(string(out), fmt.Sprintf("K%d", i)) /// it was way more easy...

				var b strings.Builder
				for j := 0; j < len(triangle_items); j++ {
					b.WriteString(triangle_items[j] + "\n   ")
				}
				items[i] = b.String()

			} else if pimode == "blocks" { /// maybe there is a workaround....
				/// maybe running the command when its called to page, it can be slow, is seems the only way.
				/// changing the order of printing did it... no need to rerun
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

/// Problem: item is just a string. split is not symmetrical.
// if split after %d,i, ++"\n"
// if split after %d, any after N numbers are printed. -> sym to width.

func (m model) View() string {
	var b, a strings.Builder
	a.WriteString("\n  Pi Digits Distribution\n\n")
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
				defer fmt.Printf("\033c\n  Pi Digits Distribution\n\n%s\n\t\t\t  %s\n\t\t    h/l ←/→ page • q: quit\n", item, m.paginator.View())
			}
		}
	}

	return a.String() + b.String()
}

func window_viz() {
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
		cmd := exec.Command("./epiphi", "phi", "0", "100", "25", "100")

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


func Pager(mode string) {
	if mode == "sequentially" {
		window_viz()

	}
	p := tea.NewProgram(newModel(mode))
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
