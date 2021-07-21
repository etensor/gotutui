package paginator

// A simple program demonstrating the paginator component from the Bubbles
// component library.

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	//"regexp"

	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/lipgloss"

	tea "github.com/charmbracelet/bubbletea"
)

// pwd of pi_viz is not paginator, but root.

func newModel() model {
	//physicalWidth, _, _ := term.GetSize(int(os.Stdout.Fd()))

	var items [10]string
	for i := 0; i < 10; i++ {
		//text := fmt.Sprintf("Item %d", i)
		//command := fmt.Sprintf("zsh ../seqpi/pi_viz.sh 100 %d", i)
		//out, err := exec.Command("echo", "hello world").Output()
		cmd := exec.Command("zsh", "./paginator/pi_viz.sh", "300", fmt.Sprintf("%d", i)) // finally.
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
			//items[i] = string(out) /// string doesn't get fit to screen.
			//fmt.Printf("\t,, %d \n", strings.Count(string(out),""))
			//fmt.Printf("\n now wtf: \v %s", out)

			/*
				if physicalWidth < len(string(out)) {
					j := 1
					k := len(string(out))/physicalWidth

					for j <= k {
							//items[i] += string(out)[((j-1)*physicalWidth):(j*physicalWidth)] + "\n    "
						j++
					}
				}
			*/
			//exp := fmt.Sprintf("%d(?=\\033[0;31m)",i) 				/// go regexp doesnt support lookarounds, perl syntax
			//exp := fmt.Sprintf(`\033[01;31m`+"%d"+`\033[0m`,i) 						//	///	/	/	 after i, back to normal.
			//r := regexp.MustCompile(exp)										//// problem was solved, but it was 01;31, not 0;31 tone
			//str := string(out)
			//exp := fmt.Sprintf(`\033[01;31m`+"%d",i) 																
			//triangle_items := strings.SplitAfter(string(out),fmt.Sprint(i))
			fmt.Printf("%+q",string(out))										/// This prints everything raw, see final char, split later.
			triangle_items := strings.SplitAfter(string(out),"01;31m")						 /// it was way more easy...
			// this causes \n before cypher $i
			// for after: 														/// if it fails, check fmt.Printf("%+q",string(out)) result and fix.
			
			/// fmt.Printf("%+q",string(out)) solved symbols used to colour findings 

			//triangle_items := r.Split(str,-1)									
			//triangle_items := strings.SplitAfter(str, exp)
			
			var b strings.Builder

			for j:=0;j<len(triangle_items);j++ {
				b.WriteString(triangle_items[j]+"\n   ")
			}

			/// Works but as grep colors cyphers red, 0;31m is added each time, then is found and split,
			/// breaking the string.
				// ` ` between regexp.MustCompile, exclude 0;31m // \033[0;31m  			<- valid for echo. not only available
				// line 62

			items[i] = b.String()
		}
		/*
			if output, err := cmd.Output(); err != nil {
				items[i] = string(output)
			} else {
				fmt.Print("\t\v error: ", err)
			}*/

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
	}
}

type model struct {
	items     [10]string
	paginator paginator.Model
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
	var b strings.Builder
	b.WriteString("\n  Paginator Example\n\n")
	start, end := m.paginator.GetSliceBounds(len(m.items))
	for _, item := range m.items[start:end] {
		b.WriteString("   ")
		b.WriteString(item)
		b.WriteString("\n\n")
	}
	b.WriteString("  " + m.paginator.View())
	b.WriteString("\n\n  h/l ←/→ page • q: quit\n")
	return b.String()
}

func Pager() {

	p := tea.NewProgram(newModel())
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
