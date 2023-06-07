package main

import (
	"fmt"
	"wikipass/pkg/aeswrapper"
	c "wikipass/pkg/consts"
	"wikipass/pkg/gui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(c.LogWindowWidth, c.LogWindowHeight, "Wikipass")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	// plain := []string{
	// 	"hello world, test message for encryption",
	// 	"second line",
	// 	"another one",
	// 	"another one",
	// 	"actually... another one",
	// }
	key := aeswrapper.GenKey("masterpassword")
	// aeswrapper.EncryptAES(c.EncryptionFile, plain, key)
	plain, err := aeswrapper.DecryptAES(c.EncryptionFile, key)
	if err != nil {
		fmt.Println("key not correct: ", err)
	} else {
		for _, word := range plain {
			fmt.Println(word)
		}
	}

	login := gui.InitLogin()

	for !rl.WindowShouldClose() {
		// UPDATE
		login.UpdateLogin()

		// DRAW
		rl.BeginDrawing()
		login.DrawLogin()
		rl.EndDrawing()
	}
}
