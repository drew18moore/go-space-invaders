package game

import (
	"game/assets"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type TitleScene struct{
	gameState *Game
}

func (s *TitleScene) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		s.gameState.sceneManager.GoTo(&GameScene{
			gameState: s.gameState,
		})
	}

	return nil
}

func (s *TitleScene) Draw(screen *ebiten.Image) {
	msg := "Press SPACE to play"
	bounds, _ := font.BoundString(assets.ScoreFont, msg)
	text.Draw(screen, msg, assets.ScoreFont, s.gameState.Config.ScreenWidth/2-bounds.Max.X.Ceil()/2, s.gameState.Config.ScreenHeight/2-bounds.Max.Y.Ceil()/2, color.White)
}
