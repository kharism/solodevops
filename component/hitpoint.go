package component

import "github.com/yohamta/donburi"

type HitPointData struct {
	HitPoint    int
	MaxHitPoint int
}

var HitPoint = donburi.NewComponentType[HitPointData]()
