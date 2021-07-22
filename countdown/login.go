package countdown

import (
	"log"
	"os"

	"fmt"
	"os/exec"

	"github.com/marcusolsson/tui-go"
)

var logo = `
	  	______ _____ _____  _   __  ______   ___     _______
     / ____/_  __/ ____/ / | / / /  ___/ /    \   /  _   \
    / __/   / / / __/   /  |/ /  \__ \  /  /\  ) /  -- __/
   / /___  / / / /___  / /|  /  ___/ / (  \/  / /  /\  \
  /_____/ /_/ /_____/ /_/ |_/  /____/   \____/ /__/  \__\ `

func Login() {
	number := tui.NewEntry()
	number.SetFocused(true)

	from := tui.NewEntry()
	//from.SetEchoMode(tui.EchoModePassword)
	to := tui.NewEntry()
	in_screen := tui.NewEntry()
	millis := tui.NewEntry()

	form := tui.NewGrid(0, 0)
	form.AppendRow(tui.NewLabel("Number: "), number)
	form.AppendRow(tui.NewLabel("From #: "), from)
	form.AppendRow(tui.NewLabel("To   #: "), to)
	form.AppendRow(tui.NewLabel("see  #: "), in_screen)
	form.AppendRow(tui.NewLabel("refresh rate: "), millis)

	status := tui.NewStatusBar("Ready.")

	run := tui.NewButton("[Run]")

	exit := tui.NewButton("[Salir]")
		exit.OnActivated(func(b *tui.Button) {
			status.SetText("ctrl+c | Esc para salir.")
	})

	buttons := tui.NewHBox(
		tui.NewSpacer(),
		tui.NewPadder(1, 0, run),
		tui.NewPadder(1, 0, exit),
	)

	window := tui.NewVBox(
		tui.NewPadder(3, 1, tui.NewLabel(logo)),
		tui.NewPadder(12, 0, tui.NewLabel("ephiphi, c++ digit analyzer.\n\t\t\t\t   e,pi,phi and 4 numbers")),
		tui.NewPadder(1, 2, form),
		buttons,
	)
	window.SetBorder(true)

	wrapper := tui.NewVBox(
		tui.NewSpacer(),
		window,
		tui.NewSpacer(),
	)
	content := tui.NewHBox(tui.NewSpacer(), wrapper, tui.NewSpacer())

	root := tui.NewVBox(
		content,
		status,
	)

	//tui.DefaultFocusChain.Set(user, password, login, register)
	tui.DefaultFocusChain.Set(number,from,to,in_screen,millis, run, exit)

	ui, err := tui.New(root)
	if err != nil {
		log.Fatal(err)
	}

	ui.SetKeybinding("Esc", func() { ui.Quit() })
	ui.SetKeybinding("ctrl+c", func() { ui.Quit() })

	run.OnActivated(func(b *tui.Button) {
		status.SetText("Inicializado.")
		command := exec.Command(fmt.Sprintf("./epiphi %s %s %s %s %s",number.Text(), from.Text(), to.Text(), in_screen.Text(), millis.Text()))
		command.Stdout = os.Stdout
		//out,err := command.CombinedOutput()

		if err != nil {
			fmt.Println(err.Error())
		}else{
			command.Run()
		}
		
	})

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}
