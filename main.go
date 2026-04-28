package main

import (
	"github.com/edi/cli"
	"github.com/edi/state"
)

func main() {
	bb := state.BitBoard{}

	bb.Flag(state.Pos(1, 2))

	kNeighbors := state.QNeighbors(&bb, state.Pos(1, 2))

	cli.PrintBitBoard(&bb)
	cli.PrintBitBoard(&kNeighbors)
}
