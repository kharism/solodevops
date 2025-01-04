package system

import (
	"github.com/kharism/callforgetty/component"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	ecslib "github.com/yohamta/donburi/ecs"
)

type PlayerMoveSystem struct {
	PlayerIndex *donburi.Entry
}

func (p *PlayerMoveSystem) Update(ecs *ecslib.ECS) {
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		component.ScreenPos.Get(p.PlayerIndex).X += 3
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		component.ScreenPos.Get(p.PlayerIndex).X -= 3
	}
}
