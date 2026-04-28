package state

import "math/bits"

type BitBoard struct {
	hi uint64
	lo uint64
}

func (bb *BitBoard) Copy() BitBoard {
	return BitBoard{
		lo: bb.lo,
		hi: bb.hi,
	}
}

func (bb *BitBoard) Flag(pos Position) {
	if pos < 64 {
		bb.lo |= 1 << pos
	} else {
		bb.hi |= 1 << (pos - 64)
	}
}

func (bb *BitBoard) Unflag(pos Position) {
	if pos < 64 {
		bb.lo = bb.lo &^ (1 << pos)
	} else {
		bb.hi = bb.hi &^ (1 << (pos - 64))
	}
}

func (bb *BitBoard) Flagged(pos Position) bool {
	if pos < 64 {
		return bb.lo & (1 << pos) != 0
	} else {
		return bb.hi & (1 << (pos - 64)) != 0
	}
}

func (bb *BitBoard) Empty() bool {
	return bb.lo == 0 && bb.hi == 0
}

// Returns the "lowest" position on the board, meaning that which is the
// closest to the bottom-right corner, and unflags it. If the bitboard is
// empty, then NULL_POS is returned.
func (bb *BitBoard) Next() Position {
	switch {
	case bb.hi != 0:
		pos := Position(bits.TrailingZeros64(bb.hi) + 64)
		bb.hi &= bb.hi - 1
		return pos
	case bb.lo != 0:
		pos := Position(bits.TrailingZeros64(bb.lo))
		bb.lo &= bb.lo - 1
		return pos
	default:
		return NULL_POS
	}
}

// Returns the number of flagged positions on this board.
func (bb *BitBoard) Count() int {
	return bits.OnesCount64(bb.lo) + bits.OnesCount64(bb.hi)
}
