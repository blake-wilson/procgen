package gen

type Coord struct {
	X int64
	Y int64
}

type Tile struct {
	Width  int64
	Height int64

	Location Coord

	// The attributes belonging to the terrain.  Indexed by the attribute
	// name - maps to the arbitrary attribute value
	Attributes map[string]interface{}
}

func NewTile(width, height int64, location Coord, attributes map[string]interface{}) *Tile {
	if width < 0 {
		// left coordinate must precede right
		return nil
	}

	return &Tile{
		Width:      width,
		Height:     height,
		Location:   location,
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
