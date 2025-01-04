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
		playerPos := component.ScreenPos.Get(p.PlayerIndex)
		playerHitbox := component.HitBox.Get(p.PlayerIndex)
		cooldownTime := now.Add(200 * time.Millisecond)
		hurtBoxentity := ecs.World.Create(component.HurtBox)
		hurtBoxentry := ecs.World.Entry(hurtBoxentity)
		hurtbox := component.HurtBox.Get(hurtBoxentry)
		hurtbox.X = playerPos.X + float64(playerHitbox.Width)
		hurtbox.Y = playerPos.Y
		hurtbox.Height = playerHitbox.Height
		hurtbox.Width = 50
		component.EventQueue.AddEvent(&ChangeSpriteEvent{
			playerEntry:  p.PlayerIndex,
			newSprite:    p.PlayerOriSprite,
			time:         cooldownTime,
			hurtboxEntry: hurtBoxentry,
		})
		p.finishCooldown = cooldownTime

	}
}

type ChangeSpriteEvent struct {
	playerEntry  *donburi.Entry
	newSprite    *ebiten.Image
	time         time.Time
	hurtboxEntry *donburi.Entry
}

func (s *ChangeSpriteEvent) Execute(ecs *ecs.ECS) {
	component.Sprite.Get(s.playerEntry).Image = s.newSprite
	ecs.World.Remove(s.hurtboxEntry.Entity())
}
func (s *ChangeSpriteEvent) GetTime() time.Time {
	return s.time
}
