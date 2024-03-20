package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

type Config struct {
	ScreenWidth  int
	ScreenHeight int
}

type Game struct {
	player *Player
	Config *Config
}

func NewGame() *Game {
	g := &Game{
		Config: &Config{
			ScreenWidth:  ScreenWidth,
			ScreenHeight: ScreenHeight,
		},
	}

	g.player = NewPlayer(g)
	return g
}

func (g *Game) Update() error {
	g.player.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
