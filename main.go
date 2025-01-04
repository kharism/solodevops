package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/joelschutz/stagehand"
	"github.com/kharism/hanashi/core"
)

const (
	screenWidth  = 600
	screenHeight = 400
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

type GlobalGameState struct {
}

type LayouterImpl struct {
}

var Layout *LayouterImpl

func (l *LayouterImpl) GetLayout() (width, height int) {
	return 600, 400
}
func (l *LayouterImpl) GetNamePosition() (x, y int) {
	return 0, 400 - 50
}
func (l *LayouterImpl) GetTextPosition() (x, y int) {
	return 0, 400 - 30
}
func main() {
	flag.Parse()
	core.DetectKeyboardNext = func() bool {
		return inpututil.IsKeyJustReleased(ebiten.KeySpace)
	}
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("sdadadad")
	// Level := scene.GenerateLayout1()
	Layout = &LayouterImpl{}

	state := &GlobalGameState{}
	// combatScene := &scene.CombatScene{}
	// rewardScene := &scene.RewardScene{}
	// openingScene := &HanashiScene{scene: OpeningScene(Layout)}
	combatScene := &CombatScene{}
	ruleSet := map[stagehand.Scene[*GlobalGameState]][]stagehand.Directive[*GlobalGameState]{}
	manager := stagehand.NewSceneDirector[*GlobalGameState](combatScene, state, ruleSet)
	if err := ebiten.RunGame(manager); err != nil {
		log.Fatal(err)
	}
}
