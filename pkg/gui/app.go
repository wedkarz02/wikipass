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
	TextBox     rl.Rectangle
	InputNum    *Text
	Passwords   []*Text
}

func InitApp() *App {
	app := new(App)

	app.Active = false
	app.AuthError = nil
	app.Fonts = InitFonts()

	app.MenuSect = rl.Rectangle{
		X:      0,
		Y:      0,
		Width:  c.AppWindowWidth / 3.5,
		Height: c.AppWindowHeight,
	}

	app.MenuGenText = &Text{
		Content:  "Generate New\nPasswords",
		Font:     app.Fonts["arialb"],
		FontSize: 32,
		Color:    WhiteColor,
		Hidden:   false,
	}

	app.TextBox = rl.Rectangle{
		X:      20,
		Y:      app.MenuGenText.Size().Y + 60,
		Width:  app.MenuSect.Width - 2*20,
		Height: 40,
	}

	app.InputNum = &Text{
		Content:  "",
		Font:     app.Fonts["jbmb"],
		FontSize: 28,
		Color:    WhiteColor,
		Hidden:   false,
	}

	return app
}

func (app App) IsLogin() bool {
	return false
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

	if len(app.Passwords) == 0 {
		rl.DrawTextEx(app.Fonts["arialb"],
			"Nothing to show yet. Generate some passwords.",
			rl.Vector2{
				X: app.MenuSect.Width + 30,
				Y: 30},
			float32(22), 0,
			DarkGreyColor)
	}

	rl.DrawTextEx(app.MenuGenText.Font,
		app.MenuGenText.Content,
		rl.Vector2{
			X: 20, Y: 30},
		float32(app.MenuGenText.FontSize), 0,
		app.MenuGenText.Color)

	DrawTextBox(app,
		app.TextBox,
		app.InputNum,
		rl.Vector2{
			X: app.TextBox.X + 5,
			Y: app.TextBox.Y + app.TextBox.Height/2 - app.InputNum.Size().Y/2},
		BlackColor, false)
}
