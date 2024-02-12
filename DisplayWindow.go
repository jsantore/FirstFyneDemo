package main

import (
	"fyne.io/fyne/v2"
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
