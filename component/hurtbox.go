package component

import "github.com/yohamta/donburi"

type HurtBoxData struct {
	X, Y          float64
	Width, Height int
}

// ScreenPos define position of object on the screen
var HurtBox = donburi.NewComponentType[HurtBoxData]()
