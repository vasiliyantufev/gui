package component

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func GetTabFiles() *container.TabItem {
	labelPlug := widget.NewLabel("Список файлов")
	return container.NewTabItem("Файлы", labelPlug)
}
