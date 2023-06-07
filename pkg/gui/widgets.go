package gui

import (
	"strings"
	c "wikipass/pkg/consts"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawTextBox(tb rl.Rectangle, text *Text, tbColor rl.Color, hidden bool) {
	msg := text.Content

	if hidden {
		msg = strings.Repeat("*", len(msg))
	}

	size := text.Size()
	charSize := rl.MeasureTextEx(text.Font, " ", float32(text.FontSize), 0)
	maxLen := int(tb.Width/charSize.X) - 2

	time := rl.GetTime()
	fracTime := time - float64(int(time))

	rl.DrawRectangleRec(tb, tbColor)
	rl.DrawRectangleLinesEx(tb, 2.0, rl.Gray)

	rl.DrawTextEx(text.Font,
		"Master Password",
		rl.Vector2{X: tb.X + 10, Y: tb.Y + 6},
		20, 0,
		Grey)

	if len(msg) < maxLen {
		rl.DrawTextEx(text.Font,
			msg,
			rl.Vector2{X: tb.X + 10, Y: tb.Y + 24},
			float32(text.FontSize), 0,
			text.Color)

		if fracTime > 0.5 {
			rl.DrawTextEx(text.Font,
				"|",
				rl.Vector2{X: tb.X + 10 + size.X, Y: tb.Y + 24},
				float32(text.FontSize), 0,
				text.Color)
		}
	} else {
		rl.DrawTextEx(text.Font,
			msg[:maxLen],
			rl.Vector2{X: tb.X + 10, Y: tb.Y + 24},
			float32(text.FontSize), 0,
			text.Color)
	}
}

func DrawButton(btn rl.Rectangle, text *Text, btnColor rl.Color, hoverColor rl.Color) {
	size := text.Size()
	textPos := rl.Vector2{
		X: c.LogWindowWidth/2 - size.X/2,
		Y: btn.Y + btn.Height/2 - size.Y/2,
	}

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
