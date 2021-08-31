package board

import (
	"math/rand"
	"simple_snake/coords"
	"simple_snake/fields"
)

type Broard [][]*fields.Field
type Food coords.Point

func (f *Food) IsThere(x,y int) bool {
	return f.X==x && f.Y==y
}

func NewBoard(w int, h int) *Broard {
	result := make(Broard, 0, w)
	for i := 0; i < w; i++ {
		result = append(result, make([]*fields.Field, 0, h))
		for j := 0; j < h; j++ {
			result[i] = append(result[i], fields.NewField(fields.Free))
		}
	}
	return &result
}

func NewFoodRandom(w int, h int) *Food {
	x, y := rand.Intn(w), rand.Intn(h)
	result :=Food(*coords.NewPoint(x,y))
	return &result
}

func (b Broard) GetSize() (int,int) {
	if len(b) == 0 {
		return 0, 0
	}
	return len(b), len(b[0])
}