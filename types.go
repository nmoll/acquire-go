package main

type ActionType int

const (
	PlayerActionPlaceTile ActionType = iota + 1
	PlayerActionChooseCompany
	PlayerActionPurchaseShare
)

type PlayerAction struct {
	Type    ActionType
	Payload interface{}
}

type PlayerActionPlaceTilePayload struct {
	Tile int
}

func NewPlaceTileAction(tile int) *PlayerAction {
	return &PlayerAction{
		Type: PlayerActionPlaceTile,
		Payload: PlayerActionPlaceTilePayload{
			Tile: tile,
		},
	}
}

type GameEventType int

const (
	TilePlaced GameEventType = iota + 1
	CompanyStarted
	SharesPurchased
	Merge
)

type GameEvent struct {
	Type    GameEventType
	Payload interface{}
}

type GameEventTilePlacedPayload struct {
	Tile int
}
