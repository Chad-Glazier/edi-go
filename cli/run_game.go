package cli

import (
	"fmt"
	"time"

	"github.com/edi/search"
	"github.com/edi/state"
)

func RunGame(white, black search.SearchFunc, turnTimer time.Duration) {
	board := state.InitialState()
	clearScreen()
	PrintState(&board)

	player := state.WHITE

	for len(board.Successors()) != 0 {
		var move *state.Move
		if player == state.WHITE {
			move = white(&board, turnTimer)
		} else {
			move = black(&board, turnTimer)
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
