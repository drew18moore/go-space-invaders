package game

import (
	"fmt"
	"game/assets"
	"game/pkg/scenes"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
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
	sceneManager   *scenes.SceneManager
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
	g.sceneManager = scenes.NewSceneManager()

	return g
}

func (g *Game) Update() error {
	g.player.Update()
	g.enemyFormation.Update(g)
	g.input.Update(g)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
	g.enemyFormation.Draw(screen)

	text.Draw(screen, fmt.Sprintf("%06d", g.score), assets.ScoreFont, g.Config.ScreenWidth/2-100, 50, color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Reset() {
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
