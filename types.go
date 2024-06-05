package main

type CellState int

const (
	Empty CellState = iota
	Tile
	American
	Continental
	Festival
	Imperial
	Luxor
	Tower
	Worldwide
)

type HotelType int

const (
	AmericanHotel    HotelType = HotelType(American)
	ContinentalHotel HotelType = HotelType(Continental)
	FestivalHotel    HotelType = HotelType(Festival)
	ImperialHotel    HotelType = HotelType(Imperial)
	LuxorHotel       HotelType = HotelType(Luxor)
	TowerHotel       HotelType = HotelType(Tower)
	WorldwideHotel   HotelType = HotelType(Worldwide)
)

func (h HotelType) String() string {
	switch h {
	case AmericanHotel:
		return "American"
	case ContinentalHotel:
		return "Continental"
	case FestivalHotel:
		return "Festival"
	case ImperialHotel:
		return "Imperial"
	case LuxorHotel:
		return "Luxor"
	case TowerHotel:
		return "Tower"
	case WorldwideHotel:
		return "Worldwide"
	default:
		return "Unknown"
	}
}

var AllHotelTypes = []HotelType{
	AmericanHotel,
	ContinentalHotel,
	FestivalHotel,
	ImperialHotel,
	LuxorHotel,
	TowerHotel,
	WorldwideHotel,
}

type ActionType int

const (
	PlayerActionPlaceTile ActionType = iota + 1
	PlayerActionChooseHotel
	PlayerActionPurchaseShare
	PlayerActionEndTurn
)

func (a ActionType) String() string {
	switch a {
	case PlayerActionPlaceTile:
		return "Place Tile"
	case PlayerActionChooseHotel:
		return "Choose Hotel"
	case PlayerActionPurchaseShare:
		return "Purchase Share"
	case PlayerActionEndTurn:
		return "End Turn"
	default:
		return "Unknown"
	}
}

type PlayerAction struct {
	Type    ActionType
	Payload interface{}
}

type PlayerActionPlaceTilePayload struct {
	Tile int
}

type PlayerActionChooseHotelPayload struct {
	HotelType HotelType
}

type PlayerActionPurchaseSharePayload struct {
	HotelType HotelType
}

func NewPlaceTileAction(tile int) PlayerAction {
	return PlayerAction{
		Type: PlayerActionPlaceTile,
		Payload: PlayerActionPlaceTilePayload{
			Tile: tile,
		},
	}
}

func NewChooseHotelAction(hotelType HotelType) PlayerAction {
	return PlayerAction{
		Type: PlayerActionChooseHotel,
		Payload: PlayerActionChooseHotelPayload{
			HotelType: hotelType,
		},
	}
}

func NewPurchaseShareAction(hotelType HotelType) PlayerAction {
	return PlayerAction{
		Type: PlayerActionPurchaseShare,
		Payload: PlayerActionPurchaseSharePayload{
			HotelType: hotelType,
		},
	}
}

func NewEndTurnAction() PlayerAction {
	return PlayerAction{
		Type: PlayerActionEndTurn,
	}
}

type AvailableAction struct {
	Type    ActionType
	Payload interface{}
}

type AvailableActionChooseTilePayload struct {
	Tiles []int
}

type AvailableActionChooseHotelPayload struct {
	Hotels []*Hotel
}

type AvailableActionChooseSharePayload struct {
	AvailableShares []AvailableShare
}

func NewChooseTileAvailableAction(tiles []int) AvailableAction {
	return AvailableAction{
		Type: PlayerActionPlaceTile,
		Payload: AvailableActionChooseTilePayload{
			Tiles: tiles,
		},
	}
}

func NewEndTurnAvailableAction() AvailableAction {
	return AvailableAction{
		Type: PlayerActionEndTurn,
	}
}

func NewChooseHotelAvailableAction(hotels []*Hotel) AvailableAction {
	return AvailableAction{
		Type: PlayerActionChooseHotel,
		Payload: AvailableActionChooseHotelPayload{
			Hotels: hotels,
		},
	}
}

func NewChooseShareAvailableAction(shares []AvailableShare) AvailableAction {
	return AvailableAction{
		Type: PlayerActionPurchaseShare,
		Payload: AvailableActionChooseSharePayload{
			AvailableShares: shares,
		},
	}
}
