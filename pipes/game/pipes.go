package game

import (
	"pipes/coords"
	"time"
)

type Pipe struct {
	body      []*coords.PointDirect
	direct    *coords.Direct
	duration  time.Duration
	needReset bool
	ch        chan bool
}

func (p *Pipe) Play(g *Game) {
	if g.Paused {
		return
	}
	p.move(g)
}

func (p *Pipe) move(g *Game) {
	b := p.body[len(p.body)-1]
	p.body = append(p.body, coords.NewPointDirect(&coords.Point{coords.Basic{
		X: b.P.X + b.D.X,
		Y: b.P.Y + b.D.Y,
	}}, p.direct, nil))
	x, y := b .P.X, b.P.Y
	if x < 0 || y < 0 || x >= len(*g.Board) || y >= len((*g.Board)[0]) || p.colision() {
		g.Paused = true
		g.Lose()
		g.PauseButton.Text = "try again"
		g.PauseButton.OnTapped = func() {
			p.needReset = true
			g.PauseButton.OnTapped = g.normalButtonFunction(g.PauseButton)
			g.PauseButton.Refresh()
		}
		return
	}
	p.duration-=10
}

func (p *Pipe) NeedReset() bool {
	return p.needReset
}

func (p *Pipe) colision() bool {
	lenght := len(p.body)-1
	b := p.body[lenght]
	for i := 0; i < lenght; i++ {
		if p.body[i].P.Collision(b.P) {
			return true
		}
	}
	return false
}

func NewSnake(head *coords.Point) *Pipe {
	d := &coords.Direct{}
	d.Random()
	body := make([]*coords.PointDirect,0,1000)
	body = append(body, coords.NewPointDirect(head, d, nil))
	return &Pipe{
		body:      body,
		direct:    d,
		duration:  time.Second/10,
		needReset: false,
		ch:        make(chan bool),
	}
}
