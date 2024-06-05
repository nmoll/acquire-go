package main

import "testing"

func TestGameEngine_handleAction(t *testing.T) {
	t.Run("should play actions", func(t *testing.T) {
		engine := NewGameEngine(GameEngineConfig{seed: 1})
		helper := NewGameEngineTestHelper(t, engine)

		helper.
			assertState("Action 1", ExpectedGameState{
				currentPlayer:    "1",
				availableActions: "Place Tile 7H 1A 5C 4I 7E 7C",
				board: `
				- - - - - - - - - - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -	`,
			}).
			placeTile("5C").
			assertState("Action 2", ExpectedGameState{
				currentPlayer:    "1",
				availableActions: "End Turn",
				board: `
				- - - - - - - - - - - -
				- - - - - - - - - - - -
				- - - - 0 - - - - - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -	`,
			}).
			playAction(NewEndTurnAction()).
			assertState("Action 3", ExpectedGameState{
				currentPlayer:    "2",
				availableActions: "Place Tile 7I 3B 3G 9F 10H 3H",
				board: `
				- - - - - - - - - - - -
				- - - - - - - - - - - -
				- - - - 0 - - - - - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -	`,
			}).
			placeTile("9F").
			assertState("Action 4", ExpectedGameState{
				currentPlayer:    "2",
				availableActions: "End Turn",
				board: `
				- - - - - - - - - - - -
				- - - - - - - - - - - -
				- - - - 0 - - - - - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -
				- - - - - - - - 0 - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -	`,
			}).
			playAction(NewEndTurnAction()).
			assertState("Action 5", ExpectedGameState{
				currentPlayer:    "3",
				availableActions: "Place Tile 12B 5G 9H 2C 6I 11D",
				board: `
				- - - - - - - - - - - -
				- - - - - - - - - - - -
				- - - - 0 - - - - - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -
				- - - - - - - - 0 - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -`,
			}).
			placeTile("12B").
			assertState("Action 6", ExpectedGameState{
				currentPlayer:    "3",
				availableActions: "End Turn",
				board: `
				- - - - - - - - - - - -
				- - - - - - - - - - - 0
				- - - - 0 - - - - - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -
				- - - - - - - - 0 - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -`,
			}).
			playAction(NewEndTurnAction()).
			assertState("Action 7", ExpectedGameState{
				currentPlayer:    "1",
				availableActions: "Place Tile 7H 1A 4I 7E 7C 2H",
				board: `
				- - - - - - - - - - - -
				- - - - - - - - - - - 0
				- - - - 0 - - - - - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -
				- - - - - - - - 0 - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -`,
			}).
			placeTile("7E").
			assertState("Action 8", ExpectedGameState{
				currentPlayer:    "1",
				availableActions: "End Turn",
				board: `
				- - - - - - - - - - - -
				- - - - - - - - - - - 0
				- - - - 0 - - - - - - -
				- - - - - - - - - - - -
				- - - - - - 0 - - - - -
				- - - - - - - - 0 - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -`,
			}).
			playAction(NewEndTurnAction()).
			assertState("Action 9", ExpectedGameState{
				currentPlayer:    "2",
				availableActions: "Place Tile 7I 3B 3G 10H 3H 2A",
				board: `
				- - - - - - - - - - - -
				- - - - - - - - - - - 0
				- - - - 0 - - - - - - -
				- - - - - - - - - - - -
				- - - - - - 0 - - - - -
				- - - - - - - - 0 - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -`,
			}).
			placeTile("7I").
			assertState("Action 10", ExpectedGameState{
				currentPlayer:    "2",
				availableActions: "End Turn",
				board: `
				- - - - - - - - - - - -
				- - - - - - - - - - - 0
				- - - - 0 - - - - - - -
				- - - - - - - - - - - -
				- - - - - - 0 - - - - -
				- - - - - - - - 0 - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -
				- - - - - - 0 - - - - -`,
			}).
			playAction(NewEndTurnAction()).
			assertState("Action 11", ExpectedGameState{
				currentPlayer:    "3",
				availableActions: "Place Tile 5G 9H 2C 6I 11D 2B",
				board: `
				- - - - - - - - - - - -
				- - - - - - - - - - - 0
				- - - - 0 - - - - - - -
				- - - - - - - - - - - -
				- - - - - - 0 - - - - -
				- - - - - - - - 0 - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -
				- - - - - - 0 - - - - -`,
			}).
			placeTile("2B").
			assertState("Action 12", ExpectedGameState{
				currentPlayer:    "3",
				availableActions: "End Turn",
				board: `
				- - - - - - - - - - - -
				- 0 - - - - - - - - - 0
				- - - - 0 - - - - - - -
				- - - - - - - - - - - -
				- - - - - - 0 - - - - -
				- - - - - - - - 0 - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -
				- - - - - - 0 - - - - -`,
			}).
			playAction(NewEndTurnAction()).
			assertState("Action 13", ExpectedGameState{
				currentPlayer:    "1",
				availableActions: "Place Tile 7H 1A 4I 7C 2H 8G",
				board: `
				- - - - - - - - - - - -
				- 0 - - - - - - - - - 0
				- - - - 0 - - - - - - -
				- - - - - - - - - - - -
				- - - - - - 0 - - - - -
				- - - - - - - - 0 - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -
				- - - - - - 0 - - - - -`,
			}).
			placeTile("8G").
			assertState("Action 13", ExpectedGameState{
				currentPlayer:    "1",
				availableActions: "End Turn",
				board: `
				- - - - - - - - - - - -
				- 0 - - - - - - - - - 0
				- - - - 0 - - - - - - -
				- - - - - - - - - - - -
				- - - - - - 0 - - - - -
				- - - - - - - - 0 - - -
				- - - - - - - 0 - - - -
				- - - - - - - - - - - -
				- - - - - - 0 - - - - -`,
			}).
			playAction(NewEndTurnAction()).
			assertState("Action 14", ExpectedGameState{
				currentPlayer:    "2",
				availableActions: "Place Tile 3B 3G 10H 3H 2A 8A",
				board: `
				- - - - - - - - - - - -
				- 0 - - - - - - - - - 0
				- - - - 0 - - - - - - -
				- - - - - - - - - - - -
				- - - - - - 0 - - - - -
				- - - - - - - - 0 - - -
				- - - - - - - 0 - - - -
				- - - - - - - - - - - -
				- - - - - - 0 - - - - -`,
			}).
			placeTile("3B").
			assertState("Action 15", ExpectedGameState{
				currentPlayer:    "2",
				availableActions: "Choose Hotel A C F I L T W",
				board: `
				- - - - - - - - - - - -
				- 0 0 - - - - - - - - 0
				- - - - 0 - - - - - - -
				- - - - - - - - - - - -
				- - - - - - 0 - - - - -
				- - - - - - - - 0 - - -
				- - - - - - - 0 - - - -
				- - - - - - - - - - - -
				- - - - - - 0 - - - - -`,
			}).
			playAction(NewChooseHotelAction(FestivalHotel))
		//assertState("Action 16", ExpectedGameState{
		//	currentPlayer:    "2",
		//	availableActions: "Purchase Share A C F I L T W",
		//	board: `
		//	- - - - - - - - - - - -
		//	- F F - - - - - - - - 0
		//	- - - - 0 - - - - - - -
		//	- - - - - - - - - - - -
		//	- - - - - - 0 - - - - -
		//	- - - - - - - - 0 - - -
		//	- - - - - - - 0 - - - -
		//	- - - - - - - - - - - -
		//	- - - - - - 0 - - - - -`,
		//})

	})

	t.Run("should return error if tile place action is not available", func(t *testing.T) {
		engine := NewGameEngine(GameEngineConfig{seed: 1})
		helper := NewGameEngineTestHelper(t, engine)

		helper.
			placeTile("5C").
			assertState("place tile action should not be available", ExpectedGameState{
				currentPlayer:    "1",
				availableActions: "End Turn",
				board: `
				- - - - - - - - - - - -
				- - - - - - - - - - - -
				- - - - 0 - - - - - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -	`,
			})

		helper.assertErrorWithAction(NewPlaceTileAction(1), "Place Tile is not an available action")
	})

	t.Run("should return error if player does not have tile", func(t *testing.T) {
		engine := NewGameEngine(GameEngineConfig{seed: 1})
		helper := NewGameEngineTestHelper(t, engine)

		helper.
			assertState("tile should not be available to play", ExpectedGameState{
				currentPlayer:    "1",
				availableActions: "Place Tile 7H 1A 5C 4I 7E 7C",
				board: `
				- - - - - - - - - - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -
				- - - - - - - - - - - -	`,
			})

		helper.assertErrorWithAction(NewPlaceTileAction(parseTile("3B")), "invalid tile option: 14. valid tiles are [90 0 28 99 54 30]")
	})
}
