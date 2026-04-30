package state

// Represents a player color.
type PlayerColor byte

const (
	WHITE PlayerColor = 0  // Represents the player on White
	BLACK PlayerColor = 1 // Represents the player on Black
)

// This constant dictates the initial capacity of the slice that holds child
// states. In most Amazons board states, there are 1000 or fewer moves, but in
// the early game it can be between 2000-3000. Reallocating the slice is pretty
// expensive, so we want to avoid it when possible and keep the capacity high.
// However, the successor allocations account for a very large portion of all
// memory allocations and we may want to keep this number lower to help that. 
// In some rough test runs, it looks like a capacity of ~300 makes the program
// use ~450MB of memory (at peak), while using 3000 (which avoids any 
// reallocations) peaks at ~1GB.
const SUCCESSOR_INITIAL_CAPACITY = 3000

const (
	W  int = iota // West
	NW            // Northwest
	N             // North
	NE            // Northeast
	E             // East
	SE            // Southeast
	S             // South
	SW            // Southwest
)

// The last direction; i.e., the one with the greatest underlying numeric
// value. If you want to iterate over all directions, you can iterate over
// range NUMBER_OF_DIRECTIONS.
const NUMBER_OF_DIRECTIONS = 8

// Represents a position on the 10x10 Amazons board with an index from 0 to 99.
// We use row-major ordering, so you can get the row index with position / 10
// and the column with position % 10.
type Position uint8

// Represents a null position. I.e., for functions that return a position,
// the null position should be returned if no valid position exists.
const NULL_POS Position = 100

// Converts row and column indices into a position index.
func Pos(row, col int) Position {
	return Position(row*10 + col)
}

// Converts a position index into row and column coordinates
func Coords(pos Position) (row, col int) {
	row = int(pos) / 10
	col = int(pos) % 10
	return
}

// Represents a move.
type Move struct {
	// The original position of the queen being moved.
	From Position
	// The new position of the queen being moved.
	To Position
	// The position where the queen fired her arrow.
	Arrow Position
}
