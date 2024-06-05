package main

import (
	"fmt"
	"slices"
)

type GameEngine struct {
	Board            *Board
	HotelManager     *HotelManager
	PlayerManager    *PlayerManager
	TileBag          *TileBag
	AvailableActions []AvailableAction
}

type GameEngineConfig struct {
	seed int64
}

func NewGameEngine(config GameEngineConfig) *GameEngine {
	board := NewBoard()
	tileBag := NewTileBag(config.seed)
	hotelManager := NewHotelManager(board)
	players := []*Player{
		NewPlayer("1"),
		NewPlayer("2"),
		NewPlayer("3"),
	}
	for _, player := range players {
		player.fillTileRack(tileBag)
	}

	playerManager := NewPlayerManager(players)

	return &GameEngine{
		Board:         board,
		TileBag:       tileBag,
		HotelManager:  hotelManager,
		PlayerManager: playerManager,
		AvailableActions: []AvailableAction{
			NewChooseTileAvailableAction(playerManager.CurrentPlayer().Tiles),
		},
	}
}

func (g *GameEngine) handleAction(action PlayerAction) error {
	err := g.validate(action)
	if err != nil {
		return err
	}

	var availableActions []AvailableAction

	if action.Type == PlayerActionPlaceTile {
		payload := action.Payload.(PlayerActionPlaceTilePayload)
		g.PlayerManager.CurrentPlayer().removeTile(payload.Tile)
		if g.Board.HasAdjacentTile(payload.Tile) {
			availableActions = []AvailableAction{
				NewChooseHotelAvailableAction(g.HotelManager.GetInactiveHotels()),
			}
		} else {
			availableActions = []AvailableAction{
				NewEndTurnAvailableAction(),
			}
		}
		g.Board.UpdateCellState(action.Payload.(PlayerActionPlaceTilePayload).Tile, Tile)
	} else if action.Type == PlayerActionChooseHotel {
		payload := action.Payload.(PlayerActionChooseHotelPayload)
		g.Board.UpdateUnmarkedHotel(payload.HotelType)
	} else if action.Type == PlayerActionEndTurn {
		g.PlayerManager.rotate()
		g.PlayerManager.CurrentPlayer().fillTileRack(g.TileBag)
		availableActions = []AvailableAction{
			NewChooseTileAvailableAction(g.PlayerManager.CurrentPlayer().Tiles),
		}
	}

	g.AvailableActions = availableActions

	return nil
}

func (g *GameEngine) validate(action PlayerAction) error {
	var available *AvailableAction
	for _, a := range g.AvailableActions {
		if a.Type == action.Type {
			available = &a
			break
		}
	}

	if available == nil {
		return fmt.Errorf("%s is not an available action", action.Type.String())
	}

	switch action.Type {
	case PlayerActionPlaceTile:
		tileOptions := available.Payload.(AvailableActionChooseTilePayload).Tiles
		tile := action.Payload.(PlayerActionPlaceTilePayload).Tile
		if !slices.Contains(tileOptions, tile) {
			return fmt.Errorf("invalid tile option: %d. valid tiles are %v", tile, tileOptions)
		}
	}
	return nil
}
