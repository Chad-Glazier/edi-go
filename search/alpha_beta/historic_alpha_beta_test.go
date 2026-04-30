package alpha_beta

import (
	"testing"
	"time"

	"github.com/edi/eval"
	"github.com/edi/state"
)

func BenchmarkHistoricAlphaBeta(b *testing.B) {
	board := state.InitialState()
	search := HistoricAlphaBeta(eval.KMinDist)
	for b.Loop() {
		search(&board, time.Second*5)
	}
}
