package data

import "testing"

func TestNewUniverse(t *testing.T) {
	universe := NewUniverse(2, 2)
	if len(universe.Cells) != 4 {
		t.Fatalf("Size of this cells array should be 4 but instead is %d", len(universe.Cells))
	}

	if universe.Cells[0].Location.X != 0 && universe.Cells[0].Location.Y != 0 {
		t.Fatalf("The first cell should be at 0,0.")
	}

	if universe.Cells[2].Location.X != 0 && universe.Cells[2].Location.Y != 1 {
		t.Fatalf("The third cell should be at 0,1.")
	}
}

func TestGetCoorectedCoordinates(t *testing.T) {
	universe := NewUniverse(3, 3)
	x, y := universe.GetCorrectedCoordinates(3, 0)
	if x != 0 || y != 0 {
		t.Fail()
	}

	x, y = universe.GetCorrectedCoordinates(-1, 0)
	if x != 2 || y != 0 {
		t.Fail()
	}

	x, y = universe.GetCorrectedCoordinates(0, 3)
	if x != 0 || y != 0 {
		t.Fail()
	}

	x, y = universe.GetCorrectedCoordinates(0, -1)
	if x != 0 || y != 2 {
		t.Fail()
	}
}

func TestNewUniverseFromFile(t *testing.T) {
	universe, err := NewUniverseFromFile(3, 3, "blinker.txt")
	if err != nil {
		t.Fatalf("Failed to create universe from file. Reason: %s", err.Error())
	}
	testCellLocation(t, universe, 1, 0)
	testCellLocation(t, universe, 1, 1)
	testCellLocation(t, universe, 1, 2)
}

func TestUniverseGet(t *testing.T) {
	universe := NewUniverse(2, 2)
	testCellLocation(t, universe, 1, 1)
	testCellLocation(t, universe, 0, 1)
	testCellLocation(t, universe, 1, 0)
}

func testCellLocation(t *testing.T, universe *Universe, x, y int) {
	cell := universe.Get(x, y)
	if cell.Location.X != x && cell.Location.Y != y {
		t.Fatalf("Cell location should be (%d,%d) instead of (%d, %d)",
			x,
			y,
			cell.Location.X,
			cell.Location.Y)
	}
}
