package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kharism/callforgetty/component"
	"github.com/yohamta/donburi"
	ecslib "github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

func Spriterenderer(ecs *ecslib.ECS, screen *ebiten.Image) {
	query := donburi.NewQuery(filter.Contains(
		component.Sprite,
		component.ScreenPos,
	))
	query.Each(ecs.World, func(e *donburi.Entry) {
		sprite := component.Sprite.Get(e)
		transformMatrix := ebiten.GeoM{}
		scale := component.Scale.Get(e)
		if scale.ScaleX != 0 && scale.ScaleY != 0 {
			transformMatrix.Scale(scale.ScaleX, scale.ScaleY)
		} else {
			transformMatrix.Scale(1, 1)
		}
		scrPos := component.ScreenPos.Get(e)
		transformMatrix.Translate(scrPos.X, scrPos.Y)
		screen.DrawImage(sprite.Image, &ebiten.DrawImageOptions{
			GeoM: transformMatrix,
		})
	})
}
