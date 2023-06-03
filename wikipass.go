package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	font := rl.LoadFont("./assets/fonts/JetBrainsMono-Regular.ttf")

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawTextEx(font, "Hello, world! In JBMono!", rl.Vector2{X: 130, Y: 200}, 50, 1, rl.LightGray)

		rl.EndDrawing()
	}
}
