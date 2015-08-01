package gen

type Coord struct {
	X uint64
	Y uint64
}

type Tile struct {
	Start Coord
	Stop  Coord

	// The attributes belonging to the terrain.  Indexed by the attribute
	// name - maps to the arbitrary attribute value
	Attributes map[string]interface{}
}

func (t *Tile) Width() uint64 {
	return t.Stop.X - t.Start.X
}

func NewTile(left Coord, right Coord, attributes map[string]interface{}) *Tile {
	if right.X <= left.X {
		// left coordinate must precede right
		return nil
	}

	return &Tile{
		Start:      left,
		Stop:       right,
		Attributes: attributes,
	}
}

type Map struct {
	// Represents a continuous block of terrain

	// Slice of terrain in order from starting x-coordinate to ending x-coordinate
	Tiles []Tile
}

// Gen function defines a function based on the context of the map up to
// the current point.  A gen function outputs an attribute value for the
// next tile with inputs consisting of all previous tiles.
type GenFunction func(...*Tile) string

func NewMap(tiles ...Tile) *Map {
	return &Map{
		Tiles: tiles,
	}
}
