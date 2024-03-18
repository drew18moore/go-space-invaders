package main

import (
	"game/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	err := ebiten.RunGame(game.NewGame())
	if err != nil {
		panic(err)
	}
}
