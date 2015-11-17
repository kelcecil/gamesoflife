package data

import ()

type Universe struct {
	Cells  []Cell
	Width  int
	Height int
}

func NewUniverse(w, h int) *Universe {
	universe := &Universe{
		Width:  w,
		Height: h,
	}

	raw := make([]Cell, w*h)
	for i := range raw {
		y := i / w
		x := i % w
		raw[i] = Cell{
			Location: Coordinate{
				X: x,
				Y: y,
			},
			Live:   false,
			Parent: universe,
		}
	}
	universe.Cells = raw
	return universe
}

func (u *Universe) Get(x int, y int) Cell {
	return u.Cells[y*u.Height+x]
}
