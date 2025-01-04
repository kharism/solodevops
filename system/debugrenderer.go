package system

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kharism/callforgetty/component"
	"github.com/yohamta/donburi"
	ecslib "github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

func DebugRenderer(ecs *ecslib.ECS, screen *ebiten.Image) {
	query := donburi.NewQuery(filter.Contains(
		component.HitBox,
	))
	query.Each(ecs.World, func(e *donburi.Entry) {
		hitboxdata := component.HitBox.Get(e)
		border := ebiten.NewImage(hitboxdata.Width, hitboxdata.Height)
		border.Fill(color.RGBA{G: 255, A: 255})
		TransparentFill := border.SubImage(image.Rect(2, 2, hitboxdata.Width-2, hitboxdata.Height-2))

		ebiten.NewImageFromImage(TransparentFill).Fill(color.RGBA{G: 0, A: 0})
		translate := ebiten.GeoM{}
		translate.Translate(hitboxdata.X, hitboxdata.Y)
		screen.DrawImage(border, &ebiten.DrawImageOptions{
			GeoM: translate,
		})
	})
	query2 := donburi.NewQuery(filter.Contains(
		component.HurtBox,
	))
	query2.Each(ecs.World, func(e *donburi.Entry) {
		hurtboxdata := component.HurtBox.Get(e)
		border := ebiten.NewImage(hurtboxdata.Width, hurtboxdata.Height)
		border.Fill(color.RGBA{R: 255, A: 255})
		translate := ebiten.GeoM{}
		translate.Translate(hurtboxdata.X, hurtboxdata.Y)
		screen.DrawImage(border, &ebiten.DrawImageOptions{
			GeoM: translate,
		})
	})
}
