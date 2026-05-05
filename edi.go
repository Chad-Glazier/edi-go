package edi

import (
	"time"

	"github.com/Chad-Glazier/edi/eval"
	"github.com/Chad-Glazier/edi/search"
	"github.com/Chad-Glazier/edi/search/mm"
	"github.com/Chad-Glazier/edi/state"
)

// EDI is the flagship VI for this project. At the time of writing she uses
// alpha-beta search with the History Heuristic for move ordering and the
// KMinDist function for leaf node evaluation.
type EDI struct {
	searchMethod search.SearchFunc
}

func (edi *EDI) Consult(
	board *state.Board, timeLimit time.Duration,
) *state.Move {

	if edi.searchMethod == nil {
		edi.searchMethod = mm.HistoricAlphaBeta(eval.KMinDist)
	}

	return edi.searchMethod(board, timeLimit)
}
