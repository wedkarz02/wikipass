package gui

import (
	c "wikipass/pkg/consts"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func InitApp() {
	rl.InitWindow(c.LogWindowWidth, c.LogWindowHeight, "Wikipass")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)
}

func UpdateInput(textData *[]string) {
	key := rl.GetCharPressed()

	for key > 0 {
		if key > 32 && key < 125 {
			*textData = append(*textData, string(key))
		}

		key = rl.GetCharPressed()
	}

	if rl.IsKeyPressed(rl.KeyBackspace) {
		tmpText := []string(*textData)
		if len(tmpText) > 0 {
			*textData = tmpText[:len(tmpText)-1]
		}
	}
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

func ButtonAction(btn rl.Rectangle, callBack func()) {
	if RectMouseCollision(btn) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			callBack()
			return
		}
	}

	if rl.IsKeyPressed(rl.KeyEnter) {
		callBack()
		return
	}
}
