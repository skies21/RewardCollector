package tools

import (
	"RewardCollector/table"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func CreateToolbar(win fyne.Window, data []interface{}, openedFilePath string) *widget.Toolbar {
	openButton := widget.NewToolbarAction(theme.FolderOpenIcon(), func() {
		OpenFile(win)
	})

	saveButton := widget.NewToolbarAction(theme.DocumentSaveIcon(), func() {
		SaveFile(win, data, openedFilePath)
	})

	exportButton := widget.NewToolbarAction(theme.DownloadIcon(), func() {
		ExportData(win, data)
	})

	filterButton := widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {
		filteredData := table.FilterTable(data)
		UpdateWin(win, filteredData, openedFilePath)
	})

	toolbar := widget.NewToolbar(
		openButton,
		saveButton,
		exportButton,
		filterButton,
	)
	return toolbar
}
