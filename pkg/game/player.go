package game

import (
	"game/assets"
	"game/pkg/utils"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	shootCooldown = time.Millisecond * 500
)

type Player struct {
	position      utils.Vector
	sprite        *ebiten.Image
	game          *Game
	shootCooldown *utils.Timer
	bullets       []*Bullet
	movementSpeed float64
}

func NewPlayer(game *Game) *Player {
	sprite := assets.PlayerSprite
	bounds := sprite.Bounds()

	return &Player{
		position: utils.Vector{
			X: float64(game.Config.ScreenWidth)/2 - float64(bounds.Dx())/2,
			Y: float64(game.Config.ScreenHeight) - float64(bounds.Dy())/2,
		},
		sprite:        assets.PlayerSprite,
		game:          game,
		shootCooldown: utils.NewTimer(shootCooldown),
		movementSpeed: float64(300 / ebiten.TPS()),
	}
}

func (p *Player) Update() {
	p.shootCooldown.Update()

	for _, b := range p.bullets {
		b.Update()
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	for _, b := range p.bullets {
		b.Draw(screen)
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.position.X, p.position.Y)
	screen.DrawImage(p.sprite, op)
}

func (p *Player) AddBullet(b *Bullet) {
	p.bullets = append(p.bullets, b)
}

func (p *Player) Collider() utils.Rect {
	bounds := p.sprite.Bounds()

	return utils.NewRect(
		p.position.X,
		p.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}

func (p *Player) TryMove(vector utils.Vector) {
	if vector.X != 0 && vector.Y != 0 {
		factor := p.movementSpeed / math.Sqrt(vector.X*vector.X+vector.Y*vector.Y)
		vector.X *= factor
		vector.Y *= factor
	}

	newX := p.position.X + vector.X
	newY := p.position.Y + vector.Y

	screenWidth := float64(p.game.Config.ScreenWidth)
	screenHeight := float64(p.game.Config.ScreenHeight)
	playerWidth := float64(p.sprite.Bounds().Dx())
	playerHeight := float64(p.sprite.Bounds().Dy())

	newX = math.Max(0, math.Min(newX, screenWidth-playerWidth))
	newY = math.Max(0, math.Min(newY, screenHeight-playerHeight))

	p.position.X = newX
	p.position.Y = newY
}

func (p *Player) TryShoot() {
	if p.shootCooldown.IsReady() {
		p.shootCooldown.Reset()

		bounds := p.sprite.Bounds()
		spawnPos := utils.Vector{
			X: p.position.X + (float64(bounds.Dx()) / 2),
			Y: p.position.Y,
		}
		bullet := NewBullet(spawnPos, -1, PlayerBullet)
		p.AddBullet(bullet)
	}
}
