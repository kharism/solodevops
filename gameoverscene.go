package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/joelschutz/stagehand"
)

type GameOverScene struct {
	State    *GlobalGameState
	director *stagehand.SceneDirector[*GlobalGameState]
}

func (c *GameOverScene) Update() error {

	return nil
}
func (c *GameOverScene) Draw(screen *ebiten.Image) {

}
func (s *GameOverScene) Load(state *GlobalGameState, director stagehand.SceneController[*GlobalGameState]) {
	s.director = director.(*stagehand.SceneDirector[*GlobalGameState])
}
func (s *GameOverScene) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}
func (s *GameOverScene) Unload() *GlobalGameState {
	// your unload code
	// s.scene.Events[0].Execute(s.scene)

	return s.State
}
