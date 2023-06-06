package gui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Text struct {
	Content  string
	Font     rl.Font
	FontSize int
	Color    rl.Color
}

func (text *Text) Size() rl.Vector2 {
	return rl.MeasureTextEx(text.Font, text.Content, float32(text.FontSize), 0)
}

func (text *Text) UpdateContent() {
	key := rl.GetCharPressed()

	for key > 0 {
		if key > 32 && key < 125 {
			text.Content += string(key)
		}

		key = rl.GetCharPressed()
	}

	if rl.IsKeyPressed(rl.KeyBackspace) {
		if len(text.Content) > 0 {
			text.Content = text.Content[:len(text.Content)-1]
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
