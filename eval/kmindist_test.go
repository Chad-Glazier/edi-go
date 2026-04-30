package eval

import (
	"testing"

	"github.com/edi/state"
)

func BenchmarkKMinDist(b *testing.B) {
	board := state.InitialState()
	for b.Loop() {
		KMinDist(&board)
	}
}
