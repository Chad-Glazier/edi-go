package diag

import (
	"time"

	"github.com/Chad-Glazier/edi/diag/dmm"
	"github.com/Chad-Glazier/edi/state"
)

// A union of all search report types.
type Report interface {
	dmm.AlphaBetaReport
}

// Represents a diagnostic search function. I.e., a search function that
// returns a report including relevant metrics instead of just a move.
type SearchFunc[R Report] func(board *state.Board, time time.Duration) R
