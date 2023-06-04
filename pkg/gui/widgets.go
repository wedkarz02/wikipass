package gui

import rl "github.com/gen2brain/raylib-go/raylib"

func InitTextBox(x int, y int, w int, h int) rl.Rectangle {
	return rl.Rectangle{X: float32(x), Y: float32(y), Width: float32(w), Height: float32(h)}
}

func RectMouseCollision(tb rl.Rectangle) bool {
	return rl.CheckCollisionPointRec(rl.GetMousePosition(), tb)
}

func TextBoxCursorType(tb rl.Rectangle) {
	if RectMouseCollision(tb) {
		rl.SetMouseCursor(rl.MouseCursorIBeam)
	} else {
		rl.SetMouseCursor(rl.MouseCursorDefault)
	}
}
