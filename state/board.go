package state

import "github.com/edi/bb"

// Represents a player color.
type PlayerColor byte

const (
	WHITE PlayerColor = 0  // Represents the player on White
	BLACK PlayerColor = 1 // Represents the player on Black
)

// Represents a move.
type Move struct {
	// The original position of the queen being moved.
	From bb.Position
	// The new position of the queen being moved.
	To bb.Position
	// The position where the queen fired her arrow.
	Arrow bb.Position
}

// Represents a board state.
type Board struct {
	// The occupied squares on the board. An occupied square is one that has
	// either a queen or an arrow on it.
	Occupancy bb.BitBoard
	// The positions of the black queens on the board.
	Black bb.BitBoard
	// The positions of the white queens on the board.
	White bb.BitBoard
	// The most recent move made.
	Move Move
	// The player who can make the next move.
	Player PlayerColor
}

// Applies a move to the board, mutating its state.
func (board *Board) Apply(move *Move) {
	if board.Black.Flagged(move.From) {
		board.Black.Unflag(move.From)
		board.Black.Flag(move.To)
		board.Player = WHITE
	} else {
		board.White.Unflag(move.From)
		board.White.Flag(move.To)
		board.Player = BLACK
	}
	board.Occupancy.Unflag(move.From)
	board.Occupancy.Flag(move.To)
	board.Occupancy.Flag(move.Arrow)
}
