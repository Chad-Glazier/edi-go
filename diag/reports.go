package diag

import "time"

type AlphaBetaReport interface {
	// The greatest depth of a completed search during the iterative deepening.
	GreatestDepth() int
	// The time it took to complete the search to a given depth.
	TimeToSearch(searchDepth int) time.Duration
	// The number of leaf nodes that were evaluated.
	Leaves(searchDepth int) uint64
	// The number of cutoffs that were produced at each depth.
	Cutoffs(searchDepth, cutoffDepth int) uint64
}
