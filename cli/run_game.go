package cli

import (
	"fmt"

	"github.com/edi/search"
	"github.com/edi/state"
)

func RunGame(white, black search.SearchFunc) {
	board := state.InitialState()
	clearScreen()
	PrintState(&board)

	player := state.WHITE

	for len(board.Successors()) != 0 {
		var move *state.Move
		if player == state.WHITE {
			move = white(&board, 30000)
		} else {
			move = black(&board, 30000)
		}
		player = !player
		board.Apply(move)
		clearScreen()
		PrintState(&board)
	}

	if (player == state.WHITE) {
		fmt.Println("Black Wins")
	} else {
		fmt.Println("White Wins")
	}
}