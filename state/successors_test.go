package state

import (
	"testing"
)

func BenchmarkSuccessorsInitial(b *testing.B) {
	board := InitialState()
	for b.Loop() {
		board.Successors()
	}
}

func BenchmarkSuccessorsTurn15(b *testing.B) {
	board := RandomBoard(15)
	for b.Loop() {
		board.Successors()
	}
}

func BenchmarkSuccessorsTurn30(b *testing.B) {
	board := RandomBoard(30)
	for b.Loop() {
		board.Successors()
	}
}

func BenchmarkSuccessorsTurn45(b *testing.B) {
	board := RandomBoard(45)
	for b.Loop() {
		board.Successors()
	}
}
