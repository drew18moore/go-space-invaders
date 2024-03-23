package game

import (
	"game/assets"
	"game/rect"
	"game/timer"
	"game/vector"
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
	shootCooldown     *timer.Timer
	bullets           []*Bullet
}

type Enemy struct {
	position vector.Vector
	sprite   *ebiten.Image
}

func NewEnemy(pos vector.Vector) *Enemy {
	sprite := assets.EnemySprite

	return &Enemy{
		position: pos,
		sprite:   sprite,
	}
}

func (e *Enemy) Update() {

}

func (e *Enemy) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(e.position.X, e.position.Y)
	screen.DrawImage(e.sprite, op)
}

func (p *Enemy) Collider() rect.Rect {
	bounds := p.sprite.Bounds()

	return rect.NewRect(
		p.position.X,
		p.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
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
			pos := vector.Vector{
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
		shootCooldown:     timer.NewTimer(enemyShootCooldown),
	}
}
