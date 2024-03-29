package game

import (
	"game/assets"
	"game/pkg/utils"
	"time"
)

const (
	enemyShootCooldown = time.Millisecond * 1000
)

type EnemyFormation struct {
	enemies           []*Enemy
	movementDirection int8 // 1 for moving right, -1 for moving left
	movementSpeed     float64
	shootCooldown     *utils.Timer
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
		shootCooldown:     utils.NewTimer(enemyShootCooldown),
	}
}