package gui

import (
	"wikipass/pkg/aeswrapper"
	c "wikipass/pkg/consts"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Login struct {
	Active        bool
	Fonts         Fonts
	Logo          rl.Texture2D
	WelcomeText   *Text
	IncorrectPass *Text
	InputText     *Text
	TextBox       rl.Rectangle
	UnlockBtn     rl.Rectangle
	UnlockText    *Text
	ResetText     *Text
	ResetBtn      rl.Rectangle
}

func InitLogin() *Login {
	li := new(Login) // li for LogIn

	li.Active = true
	li.Logo = InitLogo()
	li.Fonts = InitFonts()

	li.WelcomeText = &Text{
		Content:  "Enter Master Password",
		Font:     li.Fonts["arialb"],
		FontSize: 32,
		Color:    WhiteColor,
		Hidden:   false,
	}

	li.IncorrectPass = &Text{
		Content:  "Incorrect Master Password",
		Font:     li.Fonts["jbmb"],
		FontSize: 20,
		Color:    RedColor,
		Hidden:   true,
	}

	li.InputText = &Text{
		Content:  "",
		Font:     li.Fonts["jbmb"],
		FontSize: 20,
		Color:    WhiteColor,
		Hidden:   false,
	}

	li.TextBox = rl.Rectangle{
		X:      c.LogWindowWidth * 0.08,
		Y:      c.LogWindowHeight / 2,
		Width:  c.LogWindowWidth * 0.84,
		Height: 50,
	}

	li.UnlockBtn = rl.Rectangle{
		X:      li.TextBox.X,
		Y:      li.TextBox.Y + 80,
		Width:  li.TextBox.Width,
		Height: li.TextBox.Height,
	}

	li.UnlockText = &Text{
		Content:  "Unlock Wikipass",
		Font:     li.Fonts["arialb"],
		FontSize: 24,
		Color:    WhiteColor,
		Hidden:   false,
	}

	li.ResetText = &Text{
		Content:  "Reset Master Password",
		Font:     li.Fonts["arialb"],
		FontSize: 18,
		Color:    TintColor,
		Hidden:   false,
	}

	li.ResetBtn = rl.Rectangle{
		X:      c.LogWindowWidth/2 - li.ResetText.Size().X/2,
		Y:      li.UnlockBtn.Y + 80,
		Width:  li.ResetText.Size().X,
		Height: li.ResetText.Size().Y,
	}

	return li
}

func (li Login) IsLogin() bool {
	return true
}

func (li Login) UpdateLogin(app *App) {
	if !aeswrapper.CheckIfExists(c.SecretDir) {
		li.WelcomeText.Content = "Set Master Password"
	} else {
		li.WelcomeText.Content = "Enter Master Password"
	}

	CursorType(li.TextBox, rl.MouseCursorIBeam)
	li.InputText.UpdateContent()

	ButtonAction(li.UnlockBtn, true, func() {
		if aeswrapper.CheckIfExists(c.SecretDir) {
			key := aeswrapper.GenKey(li.InputText.Content)

			if aeswrapper.TestAuth(c.AuthFile, key) {
				li.IncorrectPass.Hidden = true
				li.Active = false
				app.Active = true
			} else {
				li.IncorrectPass.Hidden = false
			}
		} else {
			aeswrapper.InitSecretDir(c.SecretDir, c.IVFile, 32)
			key := aeswrapper.GenKey(li.InputText.Content)
			aeswrapper.InitAuth(c.AuthFile, key, 64)
		}

		li.InputText.Content = ""
	})

	ButtonAction(li.ResetBtn, false, func() {
		li.IncorrectPass.Hidden = true

		if aeswrapper.CheckIfExists(c.SecretDir) {
			aeswrapper.RmDir(c.SecretDir)
		}
	})
}

func (li Login) DrawLogin() {
	rl.ClearBackground(DarkGreyColor)

	rl.DrawTexture(li.Logo, c.LogWindowWidth/2-c.LogoWidth/2, 60, rl.White)

	rl.DrawTextEx(li.WelcomeText.Font,
		li.WelcomeText.Content,
		rl.Vector2{X: c.LogWindowWidth/2 - li.WelcomeText.Size().X/2,
			Y: c.LogWindowHeight/2 - 65},
		float32(li.WelcomeText.FontSize), 0,
		li.WelcomeText.Color)

	if !li.IncorrectPass.Hidden {
		rl.DrawTextEx(li.IncorrectPass.Font,
			li.IncorrectPass.Content,
			rl.Vector2{
				X: li.TextBox.X,
				Y: li.TextBox.Y - 24},
			float32(li.IncorrectPass.FontSize), 0,
			li.IncorrectPass.Color)
	}

	DrawTextBox(li,
		li.TextBox,
		li.InputText,
		rl.Vector2{
			X: li.TextBox.X + 10,
			Y: li.TextBox.Y + 24},
		BlackColor, true)

	DrawButton(li.UnlockBtn,
		li.UnlockText,
		rl.Vector2{
			X: c.LogWindowWidth/2 - li.UnlockText.Size().X/2,
			Y: li.UnlockBtn.Y + li.UnlockBtn.Height/2 - li.UnlockText.Size().Y/2},
		TintColor,
		DarkTintColor)

	if RectMouseCollision(li.ResetBtn) {
		rl.DrawTextEx(li.ResetText.Font,
			li.ResetText.Content,
			rl.Vector2{
				X: li.ResetBtn.X,
				Y: li.ResetBtn.Y},
			float32(li.ResetText.FontSize), 0,
			DarkTintColor)
	} else {
		rl.DrawTextEx(li.ResetText.Font,
			li.ResetText.Content,
			rl.Vector2{
				X: li.ResetBtn.X,
				Y: li.ResetBtn.Y},
			float32(li.ResetText.FontSize), 0,
			li.ResetText.Color)
	}
}
