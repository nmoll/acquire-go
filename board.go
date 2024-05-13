package main

const (
	totalCells = 108
)

type CellState int

const (
	Empty CellState = iota + 1
	Tile
	American
	Luxor
)

type Board struct {
	State []CellState
}

func NewBoard() *Board {
	state := make([]CellState, totalCells)
	for i := range state {
		state[i] = Empty
	}

	return &Board{
		State: state,
	}
}

func (b *Board) handleEvent(event *GameEvent) {
	if event.Type == TilePlaced {
		payload := event.Payload.(GameEventTilePlacedPayload)
		b.State[payload.Tile] = Tile
	}
}
