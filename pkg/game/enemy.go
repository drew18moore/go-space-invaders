package game

import (
	"game/assets"
	"game/pkg/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type Enemy struct {
	position utils.Vector
	sprite   *ebiten.Image
}

func NewEnemy(pos utils.Vector) *Enemy {
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

func (p *Enemy) Collider() utils.Rect {
	bounds := p.sprite.Bounds()

	return utils.NewRect(
		p.position.X,
		p.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}
