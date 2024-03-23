package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	ScreenWidth  = 1920
	ScreenHeight = 1080
)

type Config struct {
	ScreenWidth  int
	ScreenHeight int
	Fullscreen bool
}

type Game struct {
	player *Player
	Config *Config
	enemies []*Enemy
}

func NewGame() *Game {
	g := &Game{
		Config: &Config{
			ScreenWidth:  ScreenWidth,
			ScreenHeight: ScreenHeight,
			Fullscreen: false,
		},
	}

	g.player = NewPlayer(g)
	g.enemies = []*Enemy{ NewEnemy(g) }
	return g
}

func (g *Game) Update() error {
	g.player.Update()

	if inpututil.IsKeyJustPressed(ebiten.KeyF11) {
		g.Config.Fullscreen = !g.Config.Fullscreen
		ebiten.SetFullscreen(g.Config.Fullscreen)
	}

	for _, e := range g.enemies {
		for _, b := range g.player.bullets {
			if e.Collider().Intersects(b.Collider()) {
				fmt.Println("HIT")
			}
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)

	for _, e := range g.enemies {
		e.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
