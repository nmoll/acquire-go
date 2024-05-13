package main

import "testing"

func TestGameEngine(t *testing.T) {
	t.Run("Should return initial game state", func(t *testing.T) {
		engine := NewGameEngine()

		cellState := make([]CellState, totalCells)
		for i := range cellState {
			cellState[i] = Empty
		}

		want := &GameState{
			CellState: cellState,
		}

		got := engine.computeState()

		AssertGameState(t, got, want)
	})

	t.Run("Should play actions", func(t *testing.T) {
		engine := NewGameEngine()

		engine.handleAction(NewPlaceTileAction(5))
		engine.handleAction(NewPlaceTileAction(50))
		engine.handleAction(NewPlaceTileAction(6))

		cellState := make([]CellState, totalCells)
		for i := range cellState {
			cellState[i] = Empty
		}
		cellState[5] = Tile
		cellState[6] = Tile
		cellState[50] = Tile

		want := &GameState{
			CellState: cellState,
		}

		got := engine.computeState()

		AssertGameState(t, got, want)
	})
}
