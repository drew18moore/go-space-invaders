package game

import (
	"game/assets"
	"game/rect"
	"game/timer"
	"game/vector"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	shootCooldown = time.Millisecond * 500
)

type Player struct {
	position      vector.Vector
	sprite        *ebiten.Image
	game          *Game
	shootCooldown *timer.Timer
	bullets       []*Bullet
}

func NewPlayer(game *Game) *Player {
	sprite := assets.PlayerSprite
	bounds := sprite.Bounds()

	return &Player{
		position: vector.Vector{
			X: float64(game.Config.ScreenWidth)/2 - float64(bounds.Dx())/2,
			Y: float64(game.Config.ScreenHeight) - float64(bounds.Dy())/2,
		},
		sprite:        assets.PlayerSprite,
		game:          game,
		shootCooldown: timer.NewTimer(shootCooldown),
	}
}

func (p *Player) Update() {
	p.shootCooldown.Update()

	for _, b := range p.bullets {
		b.Update()
	}

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
	if ebiten.IsKeyPressed(ebiten.KeySpace) && p.shootCooldown.IsReady() {
		p.shootCooldown.Reset()

		bounds := p.sprite.Bounds()
		spawnPos := vector.Vector{
			X: p.position.X + (float64(bounds.Dx()) / 2),
			Y: p.position.Y,
		}
		bullet := NewBullet(spawnPos, -1)
		p.AddBullet(bullet)
	}

	if delta.X != 0 && delta.Y != 0 {
		factor := speed / math.Sqrt(delta.X*delta.X+delta.Y*delta.Y)
		delta.X *= factor
		delta.Y *= factor
	}

	newX := p.position.X + delta.X
	newY := p.position.Y + delta.Y

	screenWidth := float64(p.game.Config.ScreenWidth)
	screenHeight := float64(p.game.Config.ScreenHeight)
	playerWidth := float64(p.sprite.Bounds().Dx())
	playerHeight := float64(p.sprite.Bounds().Dy())

	newX = math.Max(0, math.Min(newX, screenWidth-playerWidth))
	newY = math.Max(0, math.Min(newY, screenHeight-playerHeight))

	p.position.X = newX
	p.position.Y = newY
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

func (p *Player) Collider() rect.Rect {
	bounds := p.sprite.Bounds()

	return rect.NewRect(
		p.position.X,
		p.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}
