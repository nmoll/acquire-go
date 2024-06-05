package main

import (
	"fmt"
	"math/rand"
)

type TileBag struct {
	tiles []int
	seed  int64
}

func NewTileBag(seed int64) *TileBag {
	tiles := make([]int, totalCells)
	for i := range tiles {
		tiles[i] = i
	}

	tileBag := &TileBag{
		tiles: tiles,
		seed:  seed,
	}

	tileBag.shuffle()

	return tileBag
}

func (t *TileBag) drawTile() (int, error) {
	if len(t.tiles) == 0 {
		return 0, fmt.Errorf("no tiles left")
	}

	var tile int
	tile, t.tiles = t.tiles[0], t.tiles[1:]
	return tile, nil
}

func (t *TileBag) shuffle() {
	randSource := rand.NewSource(t.seed)
	random := rand.New(randSource)
	random.Shuffle(totalCells, func(i, j int) {
		t.tiles[i], t.tiles[j] = t.tiles[j], t.tiles[i]
	})
}
