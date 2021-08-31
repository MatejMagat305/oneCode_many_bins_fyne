package game

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"simple_snake/board"
	"simple_snake/coords"
	"sync"
	"time"
)

type Game struct {
	widget.BaseWidget
	genText        *widget.Label
	Paused         bool
	Board          *board.Broard
	food          *board.Food
	snake          *Snake
	PauseButton    *widget.Button
	lose           bool
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
	_ = len(*g.Board) * len((*g.Board)[0])
}

func (g *Game) TappedSecondary( *fyne.PointEvent) {}
func (g *Game) normalButtonFunction(pause *widget.Button) func() {
	g.setPauseName(pause)
	pause.Refresh()
	return func() {
		if g.Paused && g.lose {
			g.snake.NeedReset()
		}
		g.Paused = !g.Paused
		g.setPauseName(pause)
		pause.Refresh()
	}
}

func (g *Game) setPauseName(pause *widget.Button) {
	if g.Paused {
		pause.Text = "Play"
	} else {
		pause.Text = "Pause"
	}
}

func (g *Game) BuildUI() fyne.CanvasObject {
	pause := widget.NewButton("Pause", nil)
	topButton := widget.NewButton("up", func() {
		if g.snake != nil {
			g.snake.direct = coords.NewDirect(0, -1)
		}
	})
	bottonButton := widget.NewButton("down", func() {
		if g.snake != nil {
			g.snake.direct = coords.NewDirect(0, 1)
		}
	})
	leftButton := widget.NewButton("left", func() {
		if g.snake != nil {
			g.snake.direct = coords.NewDirect(-1, 0)
		}
	})
	rightButton := widget.NewButton("right", func() {
		if g.snake != nil {
			g.snake.direct = coords.NewDirect( 1, 0)
		}
	})
	down := container.New(layout.NewGridLayout(2), g.genText, pause)
	buttons := []fyne.CanvasObject{topButton, bottonButton, leftButton, rightButton}
	g.PauseButton = pause
	pause.OnTapped = g.normalButtonFunction(pause)
	gameController := container.New(layout.NewBorderLayout(buttons[0], buttons[1], buttons[2], buttons[3]),
		buttons...)
	for i := 0; i < len(buttons); i++ {
		buttons[i].Resize(fyne.NewSize(50,50))
	}
	return container.New(layout.NewBorderLayout(down, gameController, nil, nil), down, gameController, g)
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
	if w <= 0 || h <=0 {
		return
	}
	if g.snake == nil || g.snake.NeedReset() {
		g.snake = NewSnake(coords.NewPoint(w/2, h/2))

	}
	if g.Board != nil {
		return
	}
	g.Board = board.NewBoard(w, h)
	g.food = board.NewFoodRandom(w,h)
	ch <- true
}

var (
	ch = make(chan bool)
)

func (g *Game) Animate() {
	go func() {
		<-ch
		for {
			time.Sleep(g.snake.duration - time.Duration(len(g.snake.body)-1))
			if !g.Paused {
				g.snake.Play(g)
				g.genText.SetText(fmt.Sprint("score: ", len(g.snake.body)-1))
			}
			g.Refresh()
		}
	}()
}

