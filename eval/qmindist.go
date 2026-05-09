package eval

import (
	"github.com/Chad-Glazier/edi/bb"
	"github.com/Chad-Glazier/edi/state"
)

// Partitions territory between Black and White based on who can reach a given
// square faster if their queens moved the way that chess queens do, then
// calculates a score based on the sizes of the territories.
func QMinDist(board *state.Board) float64 {

	whiteTerritory := bb.BitBoard{}
	blackTerritory := bb.BitBoard{}

	whiteFrontier := bb.BitBoard{}
	blackFrontier := bb.BitBoard{}

	for i := range 4 {
		whiteTerritory.Flag(board.White[i])
		blackTerritory.Flag(board.Black[i])

		whiteFrontier.AssignOr(
			state.QNeighbors(board.Occupancy, board.White[i]))
		blackFrontier.AssignOr(
			state.QNeighbors(board.Occupancy, board.Black[i]))
	}

	visited := whiteTerritory.Or(blackTerritory)

	for blackFrontier.NotEmpty() || whiteFrontier.NotEmpty() {

		// First, we let White and Black claim their respective territory.
		// Any new territory on the White frontier that isn't on the Black
		// frontier is claimed for White, and vice versa.
		whiteTerritory.AssignOr(whiteFrontier.AndNot(blackFrontier))
		blackTerritory.AssignOr(blackFrontier.AndNot(whiteFrontier))

		// Next, update the "visited" board to reflect that the new
		// frontiers have been explored.
		visited.AssignOr(whiteFrontier.Or(blackFrontier))

		// Finally, we expand the frontiers, omitting any previously explored
		// territory.
		blackFrontier = state.
			QFrontier(board.Occupancy, blackFrontier).
			AndNot(visited)
		whiteFrontier = state.
			QFrontier(board.Occupancy, whiteFrontier).
			AndNot(visited)
	}

	return float64(whiteTerritory.Count()-blackTerritory.Count()) /
		float64(visited.Count()-7)
}
