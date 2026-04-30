package vi

import (
	"time"

	"github.com/edi/state"
)

type VI interface {
	Consult(board *state.Board, timeLimit time.Duration) *state.Move
}
