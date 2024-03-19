package main

import (
	"RewardCollector/tools"
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2/app"
	"os"
	"strings"
)

func main() {
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}
	var jsonData []byte
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") {
			jsonData, err = os.ReadFile(file.Name())
			if err != nil {
				fmt.Println("Error reading file:", err)
				return
			}
			break
		}
	}

	if jsonData == nil {
		jsonData = []byte(`[{"amount":1,"rarity":"COMMON","reward":"no .json found"}]`)
	}

	var data []interface{}
	if err := json.Unmarshal(jsonData, &data); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	myApp := app.New()
	win := myApp.NewWindow("Rewards")

	tools.UpdateWin(win, data)

	win.ShowAndRun()
}
