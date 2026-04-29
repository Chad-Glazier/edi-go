package search

import "github.com/edi/state"

// Looks at a given boardstate and yields a recommended move within the given
// time limit (which is in milliseconds)
type SearchFunc func(board *state.Board, milliseconds int) *state.Move
