package state

// Returns an unordered slice of all possible subsequent board states.
func (board *Board) Successors() []Board {

	// At most, there are roughly possible moves in an Amazons board state.
	// When we recursively expand child nodes, we still only have as many
	// children expanded as the search tree is deep. Since the search tree only
	// gets to a depth of about 3-10 for the vast majority of the game, we can
	// afford to allocate the full maximum width; sacrificing this meager
	// amount of memory to ensure that we dont have to do any (slow) slice
	// reallocations.
	successors := make([]Board, 0, 3000)

	if board.Player == WHITE {
		i1 := board.White
		for from := i1.Next(); from != NULL_POS; from = i1.Next() {

			i2 := QNeighbors(board.Occupancy, from)
			for to := i2.Next(); to != NULL_POS; to = i2.Next() {

				board.White.Unflag(from)
				board.White.Flag(to)

				board.Occupancy.Unflag(from)
				board.Occupancy.Flag(to)

				i3 := QNeighbors(board.Occupancy, to)
				for arrow := i3.Next(); arrow != NULL_POS; arrow = i3.Next() {

					board.Occupancy.Flag(arrow)

					successors = append(successors, Board{
						Occupancy: board.Occupancy,
						White:     board.White,
						Black:     board.Black,
						Player:    BLACK,
						Move: Move{
							From:  from,
							To:    to,
							Arrow: arrow,
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
		i1 := board.Black
		for from := i1.Next(); from != NULL_POS; from = i1.Next() {

			i2 := QNeighbors(board.Occupancy, from)
			for to := i2.Next(); to != NULL_POS; to = i2.Next() {

				board.Black.Unflag(from)
				board.Black.Flag(to)

				board.Occupancy.Unflag(from)
				board.Occupancy.Flag(to)

				i3 := QNeighbors(board.Occupancy, to)
				for arrow := i3.Next(); arrow != NULL_POS; arrow = i3.Next() {

					board.Occupancy.Flag(arrow)

					successors = append(successors, Board{
						Occupancy: board.Occupancy,
						White:     board.White,
						Black:     board.Black,
						Player:    WHITE,
						Move: Move{
							From:  from,
							To:    to,
							Arrow: arrow,
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
