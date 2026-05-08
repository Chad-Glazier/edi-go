package state

import (
	"testing"
)

func TestApplyMove(t *testing.T) {
	board := RandomBoard(12)
	children := board.Successors()


	for _, child := range children {
		// Ensure that applying the move to the initial board yields the child.
		applied, err := Apply(board, child.Move)
		if err != nil {
			t.Errorf("Expected move to be legal %v %s", child.Move, err.Error())
			continue
		}
		if *applied != child {
			t.Errorf("Expected inferred move to yield child %v", child.Move)
		}
	}
}
