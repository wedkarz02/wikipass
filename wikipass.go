package main

import (
	"fmt"
	"strings"
	c "wikipass/pkg/consts"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(c.LogWindowWidth, c.LogWindowHeight, "Wikipass")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	font := rl.LoadFontEx("./assets/fonts/JetBrainsMono-Bold.ttf", 40, nil)

	bgColor := rl.Color{R: 0x2F, G: 0x36, B: 0x3D, A: 0xFF}
	textColor := rl.Color{R: 0xD1, G: 0xD5, B: 0xDA, A: 0xFF}

	text := "Input your\nmaster password:"
	textSize := rl.MeasureTextEx(font, text, 40, 0)

	textX := c.LogWindowWidth/2 - textSize.X/2
	textY := c.LogWindowHeight/2 - textSize.Y/2 - 80

	var inputText []string
	var hidden []string

	textBox := rl.Rectangle{X: c.LogWindowWidth * 0.12,
		Y:      c.LogWindowHeight/2 - 25,
		Width:  c.LogWindowWidth * 0.75,
		Height: 50}

	mouseOnText := false

	for !rl.WindowShouldClose() {
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), textBox) {
			mouseOnText = true
		} else {
			mouseOnText = false
		}

		if mouseOnText {
			rl.SetMouseCursor(rl.MouseCursorIBeam)
		} else {
			rl.SetMouseCursor(rl.MouseCursorDefault)
		}

		key := rl.GetCharPressed()

		for key > 0 {
			if key > 32 && key < 125 {
				inputText = append(inputText, string(key))
				hidden = append(hidden, "*")
			}

			key = rl.GetCharPressed()
		}

		if rl.IsKeyPressed(rl.KeyBackspace) {
			if len(inputText) > 0 {
				inputText = inputText[:len(inputText)-1]
				hidden = hidden[:len(hidden)-1]
			}
		}

		rl.BeginDrawing()

		rl.ClearBackground(bgColor)
		rl.DrawTextEx(font,
			text,
			rl.Vector2{X: textX, Y: textY},
			40, 0, textColor)

		rl.DrawRectangleRec(textBox, rl.LightGray)
		rl.DrawRectangleLinesEx(textBox, 1, rl.Black)

		rl.DrawTextEx(font, strings.Join(hidden, ""),
			rl.Vector2{X: textBox.X + 5, Y: textBox.Y + 8},
			40, 0, rl.Black)

		rl.EndDrawing()
	}

	fmt.Println(strings.Join(inputText, ""))
}
