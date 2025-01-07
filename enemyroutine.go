package main

import (
	"time"

	"github.com/kharism/callforgetty/component"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	ecslib "github.com/yohamta/donburi/ecs"
)

var ALREADY_FIRED = "ALREADY_FIRED"
var WARM_UP = "WARM_UP"
var CURRENT_STRATEGY = "CUR_STRATEGY"

func monsterRoutine(ecs *ecslib.ECS, self *donburi.Entry) {
	memory := component.EnemyRoutine.Get(self).Memory
	if memory[CURRENT_STRATEGY] == "" {
		memory[CURRENT_STRATEGY] = "MOVE"
		memory[WARM_UP] = time.Now().Add(300 * time.Millisecond)
	}
	scrPos := component.ScreenPos.Get(self)
	hitBox := component.HitBox.Get(self)
	player, _ := component.PlayerTag.First(ecs.World)
	hurtboxXstart := scrPos.X + 20
	if memory[CURRENT_STRATEGY] == "MOVE" {
		if waitTime, ok := memory[WARM_UP].(time.Time); ok && waitTime.Before(time.Now()) {
			scrPos.X -= 1
			hitBox.X -= 1
			playerHitbox := component.HitBox.Get(player)
			if playerHitbox.X+float64(playerHitbox.Width) > hurtboxXstart {
				memory[CURRENT_STRATEGY] = "WARMUP"
			} else {
				memory[CURRENT_STRATEGY] = "MOVE"
			}
			memory[WARM_UP] = time.Now().Add(20 * time.Millisecond)
		}

	}
	if memory[CURRENT_STRATEGY] == "WARMUP" {
		if waitTime, ok := memory[WARM_UP].(time.Time); ok && waitTime.Before(time.Now()) {
			component.Sprite.Get(self).Image = MonsterSprite3
			memory[CURRENT_STRATEGY] = "ATK"
			memory[WARM_UP] = time.Now().Add(100 * time.Millisecond)
		}
	}
	if memory[CURRENT_STRATEGY] == "ATK" {
		if waitTime, ok := memory[WARM_UP].(time.Time); ok && waitTime.Before(time.Now()) {
			memory[CURRENT_STRATEGY] = "CHECK_1"
			component.Sprite.Get(self).Image = MonsterSprite2
			hurtBox1 := ecs.World.Create(component.HurtBox)
			entry := ecs.World.Entry(hurtBox1)
			hurtboxData := component.HurtBox.Get(entry)
			hurtboxData.X = scrPos.X + 20
			hurtboxData.Y = hitBox.Y
			hurtboxData.Height = hitBox.Height
			hurtboxData.Width = int(hitBox.X) - int(hurtboxData.X)
			memory[WARM_UP] = time.Now().Add(200 * time.Millisecond)
			component.EventQueue.AddEvent(&RemoveEntryEvent{time: time.Now().Add(50 * time.Millisecond), entry: entry})
		}
	}
	if memory[CURRENT_STRATEGY] == "CHECK_1" {
		if waitTime, ok := memory[WARM_UP].(time.Time); ok && waitTime.Before(time.Now()) {
			playerHitbox := component.HitBox.Get(player)
			if playerHitbox.X+float64(playerHitbox.Width) > hurtboxXstart {
				memory[CURRENT_STRATEGY] = "WARMUP"
			} else {
				memory[CURRENT_STRATEGY] = "MOVE"
			}
		}
	}
}

type RemoveEntryEvent struct {
	entry *donburi.Entry
	time  time.Time
}

func (e *RemoveEntryEvent) Execute(ecs *ecs.ECS) {
	ecs.World.Remove(e.entry.Entity())
}
func (e *RemoveEntryEvent) GetTime() time.Time {
	return e.time
}
