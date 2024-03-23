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
	game     *Game
}

func NewEnemy(game *Game) *Enemy {
	sprite := assets.EnemySprite

	return &Enemy{
		position: vector.Vector{
			X: 0,
			Y: 0,
		},
		sprite: sprite,
		game:   game,
	}
}

func (e *Enemy) Update() {

}

func (e *Enemy) Draw(screen *ebiten.Image) {
	// op := &ebiten.DrawImageOptions{}
	screen.DrawImage(e.sprite, nil)
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