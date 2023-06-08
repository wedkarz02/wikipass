package gui

import (
	c "wikipass/pkg/consts"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type App struct {
	Active  bool
	Fonts   Fonts
	TmpText *Text
}

func InitApp() *App {
	app := new(App)

	app.Active = false
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

func (app App) UpdateApp(li *Login) {
	if li.Active {
		li.Active = false
		app.Resize()
	}
}

func (app App) DrawApp() {
	rl.ClearBackground(DarkGreyColor)

	rl.DrawTextEx(app.TmpText.Font,
		app.TmpText.Content,
		rl.Vector2{X: 350, Y: 250},
		float32(app.TmpText.FontSize), 0,
		app.TmpText.Color)
}
