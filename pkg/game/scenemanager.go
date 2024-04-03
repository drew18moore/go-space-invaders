package game

import "github.com/hajimehoshi/ebiten/v2"

type Scene interface {
	Update() error
	Draw(screen *ebiten.Image)
}

type SceneManager struct {
	current Scene
}

func NewSceneManager(g *Game) *SceneManager {
	s := &SceneManager{}
	s.current = &TitleScene{
		gameState: g,
	}

	return s
}

func (s *SceneManager) Update() error {
	s.current.Update()
	return nil
}

func (s *SceneManager) Draw(screen *ebiten.Image) {
	s.current.Draw(screen)
}

func (s *SceneManager) GoTo(scene Scene) {
	s.current = scene
}
