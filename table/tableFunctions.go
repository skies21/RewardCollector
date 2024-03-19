package table

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"sort"
)

func UpdateTableContent(data []interface{}) *widget.Table {
	list := widget.NewTable(
		func() (int, int) {
			return len(data) + 1, 4
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Table Content Example")
		},
		func(cell widget.TableCellID, obj fyne.CanvasObject) {
			if cell.Row == 0 {
				switch cell.Col {
				case 0:
					obj.(*widget.Label).SetText("Rarity")
				case 1:
					obj.(*widget.Label).SetText("Game Type")
				case 2:
					obj.(*widget.Label).SetText("Reward")
				case 3:
					obj.(*widget.Label).SetText("Amount")
				}
			} else {
				row := cell.Row - 1
				if row < len(data) {
					col := cell.Col
					reward := data[row].(map[string]interface{})
					var text string
					switch col {
					case 0:
						rarity := fmt.Sprintf("%s", reward["rarity"])
						switch rarity {
						case "COMMON":
							text = "⚪️ " + rarity + " ⚪️"
						case "RARE":
							text = "🔵 " + rarity + " 🔵"
						case "EPIC":
							text = "🍆 " + rarity + " 🍆"
						case "LEGENDARY":
							text = "☀️ " + rarity + " ☀️"
						default:
							text = rarity
						}
					case 1:
						if gameType, ok := reward["gameType"].(string); ok {
							text = gameType
						}
					case 2:
						text = fmt.Sprintf("%s", reward["reward"])
					case 3:
						if reward["amount"] != nil {
							text = fmt.Sprintf("%v", reward["amount"])
						} else {
							text = "1"
						}
					}
					obj.(*widget.Label).SetText(text)
				}
			}
		})
	return list
}

func FilterTable(data []interface{}) []interface{} {
	rarityPriority := map[string]int{
		"LEGENDARY": 4,
		"EPIC":      3,
		"RARE":      2,
		"COMMON":    1,
	}

	sortedData := make([]interface{}, len(data))
	copy(sortedData, data)
	sort.SliceStable(sortedData, func(i, j int) bool {
		rarityI := sortedData[i].(map[string]interface{})["rarity"].(string)
		rarityJ := sortedData[j].(map[string]interface{})["rarity"].(string)
		return rarityPriority[rarityI] > rarityPriority[rarityJ]
	})

	return sortedData
}
