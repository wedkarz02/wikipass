package gui

import (
	"log"
	c "wikipass/pkg/consts"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Fonts map[string]rl.Font

type Text struct {
	Content  string
	Font     rl.Font
	FontSize int
	Color    rl.Color
}

func (list Fonts) AddFont(name string, alias string, defRes int32) {
	if _, inFonts := list[alias]; inFonts {
		log.Fatalln("[ERROR]: Font already loaded!")
	}

	list[alias] = rl.LoadFontEx("./assets/fonts/"+name+".ttf", defRes, nil)
}

func (text *Text) Size() rl.Vector2 {
	return rl.MeasureTextEx(text.Font, text.Content, float32(text.FontSize), 0)
}

func InitFonts() Fonts {
	fonts := make(Fonts)

	fonts.AddFont("arialbd", "arialb", 40)
	fonts.AddFont("JetBrainsMono-Bold", "jbmb", 40)

	return fonts
}

func InitLogo() rl.Texture2D {
	logo := rl.LoadImage("./assets/logo.png")
	rl.ImageResize(logo, c.LogoWidth, c.LogoHeight)
	txtLogo := rl.LoadTextureFromImage(logo)
	rl.UnloadImage(logo)

	return txtLogo
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
