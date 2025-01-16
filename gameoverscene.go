package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/joelschutz/stagehand"
)

type GameOverScene struct {
	State    *GlobalGameState
	director *stagehand.SceneDirector[*GlobalGameState]
}

func (c *GameOverScene) Update() error {
	if inpututil.IsKeyJustReleased(ebiten.KeySpace) {
		c.director.ProcessTrigger(TriggerToMain)
	}
	return nil
}
func (c *GameOverScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	geom := ebiten.GeoM{}
	geom.Scale(4, 4)
	geom.Translate(300, 200)
	colorm := ebiten.ColorScale{}
	colorm.Scale(1, 0, 0, 1)
	text.Draw(screen, "YOU DIED", PixelFontFace, &text.DrawOptions{
		LayoutOptions: text.LayoutOptions{
			PrimaryAlign: text.AlignCenter,
		},
		DrawImageOptions: ebiten.DrawImageOptions{
			GeoM:       geom,
			ColorScale: colorm,
		},
	})
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
