package main

import (
	"fyne.io/fyne/v2"
	fyneApp "fyne.io/fyne/v2/app"

	"github.com/alionapermes/nstu-cs/internal/ui"
)

// const syntax:
// const <name> [<type>] = <literal>
const appTitle = "demo app"

func main() {
	app := fyneApp.New()

	window := app.NewWindow(appTitle)
	window.Resize(fyne.NewSize(600, 400))

	appUI := ui.New()

	canvas := window.Canvas()
	canvas.SetContent(appUI.Container)

	window.ShowAndRun()
}
