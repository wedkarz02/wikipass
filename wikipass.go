package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"wikipass/pkg/aeswrapper"
	c "wikipass/pkg/consts"
	"wikipass/pkg/gui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func CheckDirExists(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}

func RmDir(path string) {
	err := os.RemoveAll(path)

	if err != nil {
		log.Fatalln("[ERROR]: Directory deletion failed: ", err)
	}
}

func main() {
	rl.InitWindow(c.LogWindowWidth, c.LogWindowHeight, "Wikipass")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	logo := gui.InitLogo()
	fonts := gui.InitFonts()

	welcomeText := gui.Text{Content: "Enter Master Password",
		Font:     fonts["arialb"],
		FontSize: 32,
		Color:    gui.WhiteColor,
	}

	inputText := gui.Text{Content: "",
		Font:     fonts["jbmb"],
		FontSize: 20,
		Color:    gui.WhiteColor,
	}

	textBox := rl.Rectangle{
		X:      c.LogWindowWidth * 0.08,
		Y:      c.LogWindowHeight / 2,
		Width:  c.LogWindowWidth * 0.84,
		Height: 50,
	}

	unlockBtn := rl.Rectangle{
		X:      textBox.X,
		Y:      textBox.Y + 80,
		Width:  textBox.Width,
		Height: textBox.Height,
	}

	for !rl.WindowShouldClose() {
		gui.TextBoxCursorType(textBox)
		inputText.UpdateContent()

		gui.ButtonAction(unlockBtn, func() {
			fmt.Println(inputText.Content)

			if CheckDirExists(c.SecretDir) {
				RmDir(c.SecretDir)
			} else {
				aeswrapper.InitSecretDir(c.SecretDir, c.IVFile, 32)
			}
		})

		rl.BeginDrawing()
		rl.ClearBackground(gui.DarkGreyColor)

		rl.DrawTexture(logo, c.LogWindowWidth/2-c.LogoWidth/2, 60, rl.White)

		rl.DrawTextEx(welcomeText.Font,
			welcomeText.Content,
			rl.Vector2{X: c.LogWindowWidth/2 - welcomeText.Size().X/2,
				Y: c.LogWindowHeight/2 - 65},
			float32(welcomeText.FontSize), 0,
			welcomeText.Color)

		gui.DrawTextBox(textBox, &inputText, gui.BlackColor, true)

		gui.DrawButton(unlockBtn,
			"Unlock Wikipass",
			fonts["arialb"],
			24,
			gui.TintColor,
			gui.DarkTintColor,
			gui.WhiteColor)

		rl.EndDrawing()
	}
}
