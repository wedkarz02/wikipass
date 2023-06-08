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
	app := gui.InitApp()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		if login.Active {
			// UPDATE
			login.UpdateLogin(app)

			// DRAW
			login.DrawLogin()
		}

		if app.Active {
			// UPDATE
			app.UpdateApp(login)

			// DRAW
			app.DrawApp()
		}

		rl.EndDrawing()
	}
}
