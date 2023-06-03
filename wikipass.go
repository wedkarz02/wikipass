package main

import (
	"fmt"
	"sync"
	"wikipass/pkg/pwder"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	font := rl.LoadFont("./assets/fonts/JetBrainsMono-Regular.ttf")

	var wg sync.WaitGroup
	var passwords []string
	passwdChan := make(chan string)

	n := 15

	for i := 0; i < n; i++ {
		wg.Add(1)
		go pwder.GenPassword(passwdChan, &wg)
	}

	for i := 0; i < n; i++ {
		passwords = append(passwords, <-passwdChan)
	}

	wg.Wait()

	for _, passwd := range passwords {
		fmt.Println(passwd)
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawTextEx(font, "Hello, world! In JBMono!", rl.Vector2{X: 130, Y: 200}, 50, 1, rl.LightGray)

		rl.EndDrawing()
	}
}
