package eval

import (
	"testing"

	"github.com/edi/state"
)

var blackHole any

func BenchmarkKMinDist(b *testing.B) {
	board := state.InitialState()
	for b.Loop() {
		blackHole = KMinDist(&board)
	}
}
