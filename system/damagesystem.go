package system

import (
	"github.com/kharism/callforgetty/component"
	"github.com/yohamta/donburi"
	ecslib "github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

func DamageSystemUpdate(ecs *ecslib.ECS) {
	hurtboxes := []*donburi.Entry{}

	query := donburi.NewQuery(filter.Contains(
		component.HurtBox,
	))
	query.Each(ecs.World, func(e *donburi.Entry) {
		hurtboxes = append(hurtboxes, e)
	})
	if len(hurtboxes) == 0 {
		return
	}
	query2 := donburi.NewQuery(filter.Contains(
		component.HitBox,
	))
	removeIdx := []int{}
	var deadChar *component.HitPointData
	query2.Each(ecs.World, func(e *donburi.Entry) {
		hitbox := component.HitBox.Get(e)
		for remIdx, hurtboxEnt := range hurtboxes {
			if !hurtboxEnt.HasComponent(component.HurtBox) {
				continue
			}
			h := component.HurtBox.Get(hurtboxEnt)
			if (hitbox.X > h.X && h.X+float64(h.Width) > hitbox.X) || (hitbox.X+float64(hitbox.Width) > h.X && hitbox.X+float64(hitbox.Width) < h.X+float64(h.Width)) {
				// register hit
				hp := component.HitPoint.Get(e)
				hp.HitPoint -= 1
				if hp.HitPoint <= 0 {
					deadChar = hp
					return
				}
				removeIdx = append(removeIdx, remIdx)
				// ecs.World.Remove(hurtboxEnt.Entity())
				continue
			}
		}
	})
	if deadChar != nil {
		deadChar.OnDead()
		return
	}
	for _, idx := range removeIdx {
		ecs.World.Remove(hurtboxes[idx].Entity())
	}
}
