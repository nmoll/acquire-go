package main

const maxShares = 25

type StockBroker struct {
	hotelManager *HotelManager
	shares       []*HotelShare
}

type HotelShare struct {
	HotelType HotelType
	Owner     *Player
}

type PlayerShares struct {
	Shares []*HotelShare
}

type AvailableShare struct {
	HotelType HotelType
	Price     int
	Count     int
}

func NewStockBroker(hotelManager *HotelManager) *StockBroker {
	shares := make([]*HotelShare, 0)

	for _, hotelType := range AllHotelTypes {
		for i := 0; i < maxShares; i++ {
			shares = append(shares, &HotelShare{
				HotelType: hotelType,
			})
		}
	}

	return &StockBroker{
		hotelManager: hotelManager,
		shares:       shares,
	}
}

func (sb *StockBroker) AwardShare(player *Player, hotelType HotelType) {
	for _, share := range sb.shares {
		if share.HotelType == hotelType && share.Owner == nil {
			share.Owner = player
			break
		}
	}
}

func (sb *StockBroker) Purchase(player *Player, hotelType HotelType) {
	player.Cash -= sb.hotelManager.GetHotel(hotelType).price()
	sb.AwardShare(player, hotelType)
}

func (sb *StockBroker) GetAvailableToPurchase() []AvailableShare {
	available := make([]AvailableShare, 0)

	for _, hotel := range sb.hotelManager.GetActiveHotels() {
		purchasedCount := 0
		for _, share := range sb.shares {
			if share.HotelType == hotel.hotelType && share.Owner != nil {
				purchasedCount++
			}
		}
		available = append(available, AvailableShare{
			HotelType: hotel.hotelType,
			Count:     maxShares - purchasedCount,
			Price:     0,
		})
	}

	return available
}

func (sb *StockBroker) GetPlayerShares(player *Player) *PlayerShares {
	shares := make([]*HotelShare, 0)
	for _, share := range sb.shares {
		if share.Owner == player {
			shares = append(shares, share)
		}
	}

	return &PlayerShares{
		Shares: shares,
	}
}

func (ps *PlayerShares) Count(hotelType HotelType) int {
	total := 0
	for _, share := range ps.Shares {
		if share.HotelType == hotelType {
			total++
		}
	}
	return total
}
