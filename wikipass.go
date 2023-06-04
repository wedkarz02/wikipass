package main

import (
	"fmt"
	"strings"
	c "wikipass/pkg/consts"
	"wikipass/pkg/gui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(c.LogWindowWidth, c.LogWindowHeight, "Wikipass")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	logo := gui.InitLogo()

	fontBold := rl.LoadFontEx("./assets/fonts/arialbd.ttf", 40, nil)
	fontJBMB := rl.LoadFontEx("./assets/fonts/JetBrainsMono-Bold.ttf", 40, nil)

	text := "Enter Master Password"
	textSize := rl.MeasureTextEx(fontBold, text, 32, 0)

	textX := c.LogWindowWidth/2 - textSize.X/2
	textY := c.LogWindowHeight/2 - 65

	var inputText []string

	textBox := gui.InitTextBox(c.LogWindowWidth*0.08,
		c.LogWindowHeight/2,
		c.LogWindowWidth*0.84,
		50)

	for !rl.WindowShouldClose() {
		gui.TextBoxCursorType(textBox)
		gui.UpdateInput(&inputText)

		rl.BeginDrawing()
		rl.ClearBackground(gui.DarkGreyColor)

		rl.DrawTexture(logo, c.LogWindowWidth/2-c.LogoWidth/2, 60, rl.White)

		rl.DrawTextEx(fontBold,
			text,
			rl.Vector2{X: textX, Y: float32(textY)},
			32, 0, gui.WhiteColor)

		gui.DrawTextBox(textBox,
			strings.Join(inputText, ""),
			fontJBMB,
			20, gui.BlackColor,
			gui.WhiteColor,
			true)

		rl.EndDrawing()
	}

	fmt.Println(strings.Join(inputText, ""))
}
