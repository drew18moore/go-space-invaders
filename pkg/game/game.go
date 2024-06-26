package game

import (
	"game/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 2048
	ScreenHeight = 1536
)

type Config struct {
	ScreenWidth  int
	ScreenHeight int
	Fullscreen   bool
}

type Game struct {
	sceneManager   *SceneManager
	player         *Player
	Config         *Config
	enemyFormation EnemyFormation
	score          int
	input          *Input
	background     *ebiten.Image
	powerups       []*Powerup
}

func NewGame() *Game {
	g := &Game{
		Config: &Config{
			ScreenWidth:  ScreenWidth,
			ScreenHeight: ScreenHeight,
			Fullscreen:   false,
		},
	}

	g.player = NewPlayer(g)
	g.enemyFormation = NewEnemyFormation(5, 10, 50, 50)
	g.sceneManager = NewSceneManager(g)
	g.background = assets.Background

	return g
}

func (g *Game) Update() error {
	g.sceneManager.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.sceneManager.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Reset() {
	g.sceneManager.GoTo(&TitleScene{
		gameState: g,
	})
	g.player = NewPlayer(g)
	g.enemyFormation = NewEnemyFormation(5, 10, 50, 50)
	g.score = 0
}
