package game

import (
	"game/assets"
	"game/pkg/utils"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Enemy struct {
	position       utils.Vector
	sprite         *ebiten.Image
	maxHealth     int
	currentHealth int
}

func NewEnemy(pos utils.Vector, health int) *Enemy {
	sprite := assets.EnemySprite
	maxHealth := health

	return &Enemy{
		position:       pos,
		sprite:         sprite,
		maxHealth:     maxHealth,
		currentHealth: maxHealth,
	}
}

func (e *Enemy) Update() {

}

func (e *Enemy) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(e.position.X, e.position.Y)
	screen.DrawImage(e.sprite, op)

	// Draw health bar
	if e.currentHealth < e.maxHealth {
		barWidth := 50
		barHeight := 5
		barX := e.position.X + float64(e.sprite.Bounds().Dx()) / 2 - float64(barWidth) / 2
		barY := e.position.Y - float64(barHeight) - 5
	
		percentage := float32(e.currentHealth) / float32(e.maxHealth)
	
		var barColor color.Color
		if percentage > 0.5 {
			barColor = color.RGBA{0, 255, 0, 255}
		} else if percentage > 0.2 {
			barColor = color.RGBA{255, 255, 0, 255}
		} else {
			barColor = color.RGBA{255, 0, 0, 255}
		}
	
		vector.DrawFilledRect(screen, float32(barX), float32(barY), float32(barWidth), float32(barHeight), color.Black, false)
		vector.DrawFilledRect(screen, float32(barX), float32(barY), float32(barWidth)*percentage, float32(barHeight), barColor, false)
	}
}

func (p *Enemy) Collider() utils.Rect {
	bounds := p.sprite.Bounds()

	return utils.NewRect(
		p.position.X,
		p.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}
