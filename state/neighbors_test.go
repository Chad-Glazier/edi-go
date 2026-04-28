package state

import (
	"testing"
)

//
// Ground-truth functions. These implementations are certainly correct, just
// terribly inefficient. We can use them to get the expected results.
//

func inBounds(row, col int) bool {
	return row >= 0 && row < 10 && col >= 0 && col < 10
}

func expectedKNeighbors(occ BitBoard, pos Position) BitBoard {

	result := BitBoard{}
	i, j := Coords(pos)

	for di := -1; di <= 1; di++ {
		for dj := -1; dj <= 1; dj++ {
			if di == 0 && dj == 0 {
				continue
			}
			nx, ny := i+di, j+dj
			if !inBounds(nx, ny) {
				continue
			}
			p := Pos(nx, ny)
			if !occ.Flagged(p) {
				result.Flag(p)
			}
		}
	}

	return result
}

func expectedQNeighbors(occ BitBoard, pos Position) BitBoard {
	result := BitBoard{}
	i, j := Coords(pos)

	dirs := [8][2]int{
		{1, 0}, {-1, 0}, {0, 1}, {0, -1},
		{1, 1}, {1, -1}, {-1, 1}, {-1, -1},
	}

	for _, d := range dirs {
		ni, nj := i+d[0], j+d[1]
		for inBounds(ni, nj) {
			p := Pos(ni, nj)
			if occ.Flagged(p) {
				break
			}
			result.Flag(p)
			ni += d[0]
			nj += d[1]
		}
	}

	return result
}

func expectedFrontier(
	occ BitBoard,
	territory BitBoard,
	neighborFn func(BitBoard, Position) BitBoard,
) BitBoard {

	result := BitBoard{}

	iter := territory.Copy()
	for pos := iter.Next(); pos != NULL_POS; pos = iter.Next() {
		neighbors := neighborFn(occ, pos)
		for n := neighbors.Next(); n != NULL_POS; n = neighbors.Next() {
			if !territory.Flagged(n) {
				result.Flag(n)
			}
		}
	}

	return result
}

//
// Tests
//

func TestKNeighbors(t *testing.T) {
	for range 50 {
		occ, _, _ := randomBoard(0.2)

		for pos := range Position(100) {
			got := KNeighbors(&occ, pos)
			expected := expectedKNeighbors(occ, pos)

			if got.Count() != expected.Count() {
				t.Fatalf("KNeighbors size mismatch at %d", pos)
			}

			for p := expected.Next(); p != NULL_POS; p = expected.Next() {
				if !got.Flagged(p) {
					t.Errorf("KNeighbors missing %d from %d", p, pos)
				}
			}
		}
	}
}

func TestQNeighbors(t *testing.T) {
	for range 50 {
		occ, _, _ := randomBoard(0.2)

		for pos := range Position(100) {
			got := QNeighbors(&occ, pos)
			expected := expectedQNeighbors(occ, pos)

			if got.Count() != expected.Count() {
				t.Fatalf("QNeighbors size mismatch at %d", pos)
			}

			iter := expected.Copy()
			for p := iter.Next(); p != NULL_POS; p = iter.Next() {
				if !got.Flagged(p) {
					t.Errorf("QNeighbors missing %d from %d", p, pos)
				}
			}
		}
	}
}

func TestKFrontier(t *testing.T) {
	for range 50 {
		occ, _, _ := randomBoard(0.2)
		territory, _, _ := randomBoard(0.2)

		got := KFrontier(&occ, &territory)
		expected := expectedFrontier(occ, territory, expectedKNeighbors)

		if got.Count() != expected.Count() {
			t.Fatalf("KFrontier size mismatch")
		}

		iter := expected.Copy()
		for p := iter.Next(); p != NULL_POS; p = iter.Next() {
			if !got.Flagged(p) {
				t.Errorf("KFrontier missing %d", p)
			}
		}
	}
}

func TestQFrontier(t *testing.T) {
	for range 50 {
		occ, _, _ := randomBoard(0.2)
		territory, _, _ := randomBoard(0.2)

		got := QFrontier(&occ, &territory)
		expected := expectedFrontier(occ, territory, expectedQNeighbors)

		if got.Count() != expected.Count() {
			t.Fatalf("QFrontier size mismatch")
		}

		iter := expected.Copy()
		for p := iter.Next(); p != NULL_POS; p = iter.Next() {
			if !got.Flagged(p) {
				t.Errorf("QFrontier missing %d", p)
			}
		}
	}
}
