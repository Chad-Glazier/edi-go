package mm

import (
	"testing"
	"time"

	"github.com/Chad-Glazier/edi/eval"
	"github.com/Chad-Glazier/edi/state"
)

var initialBoard state.Board
var randomBoard10 state.Board

func init() {
	initialBoard = state.InitialState()
	randomBoard10 = state.RandomBoard(10)
}

func BenchmarkAlphaBetaInitial(b *testing.B) {
	AlphaBeta(initialBoard, time.Second*3, eval.KMinDist)
}

func BenchmarkABWithAnalyticsInitial(b *testing.B) {
	AlphaBetaWithAnalytics(initialBoard, time.Second*3, eval.KMinDist)
}

func BenchmarkAlphaBetaRandomTurn10(b *testing.B) {
	AlphaBeta(randomBoard10, time.Second*3, eval.KMinDist)
}

func BenchmarkABWithAnalyticsRandomTurn10(b *testing.B) {
	AlphaBetaWithAnalytics(initialBoard, time.Second*3, eval.KMinDist)
}
