package game

import (
	"simple_snake/board"
	"simple_snake/coords"
	"time"
)

type Snake struct {
	body      []*coords.Point
	direct    *coords.Direct
	duration  time.Duration
	needReset bool
	ch        chan bool
}

func (s *Snake) Play(g *Game) {
	if g.Paused {
		return
	}
	g.lock.Lock()
	s.move(g)
	g.lock.Unlock()
}

func (s *Snake) move(g *Game) {
	b := s.body[len(s.body)-1]
	x,y := b.X, b.Y
	if x < 0 || y < 0 || x >= len(*g.Board) || y >= len((*g.Board)[0]) || s.collision() {
		g.Paused = true
		g.Lose()
		g.PauseButton.Text = "try again"
		g.PauseButton.OnTapped = func() {
			s.needReset = true
			g.PauseButton.OnTapped = g.normalButtonFunction(g.PauseButton)
		}
		g.PauseButton.Refresh()
		g.Refresh()
		return
	}
	if g.food.IsThere(x,y) {
		s.body = append(s.body, coords.NewPoint(x+s.direct.X, y+s.direct.Y))
		g.food = board.NewFoodRandom(g.Board.GetSize())
		return
	}
	firstIndex := len(s.body)-1
	for i := 0; i < firstIndex; i++{
		s.body[i]= s.body[i+1]
	}
	s.body[firstIndex] = coords.NewPoint(x+s.direct.X, y+s.direct.Y)

}

func (s *Snake) NeedReset() bool {
	return s.needReset
}

func (s *Snake) collision() bool {
	lenght := len(s.body)-1
	b := s.body[lenght]
	for i := 0; i < lenght; i++ {
		if s.body[i].Collision(b, s.direct) {
			return true
		}
	}
	return false
}

func NewSnake(head *coords.Point) *Snake {
	d := &coords.Direct{}
	d.Random()
	body := make([]*coords.Point,0,1000)
	body = append(body, head)
	return &Snake{
		body:      body,
		direct:    d,
		duration:  time.Second/3,
		needReset: false,
		ch:        make(chan bool),
	}
}
