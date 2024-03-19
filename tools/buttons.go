package tools

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"io"
)

func OpenFile(win fyne.Window) {
	fileDialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.ShowError(err, win)
			return
		}
		if reader == nil {
			return
		}
		defer func(reader fyne.URIReadCloser) {
			err := reader.Close()
			if err != nil {

			}
		}(reader)

		jsonData, err := io.ReadAll(reader)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		var data []interface{}
		if err := json.Unmarshal(jsonData, &data); err != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		}
		UpdateWin(win, data)
	}, win)
	fileDialog.Show()
}

func SaveFile(win fyne.Window) {
	fileDialog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
		if err != nil {
			dialog.ShowError(err, win)
			return
		}
		if writer == nil {
			return
		}
		defer func(writer fyne.URIWriteCloser) {
			err := writer.Close()
			if err != nil {

			}
		}(writer)

		// Запись данных в файл
		// Пример:
		// _, err = writer.Write([]byte("Hello, world!"))
		// if err != nil {
		// 	dialog.ShowError(err, win)
		// 	return
		// }
	}, win)
	fileDialog.Show()
}

func ExportData(win fyne.Window) {
	// Пример данных, которые будут экспортированы
	data := map[string]interface{}{
		"key": "value",
	}

	saveDialog := dialog.NewFileSave(func(uc fyne.URIWriteCloser, err error) {
		if err != nil || uc == nil {
			return
		}
		defer func(uc fyne.URIWriteCloser) {
			err := uc.Close()
			if err != nil {

			}
		}(uc)

		// Запись данных в файл
		err = json.NewEncoder(uc).Encode(data)
		if err != nil {
			dialog.ShowError(err, nil)
			return
		}
	}, win)

	// Отображение диалога
	saveDialog.SetFileName("data.json")
	saveDialog.Show()
}
