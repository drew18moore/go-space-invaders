package game

import (
	"game/assets"
	"game/pkg/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	bulletSpeedPerSecond = 750.0
)

type Bullet struct {
	position  utils.Vector
	sprite    *ebiten.Image
	direction int8
}

func NewBullet(pos utils.Vector, direction int8) *Bullet {
	bounds := assets.LaserSprite.Bounds()

	var y float64
	if direction == 1 {
		y = pos.Y
	} else if direction == -1 {
		y = pos.Y - float64(bounds.Dy())
	}

	return &Bullet{
		position: utils.Vector{
			X: pos.X - (float64(bounds.Dx()) / 2),
			Y: y,
		},
		sprite: assets.LaserSprite,
		direction: direction,
	}
}

func (b *Bullet) Update() {
	speed := bulletSpeedPerSecond / float64(ebiten.TPS())

	b.position.Y += speed * float64(b.direction)
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.position.X, b.position.Y)
	screen.DrawImage(b.sprite, op)
}

func (p *Bullet) Collider() utils.Rect {
	bounds := p.sprite.Bounds()

	return utils.NewRect(
		p.position.X,
		p.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}
