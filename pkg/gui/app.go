package gui

import (
	c "wikipass/pkg/consts"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type App struct {
	Active    bool
	AuthError error
	Fonts     Fonts
	TmpText   *Text
}

func InitApp() *App {
	app := new(App)

	app.Active = false
	app.AuthError = nil
	app.Fonts = InitFonts()

	app.TmpText = &Text{
		Content:  "hello",
		Font:     app.Fonts["arialb"],
		FontSize: 80,
		Color:    WhiteColor,
		Hidden:   false,
	}

	return app
}

func (app App) Resize() {
	rl.SetWindowSize(c.AppWindowWidth, c.AppWindowHeight)
}

func (app *App) UpdateApp(li *Login) {
	if li.Active {
		li.Active = false
		app.Resize()
	}

	// TODO: Handle this when decrypting the file
	// app.AuthError = errors.New("dziaba dziaba dziaba")
}

func (app *App) DrawApp() {
	rl.ClearBackground(DarkGreyColor)

	if app.AuthError != nil {
		rl.DrawTextEx(app.Fonts["jbmb"],
			"Authentication Error...",
			rl.Vector2{X: 100, Y: 100},
			float32(60), 0,
			RedColor)

		return
	}

	rl.DrawTextEx(app.TmpText.Font,
		app.TmpText.Content,
		rl.Vector2{X: 350, Y: 250},
		float32(app.TmpText.FontSize), 0,
		app.TmpText.Color)
}
