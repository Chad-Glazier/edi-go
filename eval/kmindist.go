package eval

import "github.com/edi/state"

// Partitions territory between Black and White based on who can reach a given
// square faster if their queens moved the way that chess kings do, then
// calculates a score based on the sizes of the territories.
func KMinDist(board *state.Board) float64 {

	whiteTerritory := board.White
	blackTerritory := board.Black

	visited := whiteTerritory.Or(blackTerritory)

	blackFrontier := blackTerritory
	whiteFrontier := whiteTerritory

	for blackFrontier.NotEmpty() || whiteFrontier.NotEmpty() {

		// First, we expand the frontiers, omitting any previously explored
		// territory.
		blackFrontier = state.
			KFrontier(board.Occupancy, blackFrontier).
			AndNot(visited)
		whiteFrontier = state.
			KFrontier(board.Occupancy, whiteFrontier).
			AndNot(visited)

		// Next, we let white and black claim their respective territory.
		// Any new territory on the black frontier that is neither on the
		// white territory nor the white frontier is claimed for black, and
		// vice versa.
		whiteTerritory.AssignOr(
			whiteFrontier.AndNot(blackFrontier.Or(blackTerritory)),
		)
		blackTerritory.AssignOr(
			blackFrontier.AndNot(whiteFrontier.Or(whiteTerritory)),
		)

		// Finally, update the "visited" board to reflect that the new
		// frontiers have been explored.
		visited.AssignOr(whiteFrontier.Or(blackFrontier))
	}

	return float64(whiteTerritory.Count()-blackTerritory.Count()) /
		float64(visited.Count()-7)
}
