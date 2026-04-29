package eval

import (
	"github.com/edi/state"
)

// A function that evaluates a board state and returns a score that reflects
// the quality of the position. Positive values indicate favorability for
// White, negative values are favorable for Black, and 0 marks a neutral
// position.
type EvaluationFunc func(board *state.Board) float64
