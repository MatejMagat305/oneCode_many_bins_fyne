package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"math/rand"
	"simple_snake/game"
	"time"
)

func show(app fyne.App) {
	rand.Seed(time.Now().UnixNano())
	g := game.NewGame()
	window := app.NewWindow("simple_snake")

	window.SetContent(g.BuildUI())
	window.Canvas().SetOnTypedRune(g.TypedRune)
	window.Resize(fyne.NewSize(50,50))
	g.Animate()
	window.Show()
}

func main() {
	appNew := app.New()
	show(appNew)
	appNew.Run()
}
