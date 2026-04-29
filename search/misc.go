package search

import (
	"time"

	"github.com/edi/state"
)

// Looks at a given boardstate and yields a recommended move within the given
// time limit.
type SearchFunc func(board *state.Board, time time.Duration) *state.Move
