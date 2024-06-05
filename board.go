package main

const (
	totalCells = 108
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

func (b *Board) UpdateCellState(cell int, state CellState) {
	b.State[cell] = state
}

func (b *Board) HasAdjacentTile(cell int) bool {
	for _, state := range b.GetAdjacentCells(cell) {
		if state == Tile {
			return true
		}
	}

	return false
}

func (b *Board) GetAdjacentCells(cell int) []CellState {
	result := make([]CellState, 0)
	if cell > 0 { // to the left
		result = append(result, b.State[cell-1])
	}
	if cell < totalCells-1 { // to the right
		result = append(result, b.State[cell+1])
	}
	return result
}

// UpdateUnmarkedHotel finds any tiles that are touching and
// converts them to the given hotel
func (b *Board) UpdateUnmarkedHotel(hotelType HotelType) {
	touchingTileIndexes := make([]int, 0)
	for i, state := range b.State {
		if state == Tile && b.HasAdjacentTile(i) {
			touchingTileIndexes = append(touchingTileIndexes, i)
		}
	}

	for _, tileIdx := range touchingTileIndexes {
		b.UpdateCellState(tileIdx, CellState(hotelType))
	}
}
