package component

import "github.com/yohamta/donburi"

type ScaleComponentData struct {
	ScaleX, ScaleY float64
}

// ScreenPos define position of object on the screen
var Scale = donburi.NewComponentType[ScaleComponentData]()
