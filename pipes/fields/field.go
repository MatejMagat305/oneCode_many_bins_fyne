package fields

import (
	"image/color"
)

const (
	Free = rune(iota+1)
	stone
)

type Field struct {
	typ     rune
	c *color.RGBA
}

func NewField(typ rune) *Field {
	return &Field{
		typ:     typ,
	}
}

func (f *Field) SetColor(c *color.RGBA) {
	f.c = c
}

func (f *Field) GetColor() *color.RGBA {
	return f.c
}

func (f *Field) IsFree() bool {
	return f.typ == Free
}
