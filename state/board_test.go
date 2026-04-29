package state

import (
	"testing"
)

func BenchmarkSuccessors(b *testing.B) {
	board := InitialState()
	for b.Loop() {
		blackHole = board.Successors()
	}
}
