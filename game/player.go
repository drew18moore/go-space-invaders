package game

import (
	"game/assets"
	"game/vector"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	position vector.Vector
	sprite   *ebiten.Image
	game *Game
}

func NewPlayer(game *Game) *Player {
	sprite := assets.PlayerSprite
	bounds := sprite.Bounds()

	return &Player{
		position: vector.Vector{
			X: float64(game.Config.ScreenWidth)/2 - float64(bounds.Dx())/2,
			Y: float64(game.Config.ScreenHeight)/2 - float64(bounds.Dy())/2,
		},
		sprite: assets.PlayerSprite,
		game: game,
	}
}

func (p *Player) Update() {
	speed := float64(300 / ebiten.TPS())

	var delta vector.Vector
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		delta.Y = speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		delta.Y = -speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		delta.X = -speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
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
