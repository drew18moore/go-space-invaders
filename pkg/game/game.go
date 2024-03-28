package game

import (
	"fmt"
	"game/assets"
	"game/pkg/scenes"
	"game/pkg/utils"
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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

	for _, b := range g.enemyFormation.bullets {
		b.Update()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyF11) {
		g.Config.Fullscreen = !g.Config.Fullscreen
		ebiten.SetFullscreen(g.Config.Fullscreen)
	}

	// Collision b/w player bullet and enemy
	for i, e := range g.enemyFormation.enemies {
		for j, b := range g.player.bullets {
			if e.Collider().Intersects(b.Collider()) {
				g.enemyFormation.enemies = append(g.enemyFormation.enemies[:i], g.enemyFormation.enemies[i+1:]...)
				g.player.bullets = append(g.player.bullets[:j], g.player.bullets[j+1:]...)
				g.enemyFormation.movementSpeed += 0.25
				g.score++
			}
		}
	}

	//Collision b/w enemy bullet and player
	for _, b := range g.enemyFormation.bullets {
		if g.player.Collider().Intersects(b.Collider()) {
			g.Reset()
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

	// Handle enemy shooting
	g.enemyFormation.shootCooldown.Update()
	if g.enemyFormation.shootCooldown.IsReady() && len(g.enemyFormation.enemies) > 0 {
		g.enemyFormation.shootCooldown.Reset()
		r := rand.New(rand.NewSource(time.Now().Unix()))
		randEnemy := g.enemyFormation.enemies[r.Intn(len(g.enemyFormation.enemies))]

		bounds := randEnemy.sprite.Bounds()
		spawnPos := utils.Vector{
			X: randEnemy.position.X + (float64(bounds.Dx()) / 2),
			Y: randEnemy.position.Y + (float64(bounds.Dy())),
		}
		bullet := NewBullet(spawnPos, 1)
		g.enemyFormation.bullets = append(g.enemyFormation.bullets, bullet)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)

	for _, e := range g.enemyFormation.enemies {
		e.Draw(screen)
	}

	for _, b := range g.enemyFormation.bullets {
		b.Draw(screen)
	}

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
