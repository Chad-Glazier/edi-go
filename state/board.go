package state

// Represents a board state.
type Board struct {
	// The occupied squares on the board. An occupied square is one that has
	// either a queen or an arrow on it.
	Occupancy BitBoard
	// The positions of the black queens on the board.
	Black BitBoard
	// The positions of the white queens on the board.
	White BitBoard
	// The most recent move made.
	Move Move
}

// Returns true if and only if White gets to make the next move.
func (board *Board) WhiteIsActive() bool {
	// If the occupancy board has an even number of flags (i.e., sum of arrows
	// and queens), then the active player is White. Otherwise, the active
	// player is Black.
	return board.Occupancy.Count()%2 == 0
}

// Returns true if and only if Black gets to make the next move.
func (board *Board) BlackIsActive() bool {
	return board.Occupancy.Count()%2 != 0
}

// Applies a move to the board, mutating its state.
func (board *Board) Apply(move *Move) {
	if board.Black.Flagged(move.From) {
		board.Black.Unflag(move.From)
		board.Black.Flag(move.To)
	} else {
		board.White.Unflag(move.From)
		board.White.Flag(move.To)
	}
	board.Occupancy.Unflag(move.From)
	board.Occupancy.Flag(move.To)
	board.Occupancy.Flag(move.Arrow)
}
