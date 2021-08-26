package coords

import "image/color"

type PointDirect struct {
	P *Point
	D *Direct
	C *color.RGBA
	RP *Point
}

func (p *PointDirect) Update() {
	p.P.X += p.D.X
	p.P.Y += p.D.Y
}

func NewPointDirect(point *Point, d *Direct, rP *Point ) *PointDirect {
	return &PointDirect{
		P: point,
		D: d,
		RP: rP,
	}
}