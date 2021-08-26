package game

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"image"
	"image/color"
)

const (
	cellSize = 8
)

type GameRenderer struct {
	render   *canvas.Raster
	objects  []fyne.CanvasObject
	imgCache *image.RGBA

	aliveColor color.Color
	deadColor  color.Color

	game *Game
}

func (g *GameRenderer) MinSize() fyne.Size {
	return fyne.NewSize(500, 500)
}

func (g *GameRenderer) Layout(size fyne.Size) {
	g.render.Resize(size)
}

func (g *GameRenderer) ApplyTheme() {
	g.aliveColor = theme.ForegroundColor()
	g.deadColor = theme.BackgroundColor()
}

func (g *GameRenderer) BackgroundColor() color.Color {
	return theme.BackgroundColor()
}

func (g *GameRenderer) Refresh() {
	canvas.Refresh(g.render)
}

func (g *GameRenderer) Objects() []fyne.CanvasObject {
	return g.objects
}

func (g *GameRenderer) Destroy() {
}

func (g *GameRenderer) draw(w, h int) image.Image {
	pixDensity := g.game.pixelDensity()
	pixW, pixH := g.game.cellForCoord(w, h, pixDensity)
	g.game.InitDirections(pixW, pixH, w, h)
	g.game.initIfNeed(pixW, pixH)
	img := g.imgCache
	if img == nil || img.Bounds().Size().X != pixW || img.Bounds().Size().Y != pixH {
		img = image.NewRGBA(image.Rect(0, 0, pixW, pixH))
		g.imgCache = img
	}
	for i := 0; i < pixW; i++ {
		row := (*g.game.Board)[i]
		for j := 0; j < pixH; j++ {
			go func(i, j int) {
				img.Set(i, j, color.White)
				img.Set(i, j, row[j].GetColor())
				ch<-true
			}(i,j)
		}
	}
	for i := 0; i < pixW; i++ {
		for j := 0; j < pixH; j++ {
			<-ch
		}
	}
	body := g.game.snake.body
	for i := 0; i < len(body); i++ {
		go func(i int) {
			img.Set(body[i].P.X, body[i].P.Y, color.RGBA{G: 255, A:255})
			ch<-true
		}(i)
	}
	for i := 0; i < len(body); i++ {
		<-ch
	}
	return img
}
