package board

import (
	"pipes/coords"
	"pipes/fields"
	"pipes/vars"
)

type Broard [][]*fields.Field

func (b Broard) MakeColor(directions []*coords.PointDirect) {
	li := len(b)
	for i := 0; i < li; i++ {
		lj := len(b[i])
		min := lj * li
		bi := b[i]
		for j := 0; j < lj; j++ {
			localMin := min
			bj := bi[j]
			for i2, direction := range directions {
				if d := direction.P.DistanceOf(i, j); d < localMin {
					bj.SetColor( vars.C[i2])
					localMin = d
				}
			}
		}
	}
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
