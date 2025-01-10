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
	GameData map[string]any
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

const (
	TriggerToCombat stagehand.SceneTransitionTrigger = iota
	TriggerToOpening
	TriggerToGameOver
	TriggerToEnding
	TriggerToMain
)

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
	openingScene := &HanashiScene{scene: OpeningScene(Layout)}
	endingScene := &HanashiScene{scene: EndingScene(Layout)}
	mainmenu := MainMenuInstance
	// mainmenu.musicPlayer = openingScene.scene.AudioInterface
	combatScene := &CombatScene{}
	gameover := &GameOverScene{}
	ruleSet := map[stagehand.Scene[*GlobalGameState]][]stagehand.Directive[*GlobalGameState]{
		mainmenu: {
			stagehand.Directive[*GlobalGameState]{Dest: openingScene, Trigger: TriggerToOpening},
		},
		openingScene: {
			stagehand.Directive[*GlobalGameState]{Dest: combatScene, Trigger: TriggerToCombat},
		},
		combatScene: {
			stagehand.Directive[*GlobalGameState]{Dest: gameover, Trigger: TriggerToGameOver},
			stagehand.Directive[*GlobalGameState]{Dest: endingScene, Trigger: TriggerToEnding},
		},
		gameover: {
			stagehand.Directive[*GlobalGameState]{Dest: mainmenu, Trigger: TriggerToMain},
		},
		endingScene: {
			stagehand.Directive[*GlobalGameState]{Dest: mainmenu, Trigger: TriggerToMain},
		},
	}
	manager := stagehand.NewSceneDirector[*GlobalGameState](mainmenu, state, ruleSet)
	openingScene.scene.Done = func() {
		manager.ProcessTrigger(TriggerToCombat)
	}
	endingScene.scene.Done = func() {
		manager.ProcessTrigger(TriggerToMain)
	}
	if err := ebiten.RunGame(manager); err != nil {
		log.Fatal(err)
	}
}
