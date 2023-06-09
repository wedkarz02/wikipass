package gui

import (
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawTextBox(cfg Config, tb rl.Rectangle, text *Text, inputPos rl.Vector2, tbColor rl.Color, hidden bool) {
	msg := text.Content

	if hidden {
		msg = strings.Repeat("*", len(msg))
	}

	size := text.Size()
	charSize := rl.MeasureTextEx(text.Font, " ", float32(text.FontSize), 0)
	maxLen := int(tb.Width/charSize.X) - 1

	time := rl.GetTime()
	fracTime := time - float64(int(time))

	rl.DrawRectangleRec(tb, tbColor)
	rl.DrawRectangleLinesEx(tb, 2.0, rl.Gray)

	if cfg.IsLogin() {
		rl.DrawTextEx(text.Font,
			"Master Password",
			rl.Vector2{X: tb.X + 10, Y: tb.Y + 6},
			20, 0,
			Grey)
	}

	if len(msg) < maxLen {
		rl.DrawTextEx(text.Font,
			msg,
			inputPos,
			float32(text.FontSize), 0,
			text.Color)

		if fracTime > 0.5 {
			rl.DrawTextEx(text.Font,
				"|",
				rl.Vector2{X: inputPos.X + size.X, Y: inputPos.Y},
				float32(text.FontSize), 0,
				text.Color)
		}
	} else {
		rl.DrawTextEx(text.Font,
			msg[:maxLen],
			inputPos,
			float32(text.FontSize), 0,
			text.Color)
	}
}

func DrawButton(btn rl.Rectangle, text *Text, textPos rl.Vector2, btnColor rl.Color, hoverColor rl.Color) {
	if RectMouseCollision(btn) {
		rl.DrawRectangleRec(btn, hoverColor)
	} else {
		rl.DrawRectangleRec(btn, btnColor)
	}

	rl.DrawTextEx(text.Font,
		text.Content,
		textPos,
		float32(text.FontSize), 0,
		text.Color)
}
