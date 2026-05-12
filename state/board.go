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

type PositionStatus uint8

const (
	VACANT PositionStatus = iota
	WHITE_QUEEN
	BLACK_QUEEN
	ARROW
)

// Returns VACANT, WHITE_QUEEN, BLACK_QUEEN, or ARROW, depending on what the
// status of the position is on the board. This function is not optimal and is
// only provided for convenience.
func (b *Board) Status(pos bb.Position) PositionStatus {
	if !b.Occupancy.Flagged(pos) {
		return VACANT
	}

	for i := range 4 {
		if b.White[i] == pos {
			return WHITE_QUEEN
		}
		if b.Black[i] == pos {
			return BLACK_QUEEN
		}
	}

	return ARROW
}
