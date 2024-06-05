package main

var basePrice = map[HotelType]int{
	TowerHotel:       200,
	LuxorHotel:       200,
	AmericanHotel:    300,
	WorldwideHotel:   300,
	FestivalHotel:    300,
	ImperialHotel:    400,
	ContinentalHotel: 400,
}

type Hotel struct {
	hotelType HotelType
	board     *Board
}

func NewHotel(hotelType HotelType, board *Board) *Hotel {
	return &Hotel{
		hotelType: hotelType,
		board:     board,
	}
}

func (h *Hotel) size() int {
	size := 0
	for _, cell := range h.board.State {
		if cell == CellState(h.hotelType) {
			size++
		}
	}
	return size
}

func (h *Hotel) price() int {
	size := h.size()
	if size >= 40 {
		return basePrice[h.hotelType] + 700
	}
	if size >= 30 {
		return basePrice[h.hotelType] + 600
	}
	if size >= 20 {
		return basePrice[h.hotelType] + 500
	}
	if size >= 10 {
		return basePrice[h.hotelType] + 400
	}
	if size >= 5 {
		return basePrice[h.hotelType] + 300
	}
	if size == 4 {
		return basePrice[h.hotelType] + 200
	}
	if size == 3 {
		return basePrice[h.hotelType] + 100
	}
	return basePrice[h.hotelType]
}

func (h *Hotel) isActive() bool {
	return h.size() > 0
}

type HotelManager struct {
	board  *Board
	hotels []*Hotel
}

func NewHotelManager(board *Board) *HotelManager {
	hotels := make([]*Hotel, len(AllHotelTypes))
	for i, _ := range AllHotelTypes {
		hotels[i] = NewHotel(AllHotelTypes[i], board)
	}

	return &HotelManager{
		board:  board,
		hotels: hotels,
	}
}

func (m *HotelManager) GetHotel(hotelType HotelType) *Hotel {
	for _, h := range m.hotels {
		if h.hotelType == hotelType {
			return h
		}
	}
	return nil
}

func (m *HotelManager) GetInactiveHotels() []*Hotel {
	inactive := make([]*Hotel, 0)
	for _, hotel := range m.hotels {
		if !hotel.isActive() {
			inactive = append(inactive, hotel)
		}
	}
	return inactive
}

func (m *HotelManager) GetActiveHotels() []*Hotel {
	active := make([]*Hotel, 0)
	for _, hotel := range m.hotels {
		if hotel.isActive() {
			active = append(active, hotel)
		}
	}
	return active
}
