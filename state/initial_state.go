package state

func InitialState() Board {
	board := Board{}

	// Set the white queens.
	board.White.Flag(Pos(3, 0))
	board.Occupancy.Flag(Pos(3, 0))
	board.White.Flag(Pos(0, 3))
	board.Occupancy.Flag(Pos(0, 3))
	board.White.Flag(Pos(0, 6))
	board.Occupancy.Flag(Pos(0, 6))
	board.White.Flag(Pos(3, 9))
	board.Occupancy.Flag(Pos(3, 9))

	// Set the black queens.
	board.Black.Flag(Pos(6, 0))
	board.Occupancy.Flag(Pos(6, 0))
	board.Black.Flag(Pos(9, 3))
	board.Occupancy.Flag(Pos(9, 3))
	board.Black.Flag(Pos(9, 6))
	board.Occupancy.Flag(Pos(9, 6))
	board.Black.Flag(Pos(6, 9))
	board.Occupancy.Flag(Pos(6, 9))

	return board
}
