package main

import (
	"fmt"
	"slices"
)

type GameEngine struct {
	Board           *Board
	HotelManager    *HotelManager
	StockBroker     *StockBroker
	PlayerManager   *PlayerManager
	TileBag         *TileBag
	AvailableAction AvailableAction
}

type GameEngineConfig struct {
	seed int64
}

func NewGameEngine(config GameEngineConfig) *GameEngine {
	board := NewBoard()
	tileBag := NewTileBag(config.seed)
	hotelManager := NewHotelManager(board)
	stockBroker := NewStockBroker(hotelManager)

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
		Board:           board,
		TileBag:         tileBag,
		HotelManager:    hotelManager,
		StockBroker:     stockBroker,
		PlayerManager:   playerManager,
		AvailableAction: NewChooseTileAvailableAction(playerManager.CurrentPlayer().Tiles),
	}
}

func (g *GameEngine) handleAction(action PlayerAction) error {
	err := g.validate(action)
	if err != nil {
		return err
	}

	var availableAction AvailableAction

	if action.Type == PlayerActionPlaceTile {
		payload := action.Payload.(PlayerActionPlaceTilePayload)
		g.PlayerManager.CurrentPlayer().removeTile(payload.Tile)
		if g.Board.HasAdjacentTile(payload.Tile) {
			availableAction = NewChooseHotelAvailableAction(g.HotelManager.GetInactiveHotels())
		} else {
			availableAction = NewEndTurnAvailableAction()
		}
		g.Board.UpdateCellState(action.Payload.(PlayerActionPlaceTilePayload).Tile, Tile)
	} else if action.Type == PlayerActionChooseHotel {
		payload := action.Payload.(PlayerActionChooseHotelPayload)
		g.Board.UpdateUnmarkedHotel(payload.HotelType)
		g.StockBroker.AwardShare(g.PlayerManager.CurrentPlayer(), payload.HotelType)
		availableAction = NewChooseShareAvailableAction(g.StockBroker.GetAvailableToPurchase())
	} else if action.Type == PlayerActionPurchaseShare {
		payload := action.Payload.(PlayerActionPurchaseSharePayload)
		g.StockBroker.Purchase(g.PlayerManager.CurrentPlayer(), payload.HotelType)
		availableAction = NewChooseShareAvailableAction(g.StockBroker.GetAvailableToPurchase())
	} else if action.Type == PlayerActionEndTurn {
		g.PlayerManager.rotate()
		g.PlayerManager.CurrentPlayer().fillTileRack(g.TileBag)
		availableAction = NewChooseTileAvailableAction(g.PlayerManager.CurrentPlayer().Tiles)
	}

	g.AvailableAction = availableAction

	return nil
}

func (g *GameEngine) validate(action PlayerAction) error {
	if g.AvailableAction.Type != action.Type {
		return fmt.Errorf("%s is not an available action", action.Type.String())
	}

	switch action.Type {
	case PlayerActionPlaceTile:
		tileOptions := g.AvailableAction.Payload.(AvailableActionChooseTilePayload).Tiles
		tile := action.Payload.(PlayerActionPlaceTilePayload).Tile
		if !slices.Contains(tileOptions, tile) {
			return fmt.Errorf("invalid tile option: %d. valid tiles are %v", tile, tileOptions)
		}
	case PlayerActionPurchaseShare:
		available := g.AvailableAction.Payload.(AvailableActionChooseSharePayload).AvailableShares
		hotelType := action.Payload.(PlayerActionPurchaseSharePayload).HotelType
		found := false
		for _, a := range available {
			if a.HotelType == hotelType {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("%s is not available to purchase", hotelType.String())
		}
	}
	return nil
}
