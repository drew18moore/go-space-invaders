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
	speed     float64
}

type BulletType int

const (
	PlayerBullet BulletType = iota
	EnemyBullet
)

func NewBullet(pos utils.Vector, direction int8, speed float64, bulletType BulletType) *Bullet {
	var sprite *ebiten.Image

	switch bulletType {
	case PlayerBullet:
		sprite = assets.PlayerLaserSprite
	case EnemyBullet:
		sprite = assets.EnemyLaserSprite
	default:
		sprite = assets.PlayerLaserSprite
	}

	bounds := assets.PlayerLaserSprite.Bounds()

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
		sprite:    sprite,
		direction: direction,
		speed: speed,
	}
}

func (b *Bullet) Update() {
	speed := b.speed / float64(ebiten.TPS())
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
