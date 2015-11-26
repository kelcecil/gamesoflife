package main

import (
	"flag"
	"fmt"
	"github.com/gosuri/uilive"
	"github.com/kelcecil/wvu-go-gameoflife/data"
	"io"
	"os"
	"time"
)

func drawScreen(writer io.Writer, universe *data.Universe) {
	fmt.Fprintf(writer, universe.String())
}

func evolveUniverse(universe *data.Universe) *data.Universe {
	workingCopy := make([]data.Cell, len(universe.Cells))
	copy(workingCopy, universe.Cells)
	for i := range universe.Cells {
		workingCopy[i].Live = universe.Cells[i].SurvivesToNextGeneration()
	}
	copy(universe.Cells, workingCopy)
	return universe
}

func run(universe *data.Universe) int {
	writer := uilive.New()
	writer.Start()
	defer writer.Stop()

	for {
		drawScreen(writer, universe)
		universe = evolveUniverse(universe)
		time.Sleep(250 * time.Millisecond)
	}

	return 0
}

func main() {
	var (
		width    int
		height   int
		filename string
	)
	flag.IntVar(&width, "width", 10, "Width of universe")
	flag.IntVar(&height, "height", 10, "Height of universe")
	flag.StringVar(&filename, "filename", "", "Filename containing initial universe.")
	flag.Parse()

	universe, err := data.NewUniverseFromFile(width, height, filename)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	os.Exit(run(universe))
}
