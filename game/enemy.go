package game

import (
	"game/assets"
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
