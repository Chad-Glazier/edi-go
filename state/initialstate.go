package state

import "github.com/edi/bb"

func InitialState() Board {
	board := Board{}

	// Set the white queens.
	board.White.Flag(bb.Pos(3, 0))
	board.Occupancy.Flag(bb.Pos(3, 0))
	board.White.Flag(bb.Pos(0, 3))
	board.Occupancy.Flag(bb.Pos(0, 3))
	board.White.Flag(bb.Pos(0, 6))
	board.Occupancy.Flag(bb.Pos(0, 6))
	board.White.Flag(bb.Pos(3, 9))
	board.Occupancy.Flag(bb.Pos(3, 9))

	// Set the black queens.
	board.Black.Flag(bb.Pos(6, 0))
	board.Occupancy.Flag(bb.Pos(6, 0))
	board.Black.Flag(bb.Pos(9, 3))
	board.Occupancy.Flag(bb.Pos(9, 3))
	board.Black.Flag(bb.Pos(9, 6))
	board.Occupancy.Flag(bb.Pos(9, 6))
	board.Black.Flag(bb.Pos(6, 9))
	board.Occupancy.Flag(bb.Pos(6, 9))

	return board
}
