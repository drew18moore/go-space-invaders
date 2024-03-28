package game

import "github.com/hajimehoshi/ebiten/v2"

type Scene interface {
	Update() error
	Draw(screen *ebiten.Image)
}

type SceneManager struct {
	current Scene
}

func (s *SceneManager) Update() error {
	return nil
}

func (s *SceneManager) Draw() {

}
