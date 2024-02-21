package main

import (
	"fyne.io/fyne/v2/app"
)

var window DisplayWindow

func main() {
	displayApp := app.New()
	mainWindow := displayApp.NewWindow("University Data in a List")
	window = DisplayWindow{}
	BuildWindow(&window, mainWindow)
	mainWindow.ShowAndRun()
}
