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
		func(tiles ...*gen.Tile) int64 {
			// Add to height based on the least significant digit of
			// the last hex value
			hexVal := tiles[len(tiles)-1].Attributes["color"].(string)
			numVal, _ := strconv.ParseInt(hexVal, 16, 64)
			mask := int64(0x0F)
			var neg int64
			if numVal%2 == 0 {
				neg = 1
			} else {
				neg = -1
			}
			return numVal & mask * neg
		})

	colorGenFunc := func(t ...*gen.Tile) string {
		// Increment hex color value
		lastTile := t[len(t)-1]
		newHexVal, _ := strconv.ParseInt(lastTile.Attributes["color"].(string), 16, 64)
		newHexVal += 500007
		newHexVal = newHexVal % 16777215
		return fmt.Sprintf("%06X", newHexVal)
	}
	g.RegisterGenerator("color", colorGenFunc)

	g.AddTile(&gen.Tile{
		Width:  20,
		Height: 100,
		Location: gen.Coord{
			X: 0,
			Y: 0,
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
