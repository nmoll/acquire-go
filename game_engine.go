package main

type GameState struct {
	CellState []CellState
}

type GameEngine struct {
	Actions []PlayerAction
	board   *Board
}

func NewGameEngine() *GameEngine {
	return &GameEngine{
		board: NewBoard(),
	}
}

func (g *GameEngine) handleAction(action *PlayerAction) {
	var event *GameEvent
	if action.Type == PlayerActionPlaceTile {
		event = &GameEvent{
			Type: TilePlaced,
			Payload: GameEventTilePlacedPayload{
				Tile: action.Payload.(PlayerActionPlaceTilePayload).Tile,
			},
		}
	}
	g.board.handleEvent(event)
}

func (g *GameEngine) computeState() *GameState {
	return &GameState{
		CellState: g.board.State,
	}
}
