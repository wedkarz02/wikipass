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

	rl.SetWindowPosition(c.LogWindowPosX, c.LogWindowPosY)

	login := gui.InitLogin()
	app := gui.InitApp()
	var passKey string

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		if login.Active {
			// UPDATE
			tmpKey := login.UpdateLogin(app)
			if tmpKey != "" {
				passKey = tmpKey
			}

			// DRAW
			login.DrawLogin()
		}

		if app.Active {
			// UPDATE
			app.UpdateApp(login, passKey)
			if app.Close {
				break
			}

			// DRAW
			app.DrawApp()
		}

		rl.EndDrawing()
	}
}
