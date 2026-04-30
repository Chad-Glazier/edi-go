package state

// Returns an unordered slice of all possible subsequent board states.
func (board *Board) Successors() []Board {

	successors := make([]Board, 0, SUCCESSOR_INITIAL_CAPACITY)

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
