package main

import (
	c "wikipass/pkg/consts"
	"wikipass/pkg/gui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(c.LogWindowWidth, c.LogWindowHeight, "Wikipass")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	login := gui.InitLogin()

	for !rl.WindowShouldClose() {
		// UPDATE
		login.UpdateLogin()

		// DRAW
		rl.BeginDrawing()
		login.DrawLogin()
		rl.EndDrawing()
	}
}
