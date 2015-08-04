package main

import (
	"fmt"

	"github.com/blake-wilson/procgen/format"
	"github.com/blake-wilson/procgen/gen"
)

// Samples are used to test the API

func linearGen() {
	// create a map Generator
	g := gen.NewGenerator()

	g.RegisterHeightGenerator(
		func(tiles ...*gen.Tile) int64 {
			// add 5 to the height of the previous tile
			return 5
		})

	g.AddTile(&gen.Tile{
		Width:  20,
		Height: 20,
		Location: gen.Coord{
			X: 0,
			Y: 0,
		},
	})
	// generate 100 tiles - use width of 20
	for i := 0; i < 100; i++ {
		g.Step(20)
	}

	tiles := g.GetTiles()
	csv := format.SprintCSV(tiles...)
	fmt.Printf(csv)
}
