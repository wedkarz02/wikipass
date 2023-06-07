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
}

func InitLogin() *Login {
	li := new(Login) // li for LogIn

	li.Logo = InitLogo()
	li.Fonts = InitFonts()

	li.WelcomeText = &Text{Content: "Enter Master Password",
		Font:     li.Fonts["arialb"],
		FontSize: 32,
		Color:    WhiteColor,
	}

	li.InputText = &Text{Content: "",
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

	return li
}

func (li Login) UpdateLogin() {
	TextBoxCursorType(li.TextBox)
	li.InputText.UpdateContent()

	ButtonAction(li.UnlockBtn, func() {
		fmt.Println(li.InputText.Content)

		if aeswrapper.CheckDirExists(c.SecretDir) {
			aeswrapper.RmDir(c.SecretDir)
		} else {
			aeswrapper.InitSecretDir(c.SecretDir, c.IVFile, 32)
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
}
