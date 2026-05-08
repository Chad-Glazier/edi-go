package state

import "github.com/Chad-Glazier/edi/bb"

// Represents a player color.
type PlayerColor byte

const (
	WHITE PlayerColor = 0 // Represents the player on White
	BLACK PlayerColor = 1 // Represents the player on Black
)

// Represents a board state.
type Board struct {
	// The occupied squares on the board. An occupied square is one that has
	// either a queen or an arrow on it.
	Occupancy bb.BitBoard
	// The positions of the black queens on the board.
	Black [4]bb.Position
	// The positions of the white queens on the board.
	White [4]bb.Position
	// The most recent move made.
	Move Move
	// The player who can make the next move.
	Player PlayerColor
}
