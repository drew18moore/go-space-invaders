package game

import (
	"game/assets"
	"game/rect"
	"game/vector"

	"github.com/hajimehoshi/ebiten/v2"
)

type Enemy struct {
	position vector.Vector
	sprite   *ebiten.Image
}

func NewEnemy(pos vector.Vector) *Enemy {
	sprite := assets.EnemySprite

	return &Enemy{
		position: pos,
		sprite: sprite,
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
