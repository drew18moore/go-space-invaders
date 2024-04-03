package game

import (
	"fmt"
	"game/assets"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type GameScene struct{
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

	g.player.Draw(screen)
	g.enemyFormation.Draw(screen)

	text.Draw(screen, fmt.Sprintf("%06d", g.score), assets.ScoreFont, g.Config.ScreenWidth/2-100, 50, color.White)
}