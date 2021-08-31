package coords

import "math/rand"

type Direct struct {
	Basic
}
func (d *Direct) SetCoords(x,y int)  {
	if x == 0 && y == 0 {
		return
	}
	x,y = changeIfFirstZero(x, y)
	x,y = changeIfFirstZero(y, x)
	d.setCoords(x,y)
}

func (d *Direct) Random() {
	if  randBool(){
		if   randBool(){
			d.SetCoords(1,0)
		}else {
			d.SetCoords(-1,0)
		}
	}else {
		if randBool() {
			d.SetCoords(0,1)
		}else {
			d.SetCoords(0,-1)
		}
	}
}

func (d *Direct) Clone() *Direct {
	return &Direct{Basic{
		X: d.X,
		Y: d.Y,
	}}
}

func randBool() bool {
	return rand.Float32() < 0.5
}
func changeIfFirstZero(n, m int) (int, int) {
	if n != 0 {
		m = 0
		n /= abs(n)
	}
	return n, m
}

func abs(n int) int {
	if n<0{
		return -n
	}
	return n
}

func NewDirect(x, y int) *Direct {
	return &Direct{
		Basic: Basic{x,y},
	}
}