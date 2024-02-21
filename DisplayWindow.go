package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type DisplayWindow struct {
	UniversitiesData *[]UniversityResponse
	DataDisplay      *widget.List
	UniversityInput  *widget.Entry
}

func getUniversityData() {
	searchTerm := window.UniversityInput.Text
	window.UniversitiesData = GetUniversityData(searchTerm)
}

func GetDataLen() int {
	if window.UniversitiesData == nil {
		return 0
	}
	return len(*window.UniversitiesData)
}

func CreateListItem() fyne.CanvasObject {
	return widget.NewLabel("Universities will appear here")
}

func UpdateListItem(itemNum widget.ListItemID, listItem fyne.CanvasObject) {
	UnivName := (*window.UniversitiesData)[itemNum].Name
	listItem.(*widget.Label).SetText(UnivName)
}

func BuildWindow(window *DisplayWindow, mainWindow fyne.Window) {
	window.UniversityInput = widget.NewEntry()
	window.UniversityInput.SetPlaceHolder("Enter University Name to search for...")
	getDataButton := widget.NewButton("Get University Data", getUniversityData)
	window.DataDisplay = widget.NewList(GetDataLen, CreateListItem, UpdateListItem)
	topPane := container.NewVBox(window.UniversityInput, getDataButton)
	leftPane := container.NewVSplit(topPane, window.DataDisplay)
	label1 := widget.NewLabel("Country:")
	countryField := widget.NewEntry()
	countryField.SetPlaceHolder("Country will appear here")
	label2 := widget.NewLabel("Website:")
	webSiteField := widget.NewEntry()
	webSiteField.SetPlaceHolder("Website will appear here")
	rightPane := container.NewGridWithColumns(2, label1, countryField, label2, webSiteField)
	contentPane := container.NewHSplit(leftPane, rightPane)
	mainWindow.SetContent(contentPane)
	mainWindow.Resize(fyne.NewSize(900, 900))
}
