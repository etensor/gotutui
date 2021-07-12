package countdown

// Vamos a ver si hay que reimportar algo en un archivo
//del mismo package
import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	duration = time.Second * 10
	interval = time.Millisecond
)

type tickMsg time.Time

type model struct {
	timeout  time.Time
	lastTick time.Time
}

func (m model) Init() tea.Cmd {
	return tick()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		}

	case tickMsg:
		t := time.Time(msg)
		if t.After(m.timeout) {
			return m, tea.Quit
		}
		m.lastTick = t
		return m, tick()
	}

	return m, nil
}

func (m model) View() string {
	t := m.timeout.Sub(m.lastTick).Milliseconds()
	secs := t / 1000
	millis := t % 1000 / 10
	return fmt.Sprintf("Este programa finalizará en %02d:%02d\n", secs, millis)

}

func tick() tea.Cmd {
	return tea.Tick(time.Duration(interval), func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func Countdown() {
	m := model{
		timeout: time.Now().Add(duration),
	}

	if err := tea.NewProgram(m).Start(); err != nil {
		fmt.Println("Hmm, no funcionó!", err)
		os.Exit(1)
	}
}
