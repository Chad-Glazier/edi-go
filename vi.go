package edi

import (
	"time"

	"github.com/Chad-Glazier/edi/state"
)

// The term VI, short for "virtual intelligence," is a program that can
// recommend a move. Unlike a search function, a VI may maintain complex
// internal state between moves and decide to change its search method.
//
// The term VI is borrowed from a videogame:
// https://masseffect.fandom.com/wiki/Virtual_Intelligence. Traditionally, we
// would call such a program "AI," but that term has been diluted by the
// the most insufferable kind of nerds.
type VI interface {
	Consult(board *state.Board, timeLimit time.Duration) *state.Move
}
