package player

import (
	"game/assets"
	"game/vector"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	position vector.Vector
	sprite   *ebiten.Image
}

func NewPlayer() *Player {
	return &Player{
		position: vector.Vector{X: 100, Y: 100},
		sprite:   assets.PlayerSprite,
	}
}

func (p *Player) Update() {
	speed := float64(300 / ebiten.TPS())

	var delta vector.Vector
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		delta.Y = speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		delta.Y = -speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		delta.X = -speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		delta.X = speed
	}

	if delta.X != 0 && delta.Y != 0 {
		factor := speed / math.Sqrt(delta.X*delta.X+delta.Y*delta.Y)
		delta.X *= factor
		delta.Y *= factor
	}

	p.position.X += delta.X
	p.position.Y += delta.Y
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.position.X, p.position.Y)
	screen.DrawImage(p.sprite, op)
}
