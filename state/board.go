package state

import "github.com/edi/bb"

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
