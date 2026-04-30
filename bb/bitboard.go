package bb

import "math/bits"

// Represents a position on the 10x10 Amazons board with an index from 0 to 99.
// We use row-major ordering, so you can get the row index with position / 10
// and the column with position % 10.
type Position uint8

// Represents a null position. I.e., for functions that return a position,
// the null position should be returned if no valid position exists.
const NULL_POS Position = 100

// Converts row and column indices into a position index.
func Pos(row, col int) Position {
	return Position(row*10 + col)
}

// Converts a position index into row and column coordinates
func Coords(pos Position) (row, col int) {
	row = int(pos) / 10
	col = int(pos) % 10
	return
}

// Represents a board where each position index (0-99, since Amazons is played
// on a 10x10 board) is either 0 or 1, which we refer to as "unflagged" and
// "flagged," respectively.
type BitBoard struct {
	hi uint64
	lo uint64
}

// Flags a bit in the bitboard.
func (bb *BitBoard) Flag(pos Position) {
	if pos < 64 {
		bb.lo |= 1 << pos
	} else {
		bb.hi |= 1 << (pos - 64)
	}
}

// Unflags a bit in the bitboard.
func (bb *BitBoard) Unflag(pos Position) {
	if pos < 64 {
		bb.lo = bb.lo &^ (1 << pos)
	} else {
		bb.hi = bb.hi &^ (1 << (pos - 64))
	}
}

// Returns true if the bit in the board is flagged and false otherwise.
func (bb *BitBoard) Flagged(pos Position) bool {
	if pos < 64 {
		return bb.lo&(1<<pos) != 0
	} else {
		return bb.hi&(1<<(pos-64)) != 0
	}
}

// Returns true if and only if the bitboard has no flags.
func (bb *BitBoard) Empty() bool {
	return bb.lo == 0 && bb.hi == 0
}

// Returns true if and only if the bitboard has at least one flag.
func (bb *BitBoard) NotEmpty() bool {
	return bb.lo != 0 || bb.hi != 0
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

// Returns the greatest flagged position index on the board. If the
// board is empty, then the null position (NULL_POS) is returned.
func (bb *BitBoard) Lsb() Position {
	switch {
	case bb.lo != 0:
		return Position(bits.TrailingZeros64(bb.lo))
	case bb.hi != 0:
		return Position(64 + bits.TrailingZeros64(bb.hi))
	default:
		return NULL_POS
	}
}

// Returns the position index of the most-significant bit in the board. If the
// board is empty, then the null position (NULL_POS) is returned .
func (bb *BitBoard) Msb() Position {
	switch {
	case bb.hi != 0:
		return Position(127 - bits.LeadingZeros64(bb.hi))
	case bb.lo != 0:
		return Position(63 - bits.LeadingZeros64(bb.lo))
	default:
		return NULL_POS
	}
}

// Performs a bitwise OR operation (a | b) and returns the result.
func (a BitBoard) Or(b BitBoard) BitBoard {
	return BitBoard{
		lo: a.lo | b.lo,
		hi: a.hi | b.hi,
	}
}

// Performs a bitwise OR operation (a | b) and assigns the result to a.
func (a *BitBoard) AssignOr(b BitBoard) {
	a.lo |= b.lo
	a.hi |= b.hi
}

// Performs a bitwise XOR operation (a ^ b) and returns the result.
func (a BitBoard) Xor(b BitBoard) BitBoard {
	return BitBoard{
		lo: a.lo ^ b.lo,
		hi: a.hi ^ b.hi,
	}
}

// Performs a bitwise XOR operation (a ^ b) and assigns the result to a.
func (a *BitBoard) AssignXor(b BitBoard) {
	a.lo ^= b.lo
	a.hi ^= b.hi
}

// Performs a bitwise AND operation (a & b) and returns the result.
func (a BitBoard) And(b BitBoard) BitBoard {
	return BitBoard{
		lo: a.lo & b.lo,
		hi: a.hi & b.hi,
	}
}

// Performs a bitwise AND operation (a & b) and assigns the result to a.
func (a *BitBoard) AssignAnd(b BitBoard) {
	a.lo &= b.lo
	a.hi &= b.hi
}

// Performs a bitwise AND-NOT operation (a &^ b) and returns the result.
func (a BitBoard) AndNot(b BitBoard) BitBoard {
	return BitBoard{
		lo: a.lo &^ b.lo,
		hi: a.hi &^ b.hi,
	}
}

// Performs a bitwise AND-NOT operation (a &^ b) and assigns the result to a.
func (a *BitBoard) AssignAndNot(b BitBoard) {
	a.lo &^= b.lo
	a.hi &^= b.hi
}

// Performs a bitwise NOT operation (a ^ b) and returns the result.
func (a BitBoard) Not() BitBoard {
	return BitBoard{
		lo: ^a.lo,
		hi: ^a.hi,
	}
}

// Performs a bitwise NOT operation (a ^ b) and assigns the result to a.
func (a *BitBoard) AssignNot() {
	a.lo = ^a.lo
	a.hi = ^a.hi
}
