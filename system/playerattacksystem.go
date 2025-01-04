package system

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/kharism/callforgetty/component"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	ecslib "github.com/yohamta/donburi/ecs"
)

type PlayerAttackSystem struct {
	PlayerIndex        *donburi.Entry
	PlayerActiveSprite *ebiten.Image
	PlayerOriSprite    *ebiten.Image
	finishCooldown     time.Time
}

func (p *PlayerAttackSystem) Update(ecs *ecslib.ECS) {
	if time.Now().Before(p.finishCooldown) {
		return
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyQ) {
		// component.ScreenPos.Get(p.PlayerIndex).X += 3
		component.Sprite.Get(p.PlayerIndex).Image = p.PlayerActiveSprite
		now := time.Now()
		cooldownTime := now.Add(200 * time.Millisecond)
		component.EventQueue.AddEvent(&ChangeSpriteEvent{
			playerEntry: p.PlayerIndex,
			newSprite:   p.PlayerOriSprite,
			time:        cooldownTime,
		})
		p.finishCooldown = cooldownTime

	}
}

type ChangeSpriteEvent struct {
	playerEntry *donburi.Entry
	newSprite   *ebiten.Image
	time        time.Time
}

func (s *ChangeSpriteEvent) Execute(ecs *ecs.ECS) {
	component.Sprite.Get(s.playerEntry).Image = s.newSprite
}
func (s *ChangeSpriteEvent) GetTime() time.Time {
	return s.time
}
