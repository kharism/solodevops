package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/joelschutz/stagehand"
	"github.com/kharism/callforgetty/component"
	"github.com/kharism/callforgetty/system"
	"github.com/yohamta/donburi"
	ecslib "github.com/yohamta/donburi/ecs"
)

type CombatScene struct {
	State    *GlobalGameState
	director *stagehand.SceneDirector[*GlobalGameState]
	Ecs      *ecslib.ECS
}

func (c *CombatScene) Update() error {
	c.Ecs.Update()
	return nil
}
func (c *CombatScene) Draw(screen *ebiten.Image) {
	screen.Clear()
	c.Ecs.DrawLayer(LayerDebug, screen)
	c.Ecs.DrawLayer(LayerCharacter, screen)
	c.Ecs.DrawLayer(LayerHP, screen)

}
func (s *CombatScene) Load(state *GlobalGameState, director stagehand.SceneController[*GlobalGameState]) {
	s.director = director.(*stagehand.SceneDirector[*GlobalGameState]) // This type assertion is important
	// s.scene.EventIndex = 0
	world := donburi.NewWorld()
	s.Ecs = ecslib.NewECS(world)
	entry := LoadPlayer(s.Ecs)
	LoadEnemy(s.Ecs)
	playerMovement := system.PlayerMoveSystem{PlayerIndex: entry}
	playerAttackSystem := system.PlayerAttackSystem{
		PlayerIndex:        entry,
		PlayerActiveSprite: OmarSprite2,
		PlayerOriSprite:    OmarSprite1,
	}
	eventQueueSystem := system.EventQueueSystem{}
	s.Ecs.AddSystem(playerMovement.Update).
		AddSystem(playerAttackSystem.Update).
		AddSystem(eventQueueSystem.Update).
		AddRenderer(LayerCharacter, system.Spriterenderer).
		AddRenderer(LayerHP, HitpointRenderer).
		AddRenderer(LayerDebug, system.DebugRenderer)
}

var startY = 300.0
var startXPlayer = 30.0
var startXMonster = 400.0

func LoadPlayer(ecs *ecslib.ECS) *donburi.Entry {
	entity := ecs.World.Create(component.ScreenPos, component.Sprite, component.Scale, component.HitPoint, component.HitBox)
	entry := ecs.World.Entry(entity)
	component.Sprite.Get(entry).Image = OmarSprite1
	scrPos := component.ScreenPos.Get(entry)
	scrPos.X = startXPlayer
	scrPos.Y = startY
	scale := component.Scale.Get(entry)
	scale.ScaleX = 2.0
	scale.ScaleY = 2.0
	hp := component.HitPoint.Get(entry)
	hp.HitPoint = 2
	hp.MaxHitPoint = 2
	component.HitBox.Set(entry, &component.HitBoxData{
		X:      startXPlayer,
		Y:      startY,
		Width:  64*2 - 50,
		Height: 64 * 2,
	})
	return entry
}
func LoadEnemy(ecs *ecslib.ECS) *donburi.Entry {
	entity := ecs.World.Create(component.ScreenPos, component.Sprite, component.Scale, component.HitPoint, component.HitBox)
	entry := ecs.World.Entry(entity)
	component.Sprite.Get(entry).Image = MonsterSprite1
	scrPos := component.ScreenPos.Get(entry)
	scrPos.X = startXMonster
	scrPos.Y = startY
	scale := component.Scale.Get(entry)
	scale.ScaleX = 2.0
	scale.ScaleY = 2.0
	hp := component.HitPoint.Get(entry)
	hp.HitPoint = 3
	hp.MaxHitPoint = 6
	component.HitBox.Set(entry, &component.HitBoxData{
		X:      startXMonster + 32,
		Y:      startY,
		Width:  64*2 - 30,
		Height: 64 * 2,
	})
	return entry
}
func (s *CombatScene) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}
func (s *CombatScene) Unload() *GlobalGameState {
	// your unload code
	// s.scene.Events[0].Execute(s.scene)

	return s.State
}
