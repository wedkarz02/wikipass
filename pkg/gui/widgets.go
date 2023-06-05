package gui

import (
	"strings"
	c "wikipass/pkg/consts"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func InitRect(x int, y int, w int, h int) rl.Rectangle {
	return rl.Rectangle{X: float32(x), Y: float32(y), Width: float32(w), Height: float32(h)}
}

func InitLogo() rl.Texture2D {
	logo := rl.LoadImage("./assets/logo.png")
	rl.ImageResize(logo, c.LogoWidth, c.LogoHeight)
	txtLogo := rl.LoadTextureFromImage(logo)
	rl.UnloadImage(logo)

	return txtLogo
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

func DrawButton(btn rl.Rectangle,
	text string,
	font rl.Font,
	fontSize float32,
	btnColor rl.Color,
	hoverColor rl.Color,
	textColor rl.Color) {

	size := rl.MeasureTextEx(font, text, fontSize, 0)
	textPos := rl.Vector2{
		X: c.LogWindowWidth/2 - size.X/2,
		Y: btn.Y + btn.Height/2 - size.Y/2,
	}

	if RectMouseCollision(btn) {
		rl.DrawRectangleRec(btn, hoverColor)
	} else {
		rl.DrawRectangleRec(btn, btnColor)
	}

	rl.DrawTextEx(font, text,
		textPos,
		fontSize, 0,
		textColor)
}
