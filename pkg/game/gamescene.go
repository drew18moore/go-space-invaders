package game

import (
	"fmt"
	"game/assets"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type GameScene struct {
	gameState *Game
}

func (s *GameScene) Update() error {
	g := s.gameState

	g.player.Update()
	g.enemyFormation.Update(g)
	g.input.Update(g)

	return nil
}

func (s *GameScene) Draw(screen *ebiten.Image) {
	g := s.gameState

	// Draw background image
	repeatX := (g.Config.ScreenWidth + s.gameState.background.Bounds().Dx() - 1) / s.gameState.background.Bounds().Dx()
	repeatY := (g.Config.ScreenHeight + s.gameState.background.Bounds().Dy() - 1) / s.gameState.background.Bounds().Dy()
	for y := 0; y < repeatY; y++ {
		for x := 0; x < repeatX; x++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x * s.gameState.background.Bounds().Dx()), float64(y * s.gameState.background.Bounds().Dy()))
			screen.DrawImage(s.gameState.background, op)
		}
	}
	
	g.player.Draw(screen)
	g.enemyFormation.Draw(screen)


	text.Draw(screen, fmt.Sprintf("%06d", g.score), assets.ScoreFont, g.Config.ScreenWidth/2-100, 50, color.White)
}
