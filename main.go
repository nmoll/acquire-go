package main

import (
	"fmt"
)

func main() {
	// actions := make([]Action, 0)
	// actions = append(actions, Action{
	// 	Type: PlaceTile,
	// })

	game := NewGameEngine()

	game.computeState()

	fmt.Println("Running acquire", len(game.Actions))
}
