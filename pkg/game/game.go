package game

import (
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

// TODOS:
// Create SceneManager struct
// Create structs for each scene (e.x title and game scenes)
// Move current Draw and Update logic into game scene
// Add SceneManager to game struct
// When SceneManager inits, set curr scene to title screen
// Create GoTo method
// Add event listener to title scene for space and GoTo game scene if pressed
