package state

import "github.com/Chad-Glazier/edi/bb"

// This constant dictates the initial capacity of the slice that holds child
// states. In most Amazons board states, there are 1000 or fewer moves, but in
// the early game it can be between 2000-3000. Reallocating the slice is pretty
// expensive, so we want to avoid it when possible and keep the capacity high.
// However, the successor allocations account for a very large portion of all
// memory allocations and we may want to keep this number lower to help that.
const SUCCESSOR_INITIAL_CAPACITY = 300

// Returns an unordered slice of all possible subsequent board states.
func (board *Board) Successors() []Board {

	successors := make([]Board, 0, SUCCESSOR_INITIAL_CAPACITY)

	if board.Player == WHITE {
		for queenIdx, from := range board.White {

			i2 := QNeighbors(board.Occupancy, from)
			for to := i2.Next(); to != bb.NULL_POS; to = i2.Next() {

				board.White[queenIdx] = to

				board.Occupancy.Unflag(from)
				board.Occupancy.Flag(to)

				i3 := QNeighbors(board.Occupancy, to)
				for arrow := i3.Next(); arrow != bb.NULL_POS; arrow = i3.Next() {

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

				board.White[queenIdx] = from

				board.Occupancy.Flag(from)
				board.Occupancy.Unflag(to)
			}
		}
	} else {
		for queenIdx, from := range board.Black {

			i2 := QNeighbors(board.Occupancy, from)
			for to := i2.Next(); to != bb.NULL_POS; to = i2.Next() {

				board.Black[queenIdx] = to

				board.Occupancy.Unflag(from)
				board.Occupancy.Flag(to)

				i3 := QNeighbors(board.Occupancy, to)
				for arrow := i3.Next(); arrow != bb.NULL_POS; arrow = i3.Next() {

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

				board.Black[queenIdx] = from

				board.Occupancy.Flag(from)
				board.Occupancy.Unflag(to)
			}
		}
	}

	return successors
}
