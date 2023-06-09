package gui

import (
	c "wikipass/pkg/consts"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type App struct {
	Active      bool
	AuthError   error
	Fonts       Fonts
	MenuSect    rl.Rectangle
	MenuGenText *Text
}

func InitApp() *App {
	app := new(App)

	app.Active = false
	app.AuthError = nil
	app.Fonts = InitFonts()

	app.MenuSect = rl.Rectangle{
		X:      0,
		Y:      0,
		Width:  c.AppWindowWidth / 3,
		Height: c.AppWindowHeight,
	}

	app.MenuGenText = &Text{
		Content:  "Generate\nNew\nPasswords",
		Font:     app.Fonts["arialb"],
		FontSize: 44,
		Color:    WhiteColor,
		Hidden:   false,
	}

	return app
}

func (app App) Resize() {
	rl.SetWindowSize(c.AppWindowWidth, c.AppWindowHeight)
	rl.SetWindowPosition(c.AppWindowPosX, c.AppWindowPosY)
}

func (app *App) UpdateApp(li *Login) {
	if li.Active {
		li.Active = false
		app.Resize()
	}

	// TODO: Handle this when decrypting the file
	// app.AuthError = errors.New("dziaba dziaba dziaba")
}

func (app *App) DrawAuthErr() {
	rl.DrawTextEx(app.Fonts["jbmb"],
		"Authentication Error...",
		rl.Vector2{X: 100, Y: 100},
		float32(60), 0,
		RedColor)
}

func (app *App) DrawApp() {
	rl.ClearBackground(WhiteColor)

	if app.AuthError != nil {
		app.DrawAuthErr()
		return
	}

	rl.DrawRectangleRec(app.MenuSect, DarkGreyColor)

	rl.DrawTextEx(app.MenuGenText.Font,
		app.MenuGenText.Content,
		rl.Vector2{
			X: 30, Y: 20},
		float32(app.MenuGenText.FontSize), 0,
		app.MenuGenText.Color)
}
