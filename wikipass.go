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

	fontBold := rl.LoadFontEx("./assets/fonts/arialbd.ttf", 40, nil)
	fontJBMB := rl.LoadFontEx("./assets/fonts/JetBrainsMono-Bold.ttf", 40, nil)

	welcomeText := gui.Text{Content: "Enter Master Password",
		Font:     fontBold,
		FontSize: 32,
		Color:    gui.WhiteColor,
	}

	inputText := gui.Text{Content: "",
		Font:     fontJBMB,
		FontSize: 20,
		Color:    gui.WhiteColor,
	}

	textBox := gui.InitRect(c.LogWindowWidth*0.08,
		c.LogWindowHeight/2,
		c.LogWindowWidth*0.84,
		50)

	unlockBtn := gui.InitRect(int(textBox.X),
		int(textBox.Y)+80,
		int(textBox.Width),
		int(textBox.Height))

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
			welcomeText.Size(),
			float32(welcomeText.FontSize), 0,
			welcomeText.Color)

		gui.DrawTextBox(textBox, &inputText, gui.BlackColor, true)

		gui.DrawButton(unlockBtn,
			"Unlock Wikipass",
			fontBold,
			24,
			gui.TintColor,
			gui.DarkTintColor,
			gui.WhiteColor)

		rl.EndDrawing()
	}
}
