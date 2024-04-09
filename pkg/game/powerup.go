package game

import (
	"game/assets"
	"game/pkg/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type Variant int

const (
	SpeedPowerup Variant = iota
)

type Powerup struct {
	position utils.Vector
	sprite   *ebiten.Image
	variant  int
}

func NewPowerup(pos utils.Vector, variant Variant) *Powerup {
	var sprite *ebiten.Image

	switch variant {
	case SpeedPowerup:
		sprite = assets.SpeedPowerupSprite
	}

	return &Powerup{
		position: utils.Vector{
			X: pos.X,
			Y: pos.Y,
		},
		sprite: sprite,
	}
}

func (p *Powerup) Collider() utils.Rect {
	bounds := p.sprite.Bounds()

	return utils.NewRect(
		p.position.X,
		p.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}
