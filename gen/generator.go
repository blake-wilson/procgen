package gen

import (
	"math/rand"
	"time"
)

type MapGenerator struct {
	// Probabilistically generates terrain based on configuration presets.
	// If no presets are given, assigns a probability to each terrain
	// attribute defined.

	// Attributes represent a name-indexable map of the possible attributes
	// a given tile on the generated map can have.  Each generated tile will
	// then have a subset of these attributes
	Attributes map[string]interface{}

	// Generators is a map indexed with the same values used to index the
	// Attributes map.  Each generator is defined only once, however.
	// Default behavior occurs if a given attribute has no corresponding generator
	Generators map[string]GenFunction

	// Special generator which generates the height delta for the next tile
	heightGen func(...*Tile) uint64

	// Maps each attribute to each other attribute. By indexing the
	// probabilityMatrix with the name of an attribute, a map is returned
	// which can then be indexed by any other attribute name to determine
	// the probability that this attribute will neighbor the attribute whose
	// name is given by the first index.
	probabilityMatrix map[string]map[string]float64

	context []*Tile

	rng *rand.Rand

	// points to the current location of the map generator
	pointer Coord
}

// Reseeds the RNG
func (g *MapGenerator) seedRNG(num int) {
	rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
}

// return all the tiles currently stored in the generator
func (g *MapGenerator) GetTiles() []*Tile {
	return g.context
}

// Generates the next tile and increments the pointer to the current tile location
func (g *MapGenerator) Step(width uint64) *Tile {
	newAttrs := make(map[string]interface{})
	for attr, _ := range g.Attributes {
		newAttrs[attr] = g.Generators[attr](g.context...)
	}
	var delta uint64

	var stop Coord
	stop.X = g.pointer.X + width
	stop.Y = g.pointer.Y + g.heightGen(g.context...)

	newTile := NewTile(g.pointer, stop, newAttrs)
	g.context = append(g.context, newTile)

	// increments the pointer for the whole generator
	g.pointer.X += width
	g.pointer.Y += delta

	return newTile
}

// Registers a generator function to be thereafter used for all newly generated
// tiles
func (g *MapGenerator) RegisterGenerator(attr string, f GenFunction) {
	g.Generators[attr] = f
}

// Unregisters the generator assigned to the given attribute, if it exists
func (g *MapGenerator) UnregisterGenerator(attr string) {
	delete(g.Generators, attr)
}

// Registers the height generator function. Subsequent calls will overwrite the
// generator
func (g *MapGenerator) RegisterHeightGenerator(f func(...*Tile) uint64) {
	g.heightGen = f
}

// Set a tile manually to the next tile position
func (g *MapGenerator) AddTile(t *Tile) {
	g.context = append(g.context, t)
	g.pointer.X += t.Stop.X - t.Start.X

	// Do not allow negative y values
	if t.Height() < 0 && uint64(-1*t.Height()) > g.pointer.Y {
		g.pointer.Y = 0
	} else {
		g.pointer.Y += uint64(t.Height())
	}
}

func NewGenerator() *MapGenerator {
	return &MapGenerator{
		Attributes:        make(map[string]interface{}),
		Generators:        make(map[string]GenFunction),
		context:           make([]*Tile, 0, 4),
		probabilityMatrix: make(map[string]map[string]float64),
		rng:               rand.New(rand.NewSource(time.Now().UTC().UnixNano())),
		pointer: Coord{
			X: 0,
			Y: 0,
		},
	}

}
