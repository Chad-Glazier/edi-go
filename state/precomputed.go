package state

// For each position index p, kAdj[p] stores a precomputed bitboard that has
// each square q flagged if and only if a chess king on p could get to q in a
// single move.
var kAdj = [100]BitBoard{}

// For each position index p and each direction d, rayExc[p][d] stores a
// precomputed bitboard such that each position is flagged if and only if it
// lies on a ray projected from p in the direction of d. This ray will exclude
// p (see rayInc for one that includes it).
var rayExc = [100][8]BitBoard{}

// For each position index p and each direction d, rayInc[p][d] stores a
// precomputed bitboard such that each position is flagged if and only if it
// lies on a ray projected from p in the direction of d. This ray will include
// p (see rayExc for one that excludes it).
var rayInc = [100][8]BitBoard{}

func init() {

	// Precompute the bitboards.
	for row := range 10 {
		for col := range 10 {
			kAdj[row*10+col] = kAdjacent(row, col)
			for dir := range NUMBER_OF_DIRECTIONS {
				rayExc[row*10+col][dir] = exclusiveRay(row, col, dir)
				rayInc[row*10+col][dir] = inclusiveRay(row, col, dir)
			}
		}
	}
}

func kAdjacent(row, col int) BitBoard {
	bb := BitBoard{}

	if row != 9 {
		bb.Flag(Pos(row+1, col))

		if col != 9 {
			bb.Flag(Pos(row+1, col+1))
		}

		if col != 0 {
			bb.Flag(Pos(row+1, col-1))
		}
	}

	if row != 0 {
		bb.Flag(Pos(row-1, col))

		if col != 9 {
			bb.Flag(Pos(row-1, col+1))
		}

		if col != 0 {
			bb.Flag(Pos(row-1, col-1))
		}
	}

	if col != 9 {
		bb.Flag(Pos(row, col+1))
	}

	if col != 0 {
		bb.Flag(Pos(row, col-1))
	}

	return bb
}

func exclusiveRay(row, col, direction int) BitBoard {
	bb := BitBoard{}

	// Rows are indexed from the top to the bottom. So, to move "north," we
	// would need to decrement the row index.

	// Column indices are indexed left-to-right, so incrementing the column
	// index is the same as moving "east."

	for {
		switch direction {
		case N:
			row--
		case NE:
			row--
			col++
		case E:
			col++
		case SE:
			row++
			col++
		case S:
			row++
		case SW:
			row++
			col--
		case W:
			col--
		case NW:
			row--
			col--
		}
		if row >= 10 || row < 0 || col >= 10 || col < 0 {
			break
		}

		bb.Flag(Pos(row, col))
	}

	return bb
}

func inclusiveRay(row, col, direction int) BitBoard {
	bb := exclusiveRay(row, col, direction)
	bb.Flag(Pos(row, col))
	return bb
}
