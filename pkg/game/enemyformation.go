package game

import (
	"game/assets"
	"game/pkg/utils"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	enemyShootCooldown = time.Millisecond * 1000
)

type EnemyFormation struct {
	enemies           []*Enemy
	movementDirection int8 // 1 for moving right, -1 for moving left
	movementSpeed     float64
	shootTimer        *utils.Timer
	bullets           []*Bullet
}

func NewEnemyFormation(rows, cols int, spacingX, spacingY float64) EnemyFormation {
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
			pos := utils.Vector{
				X: x,
				Y: y,
			}
			enemy := NewEnemy(pos)
			enemies = append(enemies, enemy)
		}
	}
	return EnemyFormation{
		enemies:           enemies,
		movementDirection: 1,
		movementSpeed:     1,
		shootTimer:        utils.NewTimer(enemyShootCooldown),
	}
}

func (ef *EnemyFormation) Update(gameState *Game) error {
	for _, b := range ef.bullets {
		b.Update()
	}

	// Move the entire formation based on the movement direction
	for _, e := range ef.enemies {
		e.position.X += float64(ef.movementDirection) * ef.movementSpeed
	}

	// Check if the formation has reached the edges of the screen
	// and reverse the direction if necessary
	for _, e := range ef.enemies {
		if ef.movementDirection == 1 && e.position.X+e.Collider().Width >= float64(gameState.Config.ScreenWidth) {
			ef.movementDirection = -1
			break
		} else if ef.movementDirection == -1 && e.position.X <= 0 {
			ef.movementDirection = 1
			break
		}
	}

	// Collision b/w player bullet and enemy
	for i, e := range ef.enemies {
		for j, b := range gameState.player.bullets {
			if e.Collider().Intersects(b.Collider()) {
				ef.enemies = append(ef.enemies[:i], ef.enemies[i+1:]...)
				gameState.player.bullets = append(gameState.player.bullets[:j], gameState.player.bullets[j+1:]...)
				ef.movementSpeed += 0.25
				gameState.score++
				ef.shootTimer.DecreaseTimer(10)

				_, ok := generateRandomPowerup(e.position) 

				if ok {
					gameState.powerups = append(gameState.powerups, NewPowerup(e.position, SpeedPowerup))
				}
			}
		}
	}

	//Collision b/w enemy bullet and player
	for _, b := range ef.bullets {
		if gameState.player.Collider().Intersects(b.Collider()) {
			gameState.Reset()
		}
	}

	// Handle enemy shooting
	ef.shootTimer.Update()
	if ef.shootTimer.IsReady() && len(ef.enemies) > 0 {
		ef.shootTimer.Reset()
		r := rand.New(rand.NewSource(time.Now().Unix()))
		randEnemy := ef.enemies[r.Intn(len(ef.enemies))]

		bounds := randEnemy.sprite.Bounds()
		spawnPos := utils.Vector{
			X: randEnemy.position.X + (float64(bounds.Dx()) / 2),
			Y: randEnemy.position.Y + (float64(bounds.Dy())),
		}
		bullet := NewBullet(spawnPos, 1, EnemyBullet)
		ef.bullets = append(ef.bullets, bullet)
	}

	return nil
}

func (ef *EnemyFormation) Draw(screen *ebiten.Image) {
	for _, e := range ef.enemies {
		e.Draw(screen)
	}

	for _, b := range ef.bullets {
		b.Draw(screen)
	}
}
