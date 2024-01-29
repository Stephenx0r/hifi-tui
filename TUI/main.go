package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/sachinsenal0x64/Hifi-Tui/TUI/components"
)

// Song informations

type Song struct {
	duration float64
	title    string
	artist   string
	album    string
}

type App struct {
	tview         *tview.Application
	rootContainer *tview.Flex
	library       *components.LibraryComponets
}

func CreateApp() *App {
	lib := components.LibraryComponets()

	a := App{
		tview:         tview.NewApplication(),
		rootContainer: tview.NewFlex(),
		library:       lib,
	}

	a.rootContainer.AddItem(a.library.Container, 0, 1, true).
		AddItem(a.library.Container, 0, 1, false)

	a.tview.SetRoot(a.rootContainer, true).SetFocus(a.rootContainer).EnableMouse(true)

	a.SetInputHandlers()

	return &a
}

func (a App) Start() error {
	return a.tview.Run()
}

func (a *App) SetInputHandlers() {
	a.rootContainer.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		focus := a.tview.GetFocus()

		switch event.Key() {
		case tcell.KeyCtrlL:
			if focus == a.library.BodyComponent {
				a.tview.SetFocus(a.library.BodyComponent)
			}
		}
		return event
	})
}

func main() {
	// app := tview.NewApplication()

	// // 2x2

	// flex := tview.NewFlex().
	// 	AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
	// 		AddItem(tview.NewBox().SetBorder(true).SetTitle("Library"), 0, 1, false).
	// 		AddItem(tview.NewBox().SetBorder(true).SetTitle("Lyrics"), 0, 1, false), 0, 1, false).

	// 	// 2x2

	// 	AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
	// 		AddItem(tview.NewBox().SetBorder(true).SetTitle("Songs"), 0, 10, false).
	// 		AddItem(tview.NewBox().SetBorder(true).SetTitle("Player"), 0, 1, false), 0, 3, false)

	// // Set up an event loop
	// app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

	// 	switch event.Rune() {
	// 	case 'k':
	// 		flex.SetBorder(true)

	// 	case 'j':
	// 		flex.SetBorder(false)
	// 	}

	// 	// Check for the q key press
	// 	if event.Rune() == 'q' {
	// 		app.Stop()
	// 		return nil
	// 	}

	// 	return event

	// })

	// // The application will run until app.Stop() is called (in response to the Escape key press)

	// if err := app.SetRoot(flex, true).Run(); err != nil {
	// 	panic(err)
	// }

	app := CreateApp()
	if err := app.Start(); err != nil {
		panic(err)
	}

}
