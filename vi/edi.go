package vi

import (
	"time"

	"github.com/edi/eval"
	"github.com/edi/search"
	"github.com/edi/search/mm"
	"github.com/edi/state"
)

// EDI is the flagship VI for this project. At the time of writing, she uses:
// - Minimax search with alpha-beta pruning.
// - The History Heuristic for move ordering.
// - The KMinDist function for leaf node evaluation.
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
