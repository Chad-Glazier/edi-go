package edi

import (
	"time"

	"github.com/Chad-Glazier/edi/eval"
	"github.com/Chad-Glazier/edi/search"
	"github.com/Chad-Glazier/edi/search/mm"
	"github.com/Chad-Glazier/edi/state"
)

// EDI is the flagship VI for this project. At the time of writing, she uses:
// - Minimax search with alpha-beta pruning.
// - The History Heuristic for move ordering.
// - The KMinDist function for leaf node evaluation.
type Arrow struct {
	searchMethod search.SearchFunc
}

func (arrow *Arrow) Consult(
	board *state.Board, timeLimit time.Duration,
) *state.Move {

	if arrow.searchMethod == nil {
		arrow.searchMethod = mm.AlphaBeta(eval.QMinDist)
	}

	return arrow.searchMethod(board, timeLimit)
}
