package main

import (
	"slices"
)

const maxTilesInHand = 6

type Player struct {
	Id    string
	Tiles []int
}

func NewPlayer(id string) *Player {
	return &Player{
		Id:    id,
		Tiles: make([]int, 0),
	}
}

func (p *Player) fillTileRack(tileBag *TileBag) {
	for len(p.Tiles) < maxTilesInHand {
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
