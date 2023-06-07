package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Tabs Example")

	label1 := widget.NewLabel("This is tab 1")
	label2 := widget.NewLabel("This is tab 2")
	label3 := widget.NewLabel("This is tab 3")

	tab1 := container.NewTabItem("Tab 1", label1)
	tab2 := container.NewTabItem("Tab 2", label2)
	tab3 := container.NewTabItem("Tab 3", label3)

	tabs := container.NewAppTabs(tab1, tab2, tab3)

	myWindow.SetContent(tabs)
	myWindow.ShowAndRun()
}
