package fields

import (
	"image/color"
	"simple_snake/vars"
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
	result := &Field{
		typ:     typ,
	}
	if result.IsFree() {
		result.SetColor(vars.FreeColor)
	}
	return result
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
