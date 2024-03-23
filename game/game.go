package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
	player         *Player
	Config         *Config
	enemyFormation EnemyFormation
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

	return g
}

func (g *Game) Update() error {
	g.player.Update()

	if inpututil.IsKeyJustPressed(ebiten.KeyF11) {
		g.Config.Fullscreen = !g.Config.Fullscreen
		ebiten.SetFullscreen(g.Config.Fullscreen)
	}

	for i, e := range g.enemyFormation.enemies {
		for j, b := range g.player.bullets {
			if e.Collider().Intersects(b.Collider()) {
				g.enemyFormation.enemies = append(g.enemyFormation.enemies[:i], g.enemyFormation.enemies[i+1:]...)
				g.player.bullets = append(g.player.bullets[:j], g.player.bullets[j+1:]...)
				g.enemyFormation.movementSpeed += 0.25
			}
		}
	}

	// Move the entire formation based on the movement direction
	for _, e := range g.enemyFormation.enemies {
		e.position.X += float64(g.enemyFormation.movementDirection) * g.enemyFormation.movementSpeed
	}

	// Check if the formation has reached the edges of the screen
	// and reverse the direction if necessary
	for _, e := range g.enemyFormation.enemies {
		if g.enemyFormation.movementDirection == 1 && e.position.X+e.Collider().Width >= float64(g.Config.ScreenWidth) {
			g.enemyFormation.movementDirection = -1
			break
		} else if g.enemyFormation.movementDirection == -1 && e.position.X <= 0 {
			g.enemyFormation.movementDirection = 1
			break
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)

	for _, e := range g.enemyFormation.enemies {
		e.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
