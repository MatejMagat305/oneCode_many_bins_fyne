package coords

import "math"

type Point struct {
	Basic
}

func (p *Point) SetCoords(x,y int)  {
	x, y = abs(x), abs(y)
	p.setCoords(x, y)
}

func (p *Point) DistanceOf(x int, y int) int {
	return int(math.Sqrt(math.Pow(float64(p.X-x),2)+math.Pow(float64(p.Y-y),2)))
}

func (p *Point) Collision(p2 *Point) bool {
	return p.X == p2.X && p.Y ==p2.Y
}


func NewPoint(x, y int) *Point {
	return &Point{
		Basic: Basic{x,y},
	}
}