package component

import "github.com/yohamta/donburi"

type HitPointData struct {
	HitPoint    int
	MaxHitPoint int
	OnDead      OnDeadFunction
}
type OnDeadFunction func()

var HitPoint = donburi.NewComponentType[HitPointData]()
