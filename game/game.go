package game

import (
	"game/assets"
	"game/vector"

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
	player  *Player
	Config  *Config
	enemies []*Enemy
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
	g.enemies = SpawnEnemies(5, 10, 50, 50)

	return g
}

func (g *Game) Update() error {
	g.player.Update()

	if inpututil.IsKeyJustPressed(ebiten.KeyF11) {
		g.Config.Fullscreen = !g.Config.Fullscreen
		ebiten.SetFullscreen(g.Config.Fullscreen)
	}

	for i, e := range g.enemies {
		for j, b := range g.player.bullets {
			if e.Collider().Intersects(b.Collider()) {
				g.enemies = append(g.enemies[:i], g.enemies[i+1:]...)
				g.player.bullets = append(g.player.bullets[:j], g.player.bullets[j+1:]...)
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

func SpawnEnemies(rows, cols int, spacingX, spacingY float64) []*Enemy {
	enemies := make([]*Enemy, 0)

	enemyWidth := float64(assets.EnemySprite.Bounds().Dx())
	enemyHeight := float64(assets.EnemySprite.Bounds().Dy())

	// Calculate total width and height of the enemy formation
	totalWidth := float64(cols)*(enemyWidth+spacingX) - spacingX

	// Calculate startX and startY to center the formation with the screen
	startX := (ScreenWidth - totalWidth) / 2
	startY := spacingY

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			x := startX + float64(col)*(spacingX+enemyWidth)
			y := startY + float64(row)*(spacingY+enemyHeight)
			pos := vector.Vector{
				X: x,
				Y: y,
			}
			enemy := NewEnemy(pos)
			enemies = append(enemies, enemy)
		}
	}
	return enemies
}
