package cli

import (
	"fmt"

	"github.com/edi/state"
)

func PrintBitBoard(bb *state.BitBoard) {
	fmt.Println()
	fmt.Println("\t   0 1 2 3 4 5 6 7 8 9 ")
	fmt.Println()

	for i := range 10 {

		fmt.Printf("\t%d ", i)

		for j := range 10 {

			if bb.Flagged(state.Pos(i, j)) {
				fmt.Print(" X")
			} else {
				fmt.Print(" .")
			}

		}

		fmt.Print("\n")
	}

	fmt.Println()
}
