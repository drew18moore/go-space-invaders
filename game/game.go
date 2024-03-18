package game

import (
	"game/player"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	player *player.Player
}

func NewGame() *Game {
	return &Game{
		player: player.NewPlayer(),
	}
}

func (g *Game) Update() error {
	g.player.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
