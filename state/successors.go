package state

import "github.com/Chad-Glazier/edi/bb"

// This constant dictates the initial capacity of the slice that holds child
// states. In most Amazons board states, there are 1000 or fewer moves, but in
// the early game it can be between 2000-3000. Reallocating the slice is pretty
// expensive, so we want to avoid it when possible and keep the capacity high.
// However, the successor allocations account for a very large portion of all
// memory allocations and we may want to keep this number lower to help that.
// In some rough test runs, it looks like a capacity of ~300 makes the program
// use ~450MB of memory (at peak), while using 3000 (which avoids any
// reallocations) peaks at ~1GB.
const SUCCESSOR_INITIAL_CAPACITY = 3000

// Returns an unordered slice of all possible subsequent board states.
func (board *Board) Successors() []Board {

	successors := make([]Board, 0, SUCCESSOR_INITIAL_CAPACITY)

	if board.Player == WHITE {
		i1 := board.White
		for from := i1.Next(); from != bb.NULL_POS; from = i1.Next() {

			i2 := QNeighbors(board.Occupancy, from)
			for to := i2.Next(); to != bb.NULL_POS; to = i2.Next() {

				board.White.Unflag(from)
				board.White.Flag(to)

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

				board.White.Flag(from)
				board.White.Unflag(to)

				board.Occupancy.Flag(from)
				board.Occupancy.Unflag(to)
			}
		}
	} else {
		i1 := board.Black
		for from := i1.Next(); from != bb.NULL_POS; from = i1.Next() {

			i2 := QNeighbors(board.Occupancy, from)
			for to := i2.Next(); to != bb.NULL_POS; to = i2.Next() {

				board.Black.Unflag(from)
				board.Black.Flag(to)

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

				board.Black.Flag(from)
				board.Black.Unflag(to)

				board.Occupancy.Flag(from)
				board.Occupancy.Unflag(to)
			}
		}
	}

	return successors
}
