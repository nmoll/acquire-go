package main

import "testing"

func AssertGameState(t *testing.T, got *GameState, want *GameState) {
	for i, v := range got.CellState {
		if v != want.CellState[i] {
			t.Errorf("Cell states do not match at index %d. got %d want %d", i, v, want.CellState[i])
		}
	}
}
