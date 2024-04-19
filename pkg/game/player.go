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
	shootTimer    *utils.Timer
	bullets       []*Bullet
	movementSpeed float64
	bulletSpeed   float64
	shields       int
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
		shootTimer:    utils.NewTimer(shootCooldown),
		movementSpeed: float64(300 / ebiten.TPS()),
		bulletSpeed:   bulletSpeedPerSecond,
		shields:       3,
	}
}

func (p *Player) Update() {
	p.shootTimer.Update()

	for _, b := range p.bullets {
		b.Update()
	}

	for i := 0; i < len(p.game.powerups); i++ {
		powerup := p.game.powerups[i]
		if p.Collider().Intersects(powerup.Collider()) {
			p.game.powerups = append(p.game.powerups[:i], p.game.powerups[i+1:]...)
			i--
			switch powerup.variant {
			case SpeedPowerup:
				p.shootTimer.DecreaseTimer(time.Millisecond * 25)
				if p.shootTimer.CurrentTarget() < 1 {
					p.shootTimer.SetDuration(time.Millisecond * 34)
				}
				if p.bulletSpeed > 4250 {
					p.bulletSpeed = 4250
				}
				p.bulletSpeed += 250.0
				p.game.score++
			case MovementPowerup:
				p.movementSpeed += 0.25
				if p.movementSpeed >= 12.5 {
					p.movementSpeed = 12.5
				}
				p.game.score++
			case ShieldPowerup:
				p.shields++
				if p.shields > 5 {
					p.shields = 5
				}
			}
		}
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
	if p.shootTimer.IsReady() {
		p.shootTimer.Reset()

		bounds := p.sprite.Bounds()
		spawnPos := utils.Vector{
			X: p.position.X + (float64(bounds.Dx()) / 2),
			Y: p.position.Y,
		}
		bullet := NewBullet(spawnPos, -1, p.bulletSpeed, PlayerBullet)
		p.AddBullet(bullet)
	}
}
