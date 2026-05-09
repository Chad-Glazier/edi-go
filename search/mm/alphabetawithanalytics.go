package mm

import (
	"time"

	"github.com/Chad-Glazier/edi/eval"
	"github.com/Chad-Glazier/edi/state"
)

type AlphaBetaAnalytics struct {
	Depth         uint8
	LeafNodes     uint64
	InteriorNodes uint64
	Duration      time.Duration
}

// Conducts a simple alpha-beta search and collects analytics as it goes. This
// implementation involves more overhead than the regular AlphaBeta function so
// you should only use this version if the analytics are important.
//
// Search analytics are sent through the returned channel at the completion of
// each depth-limited search during the iterative deepening process.
func AlphaBetaWithAnalytics(
	board state.Board,
	timeLimit time.Duration,
	heuristic eval.EvalFunc,
) (*state.Move, <-chan AlphaBetaAnalytics) {
	maxDepth := 100 - board.Occupancy.Count()
	complete := make(chan bool)
	var bestMove *state.Move

	s := &alphaBetaState{
		heuristic: heuristic,
	}

	// TODO: Implement.

	go func() {
		for depth := 1; depth <= maxDepth; depth++ {

			bestChildAtDepth := s.depthLimitedSearch(&board, depth)
			if bestChildAtDepth == nil {
				break
			}

			bestMove = &bestChildAtDepth.Move
		}
		complete <- true
	}()

	select {
	case <-time.After(timeLimit):
		return bestMove, nil
	case <-complete:
		return bestMove, nil
	}
}
