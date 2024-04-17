package game

import (
	"game/assets"
	"game/pkg/utils"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Variant int

const (
	powerupSpeedPerSecond = 250.0
	powerupSpawnChance = 1
)

const (
	SpeedPowerup Variant = iota
	MovementPowerup
	ShieldPowerup
)

type Powerup struct {
	position utils.Vector
	sprite   *ebiten.Image
	variant  Variant
}

func NewPowerup(pos utils.Vector, variant Variant) *Powerup {
	var sprite *ebiten.Image

	switch variant {
	case SpeedPowerup:
		sprite = assets.SpeedPowerupSprite
	case MovementPowerup:
		sprite = assets.MovementPowerupSprite
	case ShieldPowerup:
		sprite = assets.ShieldPowerupSprite
	}

	return &Powerup{
		position: utils.Vector{
			X: pos.X,
			Y: pos.Y,
		},
		sprite: sprite,
		variant: variant,
	}
}

func generateRandomPowerup(pos utils.Vector) (*Powerup, bool) {
	randFloat := rand.Float64()

	if randFloat < powerupSpawnChance {
		randInt := rand.Intn(3)

		var variant Variant
		switch randInt {
		case 0:
			variant = SpeedPowerup
		case 1:
			variant = MovementPowerup
		case 2:
			variant = ShieldPowerup
		default:
			panic("Invalid random number generated")
		}
		return NewPowerup(pos, variant), true
	}

	return nil, false
}

func (p *Powerup) Update() error {
	speed := powerupSpeedPerSecond / float64(ebiten.TPS())

	p.position.Y += speed
	return nil
}

func (p *Powerup) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.position.X, p.position.Y)
	screen.DrawImage(p.sprite, op)
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
