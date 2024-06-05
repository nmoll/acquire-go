package main

import (
	"slices"
)

const (
	tileRackSize = 6
	startingCash = 6000
)

type Player struct {
	Id    string
	Tiles []int
	Cash  int
}

func NewPlayer(id string) *Player {
	return &Player{
		Id:    id,
		Tiles: make([]int, 0),
		Cash:  startingCash,
	}
}

func (p *Player) fillTileRack(tileBag *TileBag) {
	for len(p.Tiles) < tileRackSize {
		tile, err := tileBag.drawTile()
		if err != nil {
			break
		}
		p.Tiles = append(p.Tiles, tile)
	}
}

func (p *Player) removeTile(tile int) {
	idx := slices.Index(p.Tiles, tile)
	slices.Delete(p.Tiles, idx, idx+1)
	p.Tiles = p.Tiles[:len(p.Tiles)-1]
}
