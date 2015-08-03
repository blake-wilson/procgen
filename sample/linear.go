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
		func(tiles ...*gen.Tile) uint64 {
			// add 5 to the height of the previous tile
			return 5
		})

	g.AddTile(&gen.Tile{
		Start: gen.Coord{
			X: 0,
			Y: 0,
		},
		Stop: gen.Coord{
			X: 20,
			Y: 20,
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
