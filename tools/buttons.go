package tools

import (
	"encoding/json"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"io"
	_ "io/ioutil"
	"os"
)

func OpenFile(win fyne.Window) {
	fileDialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err != nil || reader == nil {
			dialog.ShowError(err, win)
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
		openedFilePath := reader.URI().Name()
		var data []interface{}
		if err := json.Unmarshal(jsonData, &data); err != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		}
		UpdateWin(win, data, openedFilePath)
	}, win)
	fileDialog.Show()
}

func SaveFile(win fyne.Window, data []interface{}, openedFilePath string) {
	if openedFilePath == "" {
		dialog.ShowError(errors.New("no file opened"), win)
		return
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		dialog.ShowError(err, win)
		return
	}

	err = os.WriteFile(openedFilePath, jsonData, 0644)
	if err != nil {
		dialog.ShowError(err, win)
		return
	}
	dialog.ShowInformation("Save Success", "Data successfully saved", win)
}

func ExportData(win fyne.Window, data []interface{}) {
	saveDialog := dialog.NewFileSave(func(uc fyne.URIWriteCloser, err error) {
		if err != nil || uc == nil {
			return
		}
		defer func(uc fyne.URIWriteCloser) {
			err := uc.Close()
			if err != nil {

			}
		}(uc)
		err = json.NewEncoder(uc).Encode(data)
		if err != nil {
			dialog.ShowError(err, nil)
			return
		}
		dialog.ShowInformation("Export Success", "Data successfully exported", win)
	}, win)

	saveDialog.Show()
}
