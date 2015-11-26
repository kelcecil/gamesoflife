package data

type Cell struct {
	Live     bool
	Parent   *Universe
	Location Coordinate
}

type Coordinate struct {
	X int
	Y int
}

func (c *Cell) SurvivesToNextGeneration() bool {
	x := c.Location.X
	y := c.Location.Y

	xRange := []int{x - 1, x, x + 1}
	yRange := []int{y - 1, y, y + 1}

	neighbors := 0

	for i := range xRange {
		for j := range yRange {
			workingX := xRange[i]
			workingY := yRange[j]

			if workingX == x && workingY == y {
				continue
			}

			if c.Parent.Get(workingX, workingY).Live == true {
				neighbors++
			}
		}
	}
	if c.Live == false && neighbors == 3 {
		return true
	}
	if c.Live == true && (neighbors == 2 || neighbors == 3) {
		return true
	}
	return false
}
