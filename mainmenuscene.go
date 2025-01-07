package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/joelschutz/stagehand"
	//"github.com/kharism/grimoiregunner/scene/assets"
)

type MainMenuScene struct {
	sm           *stagehand.SceneDirector[*GlobalGameState]
	data         *GlobalGameState
	selectedMenu int
	audioPlayer  *audio.Player
	audioContext *audio.Context
	loopMusic    bool
}

var menus = []string{
	"New Game",
	"Exit",
}
var menusFunc = []func(){
	StartGame,
	Exit,
}

func StartGame() {
	MainMenuInstance.sm.ProcessTrigger(TriggerToOpening)
}
func Exit() {
	os.Exit(0)
}
func (r *MainMenuScene) Update() error {
	if r.loopMusic && !r.audioPlayer.IsPlaying() {
		r.audioPlayer.Rewind()
		r.audioPlayer.Play()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		r.selectedMenu += 1
		if r.selectedMenu == len(menus) {
			r.selectedMenu -= 1
		}
		// r.musicPlayer.QueueSFX(assets.MenuMove)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		r.selectedMenu -= 1
		if r.selectedMenu == -1 {
			r.selectedMenu += 1
		}
		// r.musicPlayer.QueueSFX(assets.MenuMove)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		menusFunc[r.selectedMenu]()
	}
	return nil
}

var MainMenuInstance = &MainMenuScene{}

func (r *MainMenuScene) Draw(screen *ebiten.Image) {
	screen.DrawImage(TitleScreen, &ebiten.DrawImageOptions{})
	// buttonBg := ebiten.NewImage(248, 50)
	// buttonBg.Fill(color.RGBA{R: 0x72, G: 0x72, B: 0x72, A: 0xFF})
	textColor := ebiten.ColorScale{}
	textColor.Scale(0.6, 0.7, 0, 1)
	instrucTextPos := ebiten.GeoM{}
	instrucTextPos.Translate(0, 360)
	text.Draw(screen, "Press 'space' to pick", PixelFontFace, &text.DrawOptions{
		DrawImageOptions: ebiten.DrawImageOptions{
			GeoM: instrucTextPos, ColorScale: textColor,
		},
	})
	for idx, i := range menus {
		pos := ebiten.GeoM{}
		if idx == r.selectedMenu {
			pos.Scale(1.5, 1)
		}
		pos.Translate(0, float64(165+55*idx))

		screen.DrawImage(MenuButtonBg, &ebiten.DrawImageOptions{GeoM: pos})
		pos.Reset()
		pos.Scale(0.6, 0.6)
		pos.Translate(50, float64(165+55*idx))
		pos2 := ebiten.GeoM{}
		pos2.Translate(50, float64(165+55*idx))
		text.Draw(screen, i, PixelFontFace, &text.DrawOptions{

			DrawImageOptions: ebiten.DrawImageOptions{GeoM: pos2, ColorScale: textColor},
		})
	}
}

func (r *MainMenuScene) Load(state *GlobalGameState, manager stagehand.SceneController[*GlobalGameState]) {
	r.sm = manager.(*stagehand.SceneDirector[*GlobalGameState]) // This type assertion is important
	r.data = state
	r.loopMusic = true
	audioContext := audio.CurrentContext()
	if audioContext == nil {
		audioContext = audio.NewContext(48000)
	}
	r.audioContext = audioContext
	gg, err := mp3.DecodeWithoutResampling(bytes.NewReader(CreepyBell))
	if err != nil {
		fmt.Println(err.Error())

	}
	r.audioPlayer, err = r.audioContext.NewPlayer(gg)
	if err != nil {
		fmt.Println(err.Error())

	}
	r.audioPlayer.Play()
	// r.musicPlayer.PlayBgm(CreepyBell, core.TypeMP3)
}
func (s *MainMenuScene) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 600, 400
}
func (s *MainMenuScene) Unload() *GlobalGameState {
	s.loopMusic = false
	s.audioPlayer.Pause()
	s.audioPlayer.Rewind()
	// s.musicPlayer.(*core.DefaultAudioInterface).GetPlayer().Pause()
	// s.musicPlayer.AudioPlayer().Rewind()
	// s.musicPlayer.AudioPlayer().Pause()
	return s.data
}
