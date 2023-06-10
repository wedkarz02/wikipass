package gui

import (
	"strconv"
	c "wikipass/pkg/consts"
	"wikipass/pkg/pwder"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type App struct {
	Active       bool
	AuthError    error
	Fonts        Fonts
	MenuSect     rl.Rectangle
	MenuGenText  *Text
	GenText      *Text
	GenBounds    *Text
	TextBox      rl.Rectangle
	InputNum     *Text
	GenBtn       rl.Rectangle
	BtnText      *Text
	InvalidInput *Text
	Passwords    []*Text
	LogoutText   *Text
	LogoutBtn    rl.Rectangle
	Close        bool
}

func InitApp() *App {
	app := new(App)

	app.Active = false
	app.AuthError = nil
	app.Close = false
	app.Fonts = InitFonts()

	app.MenuSect = rl.Rectangle{
		X:      0,
		Y:      0,
		Width:  c.AppWindowWidth / 3,
		Height: c.AppWindowHeight,
	}

	app.MenuGenText = &Text{
		Content:  "Generate New\nPasswords",
		Font:     app.Fonts["arialb"],
		FontSize: 32,
		Color:    WhiteColor,
		Hidden:   false,
	}

	app.GenText = &Text{
		Content:  "Number of Passwords",
		Font:     app.Fonts["jbml"],
		FontSize: 18,
		Color:    LightGrey,
		Hidden:   false,
	}

	app.GenBounds = &Text{
		Content:  "0 < n < 28",
		Font:     app.Fonts["jbml"],
		FontSize: 18,
		Color:    LightGrey,
		Hidden:   false,
	}

	app.TextBox = rl.Rectangle{
		X:      20,
		Y:      app.MenuGenText.Size().Y + 120,
		Width:  app.MenuSect.Width - 2*20,
		Height: 40,
	}

	app.InputNum = &Text{
		Content:  "",
		Font:     app.Fonts["jbmr"],
		FontSize: 22,
		Color:    WhiteColor,
		Hidden:   false,
	}

	app.GenBtn = rl.Rectangle{
		X:      20,
		Y:      app.TextBox.Y + 60,
		Width:  app.TextBox.Width,
		Height: app.TextBox.Height,
	}

	app.BtnText = &Text{
		Content:  "Generate",
		Font:     app.Fonts["arialb"],
		FontSize: 24,
		Color:    WhiteColor,
		Hidden:   false,
	}

	app.InvalidInput = &Text{
		Content:  "Invalid Input.",
		Font:     app.Fonts["jbmb"],
		FontSize: 22,
		Color:    RedColor,
		Hidden:   true,
	}

	app.LogoutText = &Text{
		Content:  "Logout",
		Font:     app.Fonts["arialb"],
		FontSize: 24,
		Color:    TintColor,
		Hidden:   false,
	}

	app.LogoutBtn = rl.Rectangle{
		X:      app.MenuSect.Width/2 - app.LogoutText.Size().X/2,
		Y:      c.AppWindowHeight - 80,
		Width:  app.LogoutText.Size().X,
		Height: app.LogoutText.Size().Y,
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

func (app App) CheckBoundInput() bool {
	testNum, err := strconv.Atoi(app.InputNum.Content)

	if err != nil {
		return false
	}

	if testNum > 0 && testNum <= c.MaxPasswordsLen {
		return true
	}

	return false
}

func (app *App) WrapPasswords(newEl *Text) {
	app.Passwords = app.Passwords[1:]
	app.Passwords = append(app.Passwords, newEl)
}

func (app *App) UpdateApp(li *Login) {
	if li.Active {
		li.Active = false
		app.Resize()
	}

	CursorType(app.TextBox, rl.MouseCursorIBeam)
	app.InputNum.UpdateContent()

	ButtonAction(app.GenBtn, true, func() {
		if len(app.InputNum.Content) > 0 {
			if app.CheckBoundInput() {
				// TODO: Decrypt saved passwords and append them here

				app.InvalidInput.Hidden = true

				n, _ := strconv.Atoi(app.InputNum.Content)
				passwords := pwder.GetPasswords(n)

				for _, word := range passwords {
					password := &Text{
						Content:  word,
						Font:     app.Fonts["jbmr"],
						FontSize: 22,
						Color:    DarkGreyColor,
						Hidden:   false,
					}

					if len(app.Passwords) >= c.MaxPasswordsLen {
						app.WrapPasswords(password)
					} else {
						app.Passwords = append(app.Passwords, password)
					}
				}
			} else {
				app.InvalidInput.Hidden = false
			}
		}

		app.InputNum.Content = ""
	})

	ButtonAction(app.LogoutBtn, false, func() {
		app.Close = true
	})

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
	} else {
		for i, passwd := range app.Passwords {
			currentField := rl.Rectangle{
				X:      app.MenuSect.Width,
				Y:      float32(i*int(passwd.Size().Y) + 3),
				Width:  passwd.Size().X,
				Height: passwd.Size().Y,
			}

			if RectMouseCollision(currentField) {
				rl.DrawTextEx(passwd.Font,
					passwd.Content,
					rl.Vector2{
						X: app.MenuSect.Width + 10,
						Y: float32(i*int(passwd.Size().Y) + 3)},
					float32(passwd.FontSize), 0,
					DarkTintColor)
			} else {
				rl.DrawTextEx(passwd.Font,
					passwd.Content,
					rl.Vector2{
						X: app.MenuSect.Width + 10,
						Y: float32(i*int(passwd.Size().Y) + 3)},
					float32(passwd.FontSize), 0,
					passwd.Color)
			}
		}
	}

	rl.DrawTextEx(app.MenuGenText.Font,
		app.MenuGenText.Content,
		rl.Vector2{
			X: 20, Y: 30},
		float32(app.MenuGenText.FontSize), 0,
		app.MenuGenText.Color)

	rl.DrawTextEx(app.GenText.Font,
		app.GenText.Content,
		rl.Vector2{
			X: app.TextBox.X,
			Y: app.TextBox.Y - app.GenText.Size().Y*2 - 4},
		float32(app.GenText.FontSize), 0,
		app.GenText.Color)

	rl.DrawTextEx(app.GenBounds.Font,
		app.GenBounds.Content,
		rl.Vector2{
			X: app.TextBox.X,
			Y: app.TextBox.Y - app.GenBounds.Size().Y - 4},
		float32(app.GenBounds.FontSize), 0,
		app.GenBounds.Color)

	DrawTextBox(app,
		app.TextBox,
		app.InputNum,
		rl.Vector2{
			X: app.TextBox.X + 10,
			Y: app.TextBox.Y + app.TextBox.Height/2 - app.InputNum.Size().Y/2},
		BlackColor, false)

	DrawButton(app.GenBtn,
		app.BtnText,
		rl.Vector2{
			X: app.GenBtn.X + app.GenBtn.Width/2 - app.BtnText.Size().X/2,
			Y: app.GenBtn.Y + app.GenBtn.Height/2 - app.BtnText.Size().Y/2},
		TintColor,
		DarkTintColor)

	if !app.InvalidInput.Hidden {
		rl.DrawTextEx(app.InvalidInput.Font,
			app.InvalidInput.Content,
			rl.Vector2{
				X: app.MenuSect.Width/2 - app.InvalidInput.Size().X/2,
				Y: app.GenBtn.Y + 60},
			float32(app.InvalidInput.FontSize), 0,
			app.InvalidInput.Color)
	}

	if RectMouseCollision(app.LogoutBtn) {
		rl.DrawTextEx(app.LogoutText.Font,
			app.LogoutText.Content,
			rl.Vector2{
				X: app.LogoutBtn.X,
				Y: app.LogoutBtn.Y},
			float32(app.LogoutText.FontSize), 0,
			DarkTintColor)
	} else {
		rl.DrawTextEx(app.LogoutText.Font,
			app.LogoutText.Content,
			rl.Vector2{
				X: app.LogoutBtn.X,
				Y: app.LogoutBtn.Y},
			float32(app.LogoutText.FontSize), 0,
			app.LogoutText.Color)
	}
}
