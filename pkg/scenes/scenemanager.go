package scenes

import "github.com/hajimehoshi/ebiten/v2"

type Scene interface {
	Update() error
	Draw(screen *ebiten.Image)
}

type SceneManager struct {
	current Scene
}

func NewSceneManager() *SceneManager {
	return &SceneManager{
		current: &TitleScene{},
	}
}

func (s *SceneManager) Update() error {
	return nil
}

func (s *SceneManager) Draw() {

}

func (s *SceneManager) GoTo(scene Scene) {
	s.current = scene
}
