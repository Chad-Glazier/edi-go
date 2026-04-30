package state

import "github.com/edi/bb"

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

// Represents a move.
type Move struct {
	// The original position of the queen being moved.
	From bb.Position
	// The new position of the queen being moved.
	To bb.Position
	// The position where the queen fired her arrow.
	Arrow bb.Position
}
