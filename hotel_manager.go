package main

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

func (m *HotelManager) GetInactiveHotels() []*Hotel {
	inactive := make([]*Hotel, 0)
	for _, hotel := range m.hotels {
		if !hotel.isActive() {
			inactive = append(inactive, hotel)
		}
	}
	return inactive
}
