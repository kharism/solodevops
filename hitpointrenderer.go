package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kharism/callforgetty/component"
	"github.com/yohamta/donburi"
	ecslib "github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

func HitpointRenderer(ecs *ecslib.ECS, screen *ebiten.Image) {
	query := donburi.NewQuery(filter.Contains(
		component.Sprite,
		component.ScreenPos,
	))
	query.Each(ecs.World, func(e *donburi.Entry) {
		startPosX := component.ScreenPos.Get(e).X
		startPosY := component.ScreenPos.Get(e).Y - 40
		hitpoint := component.HitPoint.Get(e)
		for i := 0; i < hitpoint.MaxHitPoint; i++ {
			transformMatrix := ebiten.GeoM{}
			transformMatrix.Translate(startPosX+float64(35*i), startPosY)
			if i < hitpoint.HitPoint {
				screen.DrawImage(Heart, &ebiten.DrawImageOptions{
					GeoM: transformMatrix,
				})
			} else {
				screen.DrawImage(HeartEmpty, &ebiten.DrawImageOptions{
					GeoM: transformMatrix,
				})
			}

		}
	})
}
