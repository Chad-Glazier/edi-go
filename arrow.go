package edi

import (
	"time"

	"github.com/Chad-Glazier/edi/eval"
	"github.com/Chad-Glazier/edi/search"
	"github.com/Chad-Glazier/edi/search/mm"
	"github.com/Chad-Glazier/edi/state"
)

// Arrow is a program that was created by Martin Müller and Theodore Tegos,
// described in their paper "Experiments in Computer Amazons." They describe
// multiple versions which are selective/non-selective; this implementation
// is non-selective. That is, it's just normal alpha-beta search which uses
// QMinDist to evaluate leaf nodes.
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
