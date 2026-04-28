package state

// Represents a board state.
type Board struct {
	// The occupied squares on the board. An occupied square is one that has
	// either a queen or an arrow on it.
	occupancy BitBoard
	// The positions of the black queens on the board.
	black [4]Position
	// The positions of the white queens on the board.
	white [4]Position
	// The player who can make a move from this position.
	active PlayerColor
}
