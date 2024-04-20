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
	stage             int
	rows              int
	cols              int
	spacingX          float64
	spacingY          float64
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
			enemy := NewEnemy(pos, 2)
			enemies = append(enemies, enemy)
		}
	}
	return EnemyFormation{
		enemies:           enemies,
		movementDirection: 1,
		movementSpeed:     1,
		shootTimer:        utils.NewTimer(enemyShootCooldown),
		stage:             1,
		rows:              rows,
		cols:              cols,
		spacingX:          spacingX,
		spacingY:          spacingY,
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
	for i := 0; i < len(ef.enemies); i++ {
		enemy := ef.enemies[i]
		for j := 0; j < len(gameState.player.bullets); j++ {
			bullet := gameState.player.bullets[j]
			if enemy.Collider().Intersects(bullet.Collider()) {
				gameState.player.bullets = append(gameState.player.bullets[:j], gameState.player.bullets[j+1:]...)
				j--
				enemy.currentHealth--
				
				if enemy.currentHealth <= 0 {
					ef.enemies = append(ef.enemies[:i], ef.enemies[i+1:]...)

					i--

					ef.movementSpeed += 0.10
					gameState.score++
					ef.shootTimer.DecreaseTimer(time.Millisecond * 2)

					pu, ok := generateRandomPowerup(enemy.position)
					if ok {
						gameState.powerups = append(gameState.powerups, pu)
					}

					enemy.currentHealth--
					break
				}
			}
		}
	}

	//Collision b/w enemy bullet and player
	for i := 0; i < len(ef.bullets); i++ {
		bullet := ef.bullets[i]
		if gameState.player.Collider().Intersects(bullet.Collider()) {
			ef.bullets = append(ef.bullets[:i], ef.bullets[i+1:]...)
			i--
			if gameState.player.shields <= 0 {
				gameState.Reset()
			} else {
				gameState.player.shields--
			}
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
		bullet := NewBullet(spawnPos, 1, bulletSpeedPerSecond, EnemyBullet)
		ef.bullets = append(ef.bullets, bullet)
	}

	if len(ef.enemies) == 0 {
		ef.NextWave()
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

func (ef *EnemyFormation) NextWave() {
	ef.bullets = nil

	ef.stage++
	enemies := make([]*Enemy, 0)

	enemyWidth := float64(assets.EnemySprite.Bounds().Dx())
	enemyHeight := float64(assets.EnemySprite.Bounds().Dy())

	// Calculate total width and height of the enemy formation
	totalWidth := float64(ef.cols)*(enemyWidth+ef.spacingX) - ef.spacingX

	// Calculate startX and startY to center the formation with the screen
	startX := (ScreenWidth - totalWidth) / 2
	startY := ef.spacingY

	for row := 0; row < ef.rows; row++ {
		for col := 0; col < ef.cols; col++ {
			x := startX + float64(col)*(ef.spacingX+enemyWidth)
			y := startY + float64(row)*(ef.spacingY+enemyHeight)
			pos := utils.Vector{
				X: x,
				Y: y,
			}
			enemy := NewEnemy(pos, 2+ef.stage-1)
			enemies = append(enemies, enemy)
		}
	}
	ef.enemies = enemies
	ef.movementSpeed = 1
	ef.shootTimer = utils.NewTimer(enemyShootCooldown)
}
