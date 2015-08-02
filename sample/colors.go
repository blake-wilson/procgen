package main

import (
	"fmt"
	"strconv"

	"github.com/blake-wilson/procgen/format"
	"github.com/blake-wilson/procgen/gen"
)

func colorGen() {
	g := gen.NewGenerator()

	g.RegisterHeightGenerator(
		func(tiles ...*gen.Tile) uint64 {
			// add 5 to the height of the previous tile
			return tiles[len(tiles)-1].Stop.Y + 5
		})

	colorGenFunc := func(t ...*gen.Tile) string {
		// Increment hex color value
		lastTile := t[len(t)-1]
		newHexVal, _ := strconv.ParseInt(lastTile.Attributes["color"].(string), 16, 64)
		newHexVal += 500
		newHexVal = newHexVal % 16777215
		return fmt.Sprintf("%06X", newHexVal)
	}
	g.RegisterGenerator("color", colorGenFunc)

	g.AddTile(&gen.Tile{
		Start: gen.Coord{
			X: 0,
			Y: 0,
		},
		Stop: gen.Coord{
			X: 20,
			Y: 20,
		},
		Attributes: map[string]interface{}{"color": "000000"},
	})

	for i := 0; i < 100; i++ {
		g.Step(20)
	}

	tiles := g.GetTiles()
	csv := format.SprintCSV(tiles...)
	fmt.Printf(csv)
}
