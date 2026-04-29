package main

import (
	"github.com/edi/cli"
	"github.com/edi/state"
)

func main() {
	board := state.InitialState()
	cli.PrintState(&board)
}
