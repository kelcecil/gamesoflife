package data

import (
	"testing"
)

func TestLiveTwoCellSurvival(t *testing.T) {
	universe := NewUniverse(3, 3)
	universe.Cells[0].Live = true
	universe.Cells[1].Live = true

	universe.Cells[4].Live = true

	cell := universe.Get(1, 1)
	if !cell.SurvivesToNextGeneration() {
		t.Fail()
	}
}

func TestLiveThreeCellSurvival(t *testing.T) {
	universe := NewUniverse(3, 3)
	universe.Cells[0].Live = true
	universe.Cells[1].Live = true
	universe.Cells[7].Live = true

	universe.Cells[4].Live = true

	cell := universe.Get(1, 1)
	if !cell.SurvivesToNextGeneration() {
		t.Fail()
	}
}

func TestDeadThreeCellBirth(t *testing.T) {
	universe := NewUniverse(3, 3)
	universe.Cells[0].Live = true
	universe.Cells[1].Live = true
	universe.Cells[7].Live = true

	universe.Cells[4].Live = false

	cell := universe.Get(1, 1)
	if !cell.SurvivesToNextGeneration() {
		t.Fail()
	}
}

func TestLiveOneCellDeath(t *testing.T) {
	universe := NewUniverse(3, 3)
	universe.Cells[2].Live = true

	universe.Cells[4].Live = true

	cell := universe.Get(1, 1)
	if cell.SurvivesToNextGeneration() {
		t.Fail()
	}
}

func TestLiveFourCellDeath(t *testing.T) {
	universe := NewUniverse(3, 3)
	universe.Cells[0].Live = true
	universe.Cells[1].Live = true
	universe.Cells[7].Live = true
	universe.Cells[8].Live = true

	universe.Cells[4].Live = true

	cell := universe.Get(1, 1)
	if cell.SurvivesToNextGeneration() {
		t.Fail()
	}
}
