package main

import (
	"fmt"

	"github.com/blake-wilson/procgen/gen"
)

// Samples are used to test the API

func main() {
	// create a map Generator
	g := gen.NewGenerator()

	g.RegisterHeightGenerator(
		func(tiles ...*gen.Tile) uint64 {
			// add 5 to the height of the previous tile
			return tiles[len(tiles)-1].Stop.Y + 5
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
	fmt.Printf("%v, %v\n", tiles[0].Start.X, tiles[0].Start.Y)
	for _, tile := range tiles {
		// print out the only values interested in: the positions of the
		// tiles. This generator did not generate any other attributes
		fmt.Printf("%v, %v\n", tile.Stop.X, tile.Stop.Y)
	}
}
