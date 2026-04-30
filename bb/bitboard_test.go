package bb

import (
	"math/rand"
	"testing"
)

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
	iter := bb
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

func BenchmarkNext(b *testing.B) {
	bb, _, _ := randomBoard(0.20)
	for b.Loop() {
		iter := bb
		for pos := iter.Next(); pos != NULL_POS; pos = iter.Next() {
			// nothing.
		}
	}
}

func TestMsbLsb(t *testing.T) {

	empty := BitBoard{}
	if empty.Msb() != NULL_POS {
		t.Errorf(
			"Expected MSB of empty board to be NULL_POS, got %d", empty.Msb(),
		)
	}
	if empty.Lsb() != NULL_POS {
		t.Errorf(
			"Expected LSB of empty board to be NULL_POS, got %d", empty.Lsb(),
		)
	}

	bb, flagged, _ := randomBoard(0.40)

	minPos := Position(99)
	maxPos := Position(0)

	for pos := range Position(100) {
		if flagged[pos] {
			if pos < minPos {
				minPos = pos
			}
			if pos > maxPos {
				maxPos = pos
			}
		}
	}

	if bb.Lsb() != minPos {
		t.Errorf("Expected LSB %d, got %d", minPos, bb.Lsb())
	}
	if bb.Msb() != maxPos {
		t.Errorf("Expected MSB %d, got %d", maxPos, bb.Msb())
	}

	for pos := range Position(100) {
		single := BitBoard{}
		single.Flag(pos)

		if single.Lsb() != pos {
			t.Errorf("Single-bit LSB failed at %d, got %d", pos, single.Lsb())
		}
		if single.Msb() != pos {
			t.Errorf("Single-bit MSB failed at %d, got %d", pos, single.Msb())
		}
	}
}

func BenchmarkBitwiseOperations(b *testing.B) {
	x, y := BitBoard{}, BitBoard{}
	for b.Loop() {
		x.Or(y)
		x.Xor(y)
		x.And(y)
		x.AndNot(y)
		x.Not()
	}
}
