package tools

import (
	"RewardCollector/table"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func CreateToolbar(win fyne.Window, data []interface{}) *widget.Toolbar {
	openButton := widget.NewToolbarAction(theme.FolderOpenIcon(), func() {
		OpenFile(win)
	})

	saveButton := widget.NewToolbarAction(theme.DocumentSaveIcon(), func() {
		SaveFile(win)
	})

	exportButton := widget.NewToolbarAction(theme.DownloadIcon(), func() {
		ExportData(win)
	})

	filterButton := widget.NewToolbarAction(theme.HelpIcon(), func() {
		filteredData := table.FilterTable(data)
		UpdateWin(win, filteredData)
	})

	toolbar := widget.NewToolbar(
		openButton,
		saveButton,
		exportButton,
		filterButton,
	)
	return toolbar
}
