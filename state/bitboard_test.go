package state

import (
	"math/rand"
	"testing"
)

// Used to "consume" values to avoid dead code or "unused variable" errors.
var blackHole any

func randomBoard(density float64) (
	bb BitBoard, flagged map[Position]bool, flagCount int,
) {
	bb = BitBoard{}
	flagged = make(map[Position]bool, 100)
	flagCount = 0

	for pos := range Position(100) {
		if rand.Float64() < density {
			flagged[pos] = true
			bb.Flag(pos)
			flagCount++
		} else {
			flagged[pos] = false
		}
	}

	return
}

func TestFlagging(t *testing.T) {

	bb, flagged, flagCount := randomBoard(0.40)

	for pos := range Position(100) {
		if flagged[pos] && !bb.Flagged(pos) {
			t.Errorf("Expected flag at %d", pos)
		}
		if !flagged[pos] && bb.Flagged(pos) {
			t.Errorf("Unexpected flag at %d", pos)
		}
	}

	iteratedPositions := 0
	for pos := range bb.Positions() {
		iteratedPositions++
		if !flagged[pos] {
			t.Errorf("Expected iterated position %d to be flagged.", pos)
		}
	}

	if iteratedPositions != flagCount {
		t.Errorf(
			"Iterated over %d flags, but expected %d",
			iteratedPositions, flagCount,
		)
	}

	iteratedPositions = 0
	iter := bb.Copy()
	for pos := iter.Next(); pos != NULL_POS; pos = iter.Next() {
		iteratedPositions++
		if !flagged[pos] {
			t.Errorf("Expected iterated position %d to be flagged.", pos)
		}
	}

	if iteratedPositions != flagCount {
		t.Errorf(
			"Iterated over %d flags, but expected %d",
			iteratedPositions, flagCount,
		)
	}

	for pos := range Position(100) {

		bb.Unflag(pos)
		if flagged[pos] {
			flagCount--
		}

		if bb.Count() != flagCount {
			t.Errorf(
				"Expected flag count %d to be %d",
				bb.Count(), flagCount,
			)
		}
	}

	if !bb.Empty() {
		t.Errorf("Expected board to be empty")
	}
}

func BenchmarkPositions(b *testing.B) {
	
	bb, _, _ := randomBoard(0.20)

	for b.Loop() {
		for pos := range bb.Positions() {
			blackHole = pos
		}
	}
}

func BenchmarkNext(b *testing.B) {
	
	bb, _, _ := randomBoard(0.20)

	for b.Loop() {
		iter := bb.Copy()
		for pos := iter.Next(); pos != NULL_POS; pos = iter.Next() {
			blackHole = pos
		}
	}
}
