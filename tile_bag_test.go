package main

import (
	"fmt"
	"testing"
)

func TestNewTileBag(t *testing.T) {
	t.Run("should initialize with 108 tiles", func(t *testing.T) {
		tileBag := NewTileBag(1)

		numTiles := len(tileBag.tiles)
		if numTiles != 108 {
			t.Errorf("tile bag should have 108 tiles, got %d", numTiles)
		}
	})

	type SeedTest struct {
		seed      int64
		wantOrder []int
	}
	seedTests := []SeedTest{
		{
			seed: 1,
			wantOrder: []int{
				90, 0, 28, 99, 54, 30, 102, 14, 74, 68, 93, 86, 23, 76, 92, 25, 101, 46, 85, 1, 13, 79, 7, 77, 82, 41, 10, 58, 56, 88, 60, 59, 22, 57, 38, 62, 71, 53, 107, 39, 63, 37, 64, 8, 32, 66, 96, 103, 81, 75, 34, 67, 3, 52, 84, 97, 80, 51, 21, 47, 87, 106, 33, 27, 18, 35, 94, 11, 95, 48, 4, 42, 5, 72, 104, 12, 2, 40, 55, 69, 16, 89, 83, 24, 73, 49, 31, 17, 19, 61, 91, 26, 43, 98, 36, 20, 78, 50, 29, 9, 15, 6, 105, 44, 45, 70, 100, 65,
			},
		},
		{
			seed: 99,
			wantOrder: []int{
				94, 50, 61, 20, 59, 36, 9, 58, 30, 92, 106, 31, 75, 1, 17, 82, 103, 73, 2, 78, 3, 44, 60, 80, 5, 7, 81, 13, 101, 48, 47, 55, 104, 53, 41, 43, 105, 95, 11, 84, 52, 18, 107, 29, 85, 19, 76, 70, 42, 83, 88, 72, 40, 96, 74, 22, 45, 86, 23, 15, 62, 21, 87, 14, 49, 26, 24, 12, 102, 69, 38, 27, 34, 37, 63, 89, 56, 90, 99, 97, 79, 93, 64, 25, 4, 77, 54, 100, 6, 91, 51, 57, 35, 8, 65, 67, 0, 33, 10, 16, 39, 32, 98, 46, 66, 71, 68, 28,
			},
		},
	}

	for _, test := range seedTests {
		t.Run(fmt.Sprintf("should create tiles in the same order for seed %d", test.seed), func(t *testing.T) {
			tileBag := NewTileBag(test.seed)

			for i := range test.wantOrder {
				if tileBag.tiles[i] != test.wantOrder[i] {
					t.Errorf("tiles not in expected order for seed. \ngot\n%v\nwant\n%v", tileBag.tiles, test.wantOrder)
					break
				}
			}
		})
	}
}

func TestDrawTile(t *testing.T) {
	t.Run("should draw tiles out of the bag in order", func(t *testing.T) {
		tileBag := NewTileBag(1)

		expectedOrder := []int{
			90, 0, 28, 99, 54, 30, 102, 14, 74, 68, 93, 86, 23, 76, 92, 25, 101, 46, 85, 1, 13, 79, 7, 77, 82, 41, 10, 58, 56, 88, 60, 59, 22, 57, 38, 62, 71, 53, 107, 39, 63, 37, 64, 8, 32, 66, 96, 103, 81, 75, 34, 67, 3, 52, 84, 97, 80, 51, 21, 47, 87, 106, 33, 27, 18, 35, 94, 11, 95, 48, 4, 42, 5, 72, 104, 12, 2, 40, 55, 69, 16, 89, 83, 24, 73, 49, 31, 17, 19, 61, 91, 26, 43, 98, 36, 20, 78, 50, 29, 9, 15, 6, 105, 44, 45, 70, 100, 65,
		}

		for i, want := range expectedOrder {
			got, err := tileBag.drawTile()
			if err != nil {
				t.Errorf("got error drawing tile %v", err)
			}
			if want != got {
				t.Errorf("wrong tile drawn at position %d. got %d want %d", i, got, want)
				break
			}
		}

		_, err := tileBag.drawTile()
		if err.Error() == "" {
			t.Errorf("expected drawTile() to return error if no tiles are left in the bag")
		}
	})
}
