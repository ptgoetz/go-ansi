package color

import "fmt"

type Color struct {
	r int
	g int
	b int
}

func (c Color) R() int {
	return c.r
}

func (c Color) G() int {
	return c.g
}

func (c Color) B() int {
	return c.b
}

func (c Color) HexString() string {
	return fmt.Sprintf("#%02X%02X%02X", c.r, c.g, c.b)
}

func New(r int, g int, b int) *Color {
	return &Color{r: r, g: g, b: b}
}
