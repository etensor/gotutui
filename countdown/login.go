package countdown

import (
	"log"

	"github.com/marcusolsson/tui-go"
)

var logo = `
	  	______ _____ _____  _   __  ______   ___     _______
     / ____/_  __/ ____/ / | / / /  ___/ /    \   /  _   \
    / __/   / / / __/   /  |/ /  \__ \  /  /\  ) /  -- __/
   / /___  / / / /___  / /|  /  ___/ / (  \/  / /  /\  \
  /_____/ /_/ /_____/ /_/ |_/  /____/   \____/ /__/  \__\ `

func Login() {
	user := tui.NewEntry()
	user.SetFocused(true)

	password := tui.NewEntry()
	password.SetEchoMode(tui.EchoModePassword)

	form := tui.NewGrid(0, 0)
	form.AppendRow(tui.NewLabel("User"), tui.NewLabel("Password"))
	form.AppendRow(user, password)

	status := tui.NewStatusBar("Ready.")

	login := tui.NewButton("[Login]")
	login.OnActivated(func(b *tui.Button) {
		status.SetText("Logged in.")
	})

	register := tui.NewButton("[Register]")

	buttons := tui.NewHBox(
		tui.NewSpacer(),
		tui.NewPadder(1, 0, login),
		tui.NewPadder(1, 0, register),
	)

	window := tui.NewVBox(
		tui.NewPadder(3, 1, tui.NewLabel(logo)),
		tui.NewPadder(12, 0, tui.NewLabel("Welcome to Skynet! Login or register.")),
		tui.NewPadder(1, 1, form),
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

	tui.DefaultFocusChain.Set(user, password, login, register)

	ui, err := tui.New(root)
	if err != nil {
		log.Fatal(err)
	}

	ui.SetKeybinding("Esc", func() { ui.Quit() })

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}
