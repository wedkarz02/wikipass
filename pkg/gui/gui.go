package gui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func UpdateInput(textData *[]string) {
	key := rl.GetCharPressed()

	for key > 0 {
		if key > 32 && key < 125 {
			*textData = append(*textData, string(key))
		}

		key = rl.GetCharPressed()
	}

	if rl.IsKeyPressed(rl.KeyBackspace) {
		tmpText := []string(*textData)
		if len(tmpText) > 0 {
			*textData = tmpText[:len(tmpText)-1]
		}
	}
}
