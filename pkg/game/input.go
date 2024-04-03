package game

import (
	"game/pkg/utils"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Input struct{}

func (i *Input) Update(gameState *Game) error {
	// Fullscreen Keybind
	if inpututil.IsKeyJustPressed(ebiten.KeyF11) {
		gameState.Config.Fullscreen = !gameState.Config.Fullscreen
		ebiten.SetFullscreen(gameState.Config.Fullscreen)
	}

	// Player movement
	var delta utils.Vector
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		delta.Y = -gameState.player.movementSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		delta.X = -gameState.player.movementSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		delta.Y = gameState.player.movementSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		delta.X = gameState.player.movementSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		gameState.player.TryShoot()
	}

	gameState.player.TryMove(delta)

	return nil
}
