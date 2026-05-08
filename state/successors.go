package state

import "github.com/Chad-Glazier/edi/bb"

// This constant dictates the initial capacity of the slice that holds child
// states. In most Amazons board states, there are 1000 or fewer moves, but in
// the early game it can be between 2000-3000. In the early turns, reallocating
// the successor slice can be expensive but over-allocating it in the mid- to
// late-game seems to incur a much more significant cost. The ideal capacity
// is clearly variable, so we should consider computing it by some heuristic
// before the search. For now, though, we just set it to a number that seems
// good after some trial-and-error with the benchmarks.
const SUCCESSOR_INITIAL_CAPACITY = 200

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
