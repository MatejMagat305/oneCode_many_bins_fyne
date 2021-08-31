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
	game := g.game
	pixDensity := game.pixelDensity()
	pixW, pixH := game.cellForCoord(w, h, pixDensity)
	game.initIfNeed(pixW, pixH)
	img := g.imgCache
	game.lock.Lock()
	if img == nil || img.Bounds().Size().X != pixW || img.Bounds().Size().Y != pixH {
		img = image.NewRGBA(image.Rect(0, 0, pixW, pixH))
		g.imgCache = img
	}
	if game.snake == nil ||game.Board == nil {
		game.lock.Unlock()
		return img
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
	body := game.snake.body
	for i := 0; i < len(body); i++ {
		go func(i int) {
			img.Set(body[i].X, body[i].Y, color.RGBA{G: 255, A:255})
			ch<-true
		}(i)
	}
	go func() {
		food :=game.food
		img.Set(food.X, food.Y, color.RGBA{
			R: 200, //
			G: 0,
			B: 200,
			A: 255,
		})
		ch<-true
	}()
	for i := 0; i < len(body)+1; i++ {
		<-ch
	}
	game.lock.Unlock()
	return img
}
