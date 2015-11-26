package data

import (
	"io/ioutil"
	"strings"
)

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

func NewUniverseFromFile(w, h int, filename string) (*Universe, error) {
	universe := NewUniverse(w, h)
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	indice := 0
	lines := strings.Split(string(contents), "\n")

	for i := 0; i < h; i++ {
		if len(lines)-1 < i {
			indice++
			continue
		}
		line := strings.Split(lines[i], " ")

		for j := 0; j < w; j++ {
			if len(line)-1 < j {
				indice++
				continue
			}
			if line[j] == "*" {
				universe.Cells[indice].Live = true
			} else {
				universe.Cells[indice].Live = false
			}
			indice++
		}
	}
	return universe, nil
}

func (u *Universe) Get(x int, y int) Cell {
	x, y = u.GetCorrectedCoordinates(x, y)
	return u.Cells[y*u.Height+x]
}

func (u *Universe) SetLive(x int, y int, value bool) {
	if u.coordinatesOutOfRange(x, y) {
		return
	}
	u.Cells[y*u.Height+x].Live = value
}

func (u *Universe) GetCorrectedCoordinates(x, y int) (int, int) {
	if x < 0 {
		x = u.Width - 1
	} else if x >= u.Width {
		x = 0
	}
	if y < 0 {
		y = u.Height - 1
	} else if y >= u.Height {
		y = 0
	}
	return x, y
}

func (u *Universe) coordinatesOutOfRange(x, y int) bool {
	if x < 0 || x > u.Width-1 || y < 0 || y > u.Height-1 {
		return true
	}
	return false
}

func (u *Universe) String() string {
	output := ""
	for i := range u.Cells {
		if i%u.Width == 0 {
			output = output + "\n"
		}
		if u.Cells[i].Live {
			output = output + "* "
		} else {
			output = output + "- "
		}
	}
	return output
}
