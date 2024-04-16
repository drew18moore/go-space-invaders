package assets

import (
	"embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//go:embed *
var assets embed.FS

var Background = mustLoadImage("background.png")

var PlayerSprite = mustLoadImage("player.png")
var PlayerLaserSprite = mustLoadImage("laser-player.png")
var EnemyLaserSprite = mustLoadImage("laser-enemy.png")
var EnemySprite = mustLoadImage("enemy.png")

var SpeedPowerupSprite = mustLoadImage("speed-powerup.png")
var MovementPowerupSprite = mustLoadImage("movement-powerup.png")
var ShieldPowerupSprite = mustLoadImage("shield-powerup.png")

var ScoreFont = mustLoadFont("font.ttf")


func mustLoadImage(name string) *ebiten.Image {
	f, err := assets.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}

func mustLoadFont(name string) font.Face {
	f, err := assets.ReadFile(name)
	if err != nil {
		panic(err)
	}

	tt, err := opentype.Parse(f)
	if err != nil {
		panic(err)
	}

	face, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    48,
		DPI:     72,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		panic(err)
	}

	return face
}
