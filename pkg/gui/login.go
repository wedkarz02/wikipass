package gui

import (
	"fmt"
	"wikipass/pkg/aeswrapper"
	c "wikipass/pkg/consts"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Login struct {
	Fonts       Fonts
	Logo        rl.Texture2D
	WelcomeText *Text
	InputText   *Text
	TextBox     rl.Rectangle
	UnlockBtn   rl.Rectangle
	ResetText   *Text
	ResetBtn    rl.Rectangle
}

func InitLogin() *Login {
	li := new(Login) // li for LogIn

	li.Logo = InitLogo()
	li.Fonts = InitFonts()

	li.WelcomeText = &Text{
		Content:  "Enter Master Password",
		Font:     li.Fonts["arialb"],
		FontSize: 32,
		Color:    WhiteColor,
	}

	li.InputText = &Text{
		Content:  "",
		Font:     li.Fonts["jbmb"],
		FontSize: 20,
		Color:    WhiteColor,
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

	li.ResetText = &Text{
		Content:  "Reset Master Password",
		Font:     li.Fonts["arialb"],
		FontSize: 18,
		Color:    TintColor,
	}

	li.ResetBtn = rl.Rectangle{
		X:      c.LogWindowWidth/2 - li.ResetText.Size().X/2,
		Y:      li.UnlockBtn.Y + 80,
		Width:  li.ResetText.Size().X,
		Height: li.ResetText.Size().Y,
	}

	return li
}

func (li Login) UpdateLogin() {
	CursorType(li.TextBox, rl.MouseCursorIBeam)
	li.InputText.UpdateContent()

	ButtonAction(li.UnlockBtn, true, func() {
		if aeswrapper.CheckDirExists(c.SecretDir) {
			key := aeswrapper.GenKey(li.InputText.Content)

			if aeswrapper.TestAuth(c.AuthFile, key) {
				fmt.Println("logged in correctly")
			} else {
				fmt.Println("key incorrect")
			}
		}
	})

	ButtonAction(li.ResetBtn, false, func() {
		if aeswrapper.CheckDirExists(c.SecretDir) {
			aeswrapper.RmDir(c.SecretDir)
		} else {
			// TODO: Remove this later, it's just to make testing easier
			aeswrapper.InitSecretDir(c.SecretDir, c.IVFile, 32)
			key := aeswrapper.GenKey(li.InputText.Content)
			aeswrapper.InitAuth(c.AuthFile, key, 64)
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

	DrawTextBox(li.TextBox, li.InputText, BlackColor, true)

	DrawButton(li.UnlockBtn,
		&Text{
			Content:  "Unlock Wikipass",
			Font:     li.Fonts["arialb"],
			FontSize: 24,
			Color:    WhiteColor},
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
