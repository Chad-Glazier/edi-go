package state

// Returns a bitboard where each neighbor of a given position is flagged, where
// two squares p and q are neighbors if and only if a chess king could move
// from p to q (accounting for squares that are already occupied by arrows or
// queens).
func KNeighbors(occupancy *BitBoard, position Position) BitBoard {
	adjacentSquares := kAdj[position]
	return BitBoard{
		lo: adjacentSquares.lo &^ occupancy.lo,
		hi: adjacentSquares.hi &^ occupancy.hi,
	}
}

// Returns the frontier of a given territory. A territory is a set of positions
// on the board, and the frontier of a territory is defined to be the set of
// positions that are adjacent to some position in the territory, excluding any
// positions that are already in the territory. Two positions p and q are
// adjacent if and only if a chess king could move from p to q in a single
// move, accounting for any arrows or queens that could obstruct such a move.
func KFrontier(occupancy *BitBoard, territory *BitBoard) BitBoard {

	frontier := BitBoard{}

	iter := territory.Copy()
	for pos := iter.Next(); pos != NULL_POS; pos = iter.Next() {

		neighbors := KNeighbors(occupancy, pos)
		frontier.lo |= neighbors.lo
		frontier.hi |= neighbors.hi
	}

	frontier.lo &^= territory.lo
	frontier.hi &^= territory.hi

	return frontier
}

// Returns a bitboard where each neighbor of a given position is flagged, where
// two squares p and q are neighbors if and only if a chess queen could move
// from p to q (accounting for squares that are already occupied by arrows or
// queens).
func QNeighbors(occupancy *BitBoard, position Position) BitBoard {

	occ := occupancy.Copy()
	occ.Unflag(position)

	neighbors := BitBoard{}

	// Iterate over the forward directions.
	for dir := W; dir < E; dir++ {

		ray := rayExc[position][dir]
		blockers := BitBoard{
			lo: ray.lo & occ.lo,
			hi: ray.hi & occ.hi,
		}

		nearestBlocker := blockers.Msb() // the direction is forward
		if nearestBlocker == NULL_POS {
			neighbors.lo |= ray.lo
			neighbors.hi |= ray.hi
			continue
		}

		blockedSegment := rayInc[nearestBlocker][dir]
		neighbors.lo |= ray.lo ^ blockedSegment.lo
		neighbors.hi |= ray.hi ^ blockedSegment.hi

	}

	// Iterate over the backward directions.
	for dir := E; dir <= SW; dir++ {

		ray := rayExc[position][dir]
		blockers := BitBoard{
			lo: ray.lo & occ.lo,
			hi: ray.hi & occ.hi,
		}

		nearestBlocker := blockers.Lsb() // the direction is backward
		if nearestBlocker == NULL_POS {
			neighbors.lo |= ray.lo
			neighbors.hi |= ray.hi
			continue
		}

		blockedSegment := rayInc[nearestBlocker][dir]
		neighbors.lo |= ray.lo ^ blockedSegment.lo
		neighbors.hi |= ray.hi ^ blockedSegment.hi

	}

	return neighbors
}

// Returns the frontier of a given territory. A territory is a set of positions
// on the board, and the frontier of a territory is defined to be the set of
// positions that are adjacent to some position in the territory, excluding any
// positions that are already in the territory. Two positions p and q are
// adjacent if and only if a chess queen could move from p to q in a single
// move, accounting for any arrows or queens that could obstruct such a move.
func QFrontier(occupancy *BitBoard, territory *BitBoard) BitBoard {

	frontier := BitBoard{}

	iter := territory.Copy()
	for pos := iter.Next(); pos != NULL_POS; pos = iter.Next() {

		neighbors := QNeighbors(occupancy, pos)
		frontier.lo |= neighbors.lo
		frontier.hi |= neighbors.hi
	}

	frontier.lo &^= territory.lo
	frontier.hi &^= territory.hi

	return frontier
}
