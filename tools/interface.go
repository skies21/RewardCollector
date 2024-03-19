package tools

import (
	"RewardCollector/table"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func UpdateWin(win fyne.Window, data []interface{}) {
	mainToolbar := CreateToolbar(win, data)
	label := widget.NewLabel("Rewards Table")
	toolbarHeight := mainToolbar.MinSize().Height
	mainToolbar.Move(fyne.NewPos(0, 0))
	mainToolbar.Resize(fyne.NewSize(win.Canvas().Size().Width, toolbarHeight))
	label.Move(fyne.NewPos(0, toolbarHeight))
	label.Resize(fyne.NewSize(win.Canvas().Size().Width, label.MinSize().Height))

	newTable := table.UpdateTableContent(data)
	newTable.Move(fyne.NewPos(0, toolbarHeight+label.MinSize().Height))
	newTable.Resize(fyne.NewSize(700, (toolbarHeight+label.MinSize().Height)+1000))

	content := container.NewWithoutLayout(mainToolbar, label, newTable)
	win.Resize(fyne.NewSize(700, (toolbarHeight+label.MinSize().Height)+600))
	win.SetContent(content)
}
