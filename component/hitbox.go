package component

import "github.com/yohamta/donburi"

type HitBoxData struct {
	X, Y          float64
	Width, Height int
}

// ScreenPos define position of object on the screen
var HitBox = donburi.NewComponentType[HitBoxData]()
