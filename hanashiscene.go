package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/joelschutz/stagehand"
	"github.com/kharism/hanashi/core"
)

type HanashiScene struct {
	scene    *core.Scene
	State    *GlobalGameState
	director *stagehand.SceneDirector[*GlobalGameState]
}

func (m *HanashiScene) Update() error {
	e := m.scene.Update()
	if e != nil {
		return e
	}

	return nil
}
func (m *HanashiScene) Draw(screen *ebiten.Image) {
	m.scene.Draw(screen)
	// m.SkipButton.Draw(screen)
	// txt := "click to continue"
	// txtOpt := text.DrawOptions{}
	// txtOpt.ColorScale.ScaleWithColor(RED)
	// txtOpt.GeoM.Scale(0.5, 0.5)
	// text.Draw(screen, txt, face, &txtOpt)
}

func (s *HanashiScene) Load(state *GlobalGameState, director stagehand.SceneController[*GlobalGameState]) {
	s.director = director.(*stagehand.SceneDirector[*GlobalGameState]) // This type assertion is important
	s.scene.Events[0].Execute(s.scene)
	// s.scene.EventIndex = 0

}
func (s *HanashiScene) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}
func (s *HanashiScene) Unload() *GlobalGameState {
	// your unload code
	// s.scene.Events[0].Execute(s.scene)
	s.scene.EventIndex = 0
	s.scene.CurCharName = ""
	s.scene.CurDialog = ""
	// s.scene.ViewableCharacters = []*core.Character{}
	s.scene.VisibleDialog = ""

	return s.State
}
