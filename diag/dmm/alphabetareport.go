package dmm

import (
	"time"

	"github.com/Chad-Glazier/edi/state"
)

// A search function that returns a search report instead of a recommended
// move. The "D" stands for "diagnostic."
type DAlphaBetaSearch = func(board *state.Board, time time.Duration) AlphaBetaReport

// Contains the information about how an alpha-beta search went.
type AlphaBetaReport interface {
	// The greatest depth of a completed search during the iterative deepening.
	GreatestDepth() int
	// The time it took to complete the search to a given depth.
	SearchDuration(searchDepth int) time.Duration
	// The number of leaf nodes that were evaluated.
	Leaves(searchDepth int) uint64
	// The number of cutoffs that were produced at each depth.
	Cutoffs(searchDepth, cutoffDepth int) uint64
}

// A struct that stores data for a report. During an alpha-beta search, this
// structure can be maintained and gradually updated.
type report struct {
	// Contains details of each prior iterated search (alpha-beta, in our
	// implementations, uses iterative deepening).
	completedSearches []completeSearch
	current           ongoingSearch
}

type completeSearch struct {
	// The time it took to complete the search.
	duration time.Duration
	// The number of leaf nodes evaluated.
	leaves uint64
	// The number of cutoffs, indexed by the depth of the cutoff.
	cutoffs []uint64
}

type ongoingSearch struct {
	leaves  uint64
	cutoffs []uint64
}

func (r *report) GreatestDepth() int {
	return len(r.completedSearches) - 1
}

func (r *report) SearchDuration(searchDepth int) time.Duration {
	return r.completedSearches[searchDepth].duration
}

func (r *report) Leaves(searchDepth int) uint64 {
	return r.completedSearches[searchDepth].leaves
}

func (r *report) Cutoffs(searchDepth, cutoffDepth int) uint64 {
	return r.completedSearches[searchDepth].cutoffs[cutoffDepth]
}
