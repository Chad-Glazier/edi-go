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

// Returns an unordered slice of all possible subsequent board states.
func (board *Board) Successors() []Board {

	// At most, there are 2176 possible moves in an Amazons board state.
	// When we recursively expand child nodes, we still only have as many
	// children expanded as the search tree is deep. Since the search tree only
	// gets to a depth of about 3-10 for the vast majority of the game, we can
	// afford to allocate the full maximum width; sacrificing this meager
	// amount of memory for (according to benchmarks) a speedup of about 2x.
	successors := make([]Board, 0, 2176)

	// If the occupancy board has an even number of flags (i.e., sum of arrows
	// and queens), then the active player is White. Otherwise, the active
	// player is Black.

	if board.Occupancy.Count()%2 == 0 {
		i1 := board.White.Copy()
		for from := i1.Next(); from != NULL_POS; from = i1.Next() {

			i2 := QNeighbors(&board.Occupancy, from)
			for to := i2.Next(); to != NULL_POS; to = i2.Next() {

				board.White.Unflag(from)
				board.White.Flag(to)

				board.Occupancy.Unflag(from)
				board.Occupancy.Flag(to)

				i3 := QNeighbors(&board.Occupancy, to)
				for arrow := i3.Next(); arrow != NULL_POS; arrow = i3.Next() {

					board.Occupancy.Flag(arrow)

					successors = append(successors, Board{
						Occupancy: board.Occupancy,
						White:     board.White,
						Black:     board.Black,
						Move: Move{
							from:  from,
							to:    to,
							arrow: arrow,
						},
					})

					board.Occupancy.Unflag(arrow)
				}

				board.White.Flag(from)
				board.White.Unflag(to)

				board.Occupancy.Flag(from)
				board.Occupancy.Unflag(to)
			}
		}
	} else {
		i1 := board.Black.Copy()
		for from := i1.Next(); from != NULL_POS; from = i1.Next() {

			i2 := QNeighbors(&board.Occupancy, from)
			for to := i2.Next(); to != NULL_POS; to = i2.Next() {

				board.Black.Unflag(from)
				board.Black.Flag(to)

				board.Occupancy.Unflag(from)
				board.Occupancy.Flag(to)

				i3 := QNeighbors(&board.Occupancy, to)
				for arrow := i3.Next(); arrow != NULL_POS; arrow = i3.Next() {

					board.Occupancy.Flag(arrow)

					successors = append(successors, Board{
						Occupancy: board.Occupancy,
						White:     board.White,
						Black:     board.Black,
						Move: Move{
							from:  from,
							to:    to,
							arrow: arrow,
						},
					})

					board.Occupancy.Unflag(arrow)
				}

				board.Black.Flag(from)
				board.Black.Unflag(to)

				board.Occupancy.Flag(from)
				board.Occupancy.Unflag(to)
			}
		}
	}

	return successors
}
