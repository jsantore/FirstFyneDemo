package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var window DisplayWindow

func main() {
	displayApp := app.New()
	mainWindow := displayApp.NewWindow("University Data in a List")
	window = DisplayWindow{}
	window.UniversityInput = widget.NewEntry()
	window.UniversityInput.SetPlaceHolder("Enter University Name to search for...")
	getDataButton := widget.NewButton("Get University Data", getUniversityData)
	window.DataDisplay = widget.NewList(GetDataLen, CreateListItem, UpdateListItem)
	topPane := container.NewVBox(window.UniversityInput, getDataButton)
	//listPane := container.NewStack(window.DataDisplay)
	//listPane.Resize(fyne.NewSize(400, 800))
	//contentPane := container.NewVBox(window.UniversityInput, getDataButton, listPane)
	contentPane := container.NewVSplit(topPane, window.DataDisplay)
	mainWindow.SetContent(contentPane)
	mainWindow.Resize(fyne.NewSize(400, 900))
	mainWindow.ShowAndRun()
}
