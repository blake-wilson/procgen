package format

import (
	"fmt"

	"github.com/blake-wilson/procgen/gen"
)

// Prints each tile as a comma-separated list of attributes beginning with the
// position. x coord, y coord, attr1, attr2, ...
func SprintCSV(tiles ...*gen.Tile) string {
	var str string
	for _, tile := range tiles {
		str += fmt.Sprintf("%v, %v", tile.Location.X, tile.Location.Y)
		for _, attr := range tile.Attributes {
			str += fmt.Sprintf(", %v", attr)
		}
		str += "\n"
	}
	return str
}
