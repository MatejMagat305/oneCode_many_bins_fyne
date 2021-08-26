package coords

type Basic struct {
	X, Y int
}

func (b *Basic) GetCoords() (int, int) {
	return b.X, b.Y
}

func (b *Basic) setCoords(x,y int)  {
	b.X = x
	b.Y = y
}

