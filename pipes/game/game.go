package game

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"pipes/board"
	"pipes/coords"
	"pipes/vars"
	"sync"
	"time"
)

type Game struct {
	widget.BaseWidget
	genText        *widget.Label
	Paused         bool
	Board          *board.Broard
	snake          *Pipe
	PauseButton    *widget.Button
	lose           bool
	fourDirections []*coords.PointDirect
	lock           sync.Mutex
	I              int64
}

func (g *Game) Lose() {
	g.lose = true
}

func (g *Game) CreateRenderer() fyne.WidgetRenderer {
	renderer := &GameRenderer{game: g}
	render := canvas.NewRaster(renderer.draw)
	render.ScaleMode = canvas.ImageScalePixels
	renderer.render = render
	renderer.objects = []fyne.CanvasObject{render}
	renderer.ApplyTheme()

	return renderer
}

func (g *Game) cellForCoord(x, y int, density float32) (int, int) {
	xpos := int(float32(x) / float32(cellSize) / density)
	ypos := int(float32(y) / float32(cellSize) / density)
	return xpos, ypos
}

func (g *Game) run() {
	g.Paused = false
}

func (g *Game) stop() {
	g.Paused = true
}

func (g *Game) toggleRun() {
	g.Paused = !g.Paused
}

func (g *Game) TypedRune(r rune) {
	if r == ' ' {
		g.toggleRun()
	}
}

func (g *Game) Tapped(ev *fyne.PointEvent) {
	if g.Board == nil {
		return
	}
	min := len(*g.Board) * len((*g.Board)[0])
	for _, direction := range g.fourDirections {
		d := direction.RP.DistanceOf(int(ev.Position.X), int(ev.Position.Y))
		if d < min {
			g.snake.direct = direction.D.Clone()
			min = d
		}
	}

}


func (g *Game) TappedSecondary( *fyne.PointEvent) {}
func (g *Game) normalButtonFunction(pause *widget.Button) func() {
	pause.Text = "Play"

	return func() {
		if g.Paused && g.lose {
			g.snake.NeedReset()
		}
		g.Paused = !g.Paused
		if g.Paused {
			pause.Text = "Play"
		} else {
			pause.Text = "Pause"
		}
		pause.Refresh()
	}
}
func (g *Game) BuildUI() fyne.CanvasObject {
	var pause *widget.Button
	pause = widget.NewButton("Pause", nil)

	title := container.New(layout.NewGridLayout(2), g.genText, pause)
	g.PauseButton = pause
	pause.OnTapped = g.normalButtonFunction(pause)
	return container.New(layout.NewBorderLayout(title, nil, nil, nil), title, g)
}

func (g *Game) pixelDensity() float32 {
	c := fyne.CurrentApp().Driver().CanvasForObject(g)
	if c == nil {
		return 1.0
	}
	pixW, _ := c.PixelCoordinateForPosition(fyne.NewPos(cellSize, cellSize))
	return float32(pixW) / float32(cellSize)
}
func NewGame() *Game {
	g := &Game{genText: widget.NewLabel("score: 0")}
	g.ExtendBaseWidget(g)
	return g
}

func (g *Game) initIfNeed(w int, h int) {
	if g.snake == nil || g.snake.NeedReset() {
		g.snake = NewSnake(coords.NewPoint(w/2, h/2))

	}
	if g.Board != nil {
		return
	}
	g.Board = board.NewBoard(w, h)
	g.Board.MakeColor(g.fourDirections)
	ch <- true
}

var (
	ch = make(chan bool)
)

func (g *Game) Animate() {
	go func() {
		<-ch
		for {
			time.Sleep(g.snake.duration)
			if !g.Paused {
				g.snake.Play(g)
				g.genText.SetText(fmt.Sprint("score: ", len(g.snake.body)-1))
			}
			g.Refresh()
		}
	}()
}

func (g *Game) InitDirections(w int, h int, rW, rH int) {
	if (g.fourDirections) != (nil) {
		return
	}
	centerX, centerY, rCenterX, rCcenterY := (w)/2, (h)/2, rW/2, rH/2
	g.fourDirections = make([]*coords.PointDirect, 0, 4)
	g.fourDirections = append(g.fourDirections, coords.NewPointDirect(coords.NewPoint(w, centerY),
		coords.NewDirect(1, 0), coords.NewPoint(rW, rCcenterY)))
	g.fourDirections = append(g.fourDirections, coords.NewPointDirect(coords.NewPoint(0, centerY),
		coords.NewDirect(-1, 0), coords.NewPoint(0, rCcenterY)))
	g.fourDirections = append(g.fourDirections, coords.NewPointDirect(coords.NewPoint(centerX, h),
		coords.NewDirect(0, 1), coords.NewPoint(rCenterX, rH)))
	g.fourDirections = append(g.fourDirections, coords.NewPointDirect(coords.NewPoint(centerX, 0),
		coords.NewDirect(0, -1), coords.NewPoint(rCenterX, 0)))
	for i, rgba := range vars.C {
		g.fourDirections[i].C = rgba
	}
}
