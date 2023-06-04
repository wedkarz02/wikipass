package gui

import (
	"strings"

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

func DrawTextBox(tb rl.Rectangle,
	text string,
	font rl.Font,
	fontSize float32,
	tbColor rl.Color,
	textColor rl.Color,
	hidden bool) {

	if hidden {
		text = strings.Repeat("*", len(text))
	}

	size := rl.MeasureTextEx(font, text, fontSize, 0)
	charSize := rl.MeasureTextEx(font, " ", fontSize, 0)
	time := rl.GetTime()
	fracTime := time - float64(int(time))

	maxLen := int(tb.Width/charSize.X) - 2

	rl.DrawRectangleRec(tb, tbColor)
	rl.DrawRectangleLinesEx(tb, 3.0, rl.Gray)

	rl.DrawTextEx(font, "Master Password",
		rl.Vector2{X: tb.X + 10, Y: tb.Y + 6},
		20, 0,
		Grey)

	if len(text) < maxLen {
		rl.DrawTextEx(font, text,
			rl.Vector2{X: tb.X + 10, Y: tb.Y + 24},
			float32(fontSize), 0,
			textColor)

		if fracTime > 0.5 {
			rl.DrawTextEx(font, "|",
				rl.Vector2{X: tb.X + 10 + size.X, Y: tb.Y + 24},
				float32(fontSize), 0, textColor)
		}
	} else {
		rl.DrawTextEx(font, text[:maxLen],
			rl.Vector2{X: tb.X + 10, Y: tb.Y + 24},
			float32(fontSize), 0,
			textColor)
	}
}
