package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

type ExpectedGameState struct {
	board           string
	availableAction string
	currentPlayer   string
}

type GameEngineTestHelper struct {
	t                 *testing.T
	engine            *GameEngine
	failMessagePrefix string
}

func NewGameEngineTestHelper(t *testing.T, engine *GameEngine) *GameEngineTestHelper {
	return &GameEngineTestHelper{
		t:      t,
		engine: engine,
	}
}

func (h *GameEngineTestHelper) placeTile(tile string) *GameEngineTestHelper {
	return h.playAction(NewPlaceTileAction(parseTile(tile)))
}

func (h *GameEngineTestHelper) playAction(action PlayerAction) *GameEngineTestHelper {
	err := h.engine.handleAction(action)
	if err != nil {
		h.t.Error(err)
	}
	return h
}

func (h *GameEngineTestHelper) assertErrorWithAction(action PlayerAction, errMessage string) {
	err := h.engine.handleAction(action)
	if err == nil {
		h.t.Errorf("Expected action [%s] to return error: %s", action.Type.String(), errMessage)
	}

	if err.Error() != errMessage {
		h.t.Errorf("Unexpected error from action. got [%s] want [%s]", err.Error(), errMessage)
	}
}

func (h *GameEngineTestHelper) assertState(messagePrefix string, expected ExpectedGameState) *GameEngineTestHelper {
	h.failMessagePrefix = messagePrefix
	h.assertBoardState(expected.board)
	h.assertAvailableAction(expected.availableAction)
	h.assertCurrentPlayer(expected.currentPlayer)
	return h
}

func (h *GameEngineTestHelper) assertBoardState(want string) {
	got := boardStateToString(h.engine.Board.State)
	if stripWhitespace(got) != stripWhitespace(want) {
		h.fail(fmt.Sprintf("Board states do not match. \nGot: \n%s \nwant\n%s", got, want))
	}
}

func (h *GameEngineTestHelper) assertAvailableAction(want string) {
	got := ""

	hotelChar := map[HotelType]string{
		AmericanHotel:    "A",
		ContinentalHotel: "C",
		FestivalHotel:    "F",
		ImperialHotel:    "I",
		LuxorHotel:       "L",
		TowerHotel:       "T",
		WorldwideHotel:   "W",
	}

	availableAction := h.engine.AvailableAction

	switch availableAction.Type {
	case PlayerActionPlaceTile:
		got += "Place Tile\n"
		tiles := availableAction.Payload.(AvailableActionChooseTilePayload).Tiles
		for _, tile := range tiles {
			got += getTileDisplay(tile) + " "
		}
	case PlayerActionChooseHotel:
		got += "Choose Hotel\n"

		for _, hotel := range availableAction.Payload.(AvailableActionChooseHotelPayload).Hotels {
			got += hotelChar[hotel.hotelType] + " "
		}
	case PlayerActionPurchaseShare:
		got += "Purchase Share\n"

		for _, available := range availableAction.Payload.(AvailableActionChooseSharePayload).AvailableShares {
			got += hotelChar[available.HotelType]
		}

	case PlayerActionEndTurn:
		got += "End Turn\n"
	}

	if stripWhitespace(got) != stripWhitespace(want) {
		h.fail(fmt.Sprintf("Available actions do not match. \nGot: \n%s \nwant\n%s", got, want))
	}
}

func (h *GameEngineTestHelper) assertCurrentPlayer(want string) *GameEngineTestHelper {
	player := h.engine.PlayerManager.CurrentPlayer()
	got := player.Id + " | $" + strconv.Itoa(player.Cash)

	playerShares := h.engine.StockBroker.GetPlayerShares(player)
	got += " A:" + strconv.Itoa(playerShares.Count(AmericanHotel))
	got += " C:" + strconv.Itoa(playerShares.Count(ContinentalHotel))
	got += " F:" + strconv.Itoa(playerShares.Count(FestivalHotel))
	got += " I:" + strconv.Itoa(playerShares.Count(ImperialHotel))
	got += " L:" + strconv.Itoa(playerShares.Count(LuxorHotel))
	got += " T:" + strconv.Itoa(playerShares.Count(TowerHotel))
	got += " W:" + strconv.Itoa(playerShares.Count(WorldwideHotel))

	if stripWhitespace(want) != stripWhitespace(got) {
		h.fail(fmt.Sprintf("Current player does not match. \nGot: \n%s \nwant\n%s", got, want))
	}
	return h
}

func (h *GameEngineTestHelper) fail(message string) {
	h.t.Errorf("%s: %s", h.failMessagePrefix, message)
}

func boardStateToString(boardState []CellState) string {
	diagram := ""
	for idx, state := range boardState {
		cellStateMap := map[CellState]string{
			Empty:       "-",
			Tile:        "0",
			American:    "A",
			Continental: "C",
			Festival:    "F",
			Imperial:    "I",
			Luxor:       "L",
			Tower:       "T",
			Worldwide:   "W",
		}

		var separator string
		if (idx+1)%12 == 0 {
			separator = "\n"
		} else {
			separator = " "
		}
		diagram += cellStateMap[state] + separator
	}
	return diagram
}

func stripWhitespace(str string) string {
	res := strings.ReplaceAll(str, "\n", "")
	res = strings.ReplaceAll(res, "\t", "")
	res = strings.ReplaceAll(res, " ", "")
	return res
}

func getTileDisplay(tile int) string {
	letters := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}

	col := (tile % 12) + 1
	row := tile / 12

	return strconv.Itoa(col) + letters[row]
}

func parseTile(tile string) int {
	rowByLetter := map[string]int{
		"A": 0,
		"B": 1,
		"C": 2,
		"D": 3,
		"E": 4,
		"F": 5,
		"G": 6,
		"H": 7,
		"I": 8,
	}

	var row, col int
	if len(tile) == 3 {
		col, _ = strconv.Atoi(string(tile[0]) + string(tile[1]))
		row = rowByLetter[string(tile[2])]
	} else {
		col, _ = strconv.Atoi(string(tile[0]))
		row = rowByLetter[string(tile[1])]
	}

	return row*12 + col - 1
}
