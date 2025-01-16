package main

import (
	"fmt"
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kharism/hanashi/core"
)

func EndingScene(layouter core.GetLayouter) *core.Scene {
	scene := core.NewScene()
	scene.SetLayouter(layouter)
	scene.FontFace = PixelFontFace
	scene.Characters = []*core.Character{
		core.NewCharacterImage("Omar", Omarov),
		core.NewCharacterImage("Smith", ManInBlack),
		// core.NewCharacterImage("Monster", Monster),
	}
	scene.Events = []core.Event{
		&core.ComplexEvent{
			Events: []core.Event{
				&core.PlayBgmEvent{
					Audio: &Unease,
					Type:  core.TypeMP3,
				},
				core.NewBgChangeEvent(Lobby, core.MoveParam{Sx: 0, Sy: 0}, nil),
				core.NewCharacterAddEvent("Omar", core.MoveParam{Sy: 120, Sx: 600, Tx: 460, Ty: 120, Speed: 10},
					&core.ScaleParam{Sx: 4, Sy: 4}),
				core.NewDialogueEvent("Omar", "*pant* *pant*\n(It died?)", PixelFontFace),
			},
		},
		core.NewDialogueEvent("Omar", "(now get to the server room\nand restart the server?)", PixelFontFace),
		&core.ComplexEvent{
			Events: []core.Event{
				core.NewBgChangeEvent(DataCenter, core.MoveParam{Sx: 0, Sy: 0}, nil),
				core.NewDialogueEvent("Omar", "(this should restart the server)", PixelFontFace),
			},
		},
		core.NewDialogueEvent("Omar", "(just need to wait for police\nto come and clean up\nthe mess)", PixelFontFace),
		&core.ComplexEvent{
			Events: []core.Event{
				core.NewCharacterAddEvent("Smith",
					core.MoveParam{Sy: 120, Sx: -300, Tx: 0, Ty: 120, Speed: 10},
					&core.ScaleParam{Sx: 4, Sy: 4}),
				core.NewDialogueEvent("", "", PixelFontFace),
			},
		},
		core.NewDialogueEvent("Smith", "I see the problem has been fixed", PixelFontFace),
		core.NewDialogueEvent("Omar", "Umm Who are you?", PixelFontFace),
		core.NewDialogueEvent("Smith", "(Show ID)\nJames Smith, Agent of\nNational Burreau of Wildlife", PixelFontFace),
		core.NewDialogueEvent("Smith", "We got your distress call about\nmonster killing someone here", PixelFontFace),
		core.NewDialogueEvent("Omar", "Yea, as you said. Its been dealt\nwith", PixelFontFace),
		core.NewDialogueEvent("Omar", "Anyway can you tell me what that\nthing is?", PixelFontFace),
		core.NewDialogueEvent("Smith", "Sure, please sit. I'll tell\nyou everything while\nwe tend your wound", PixelFontFace),
		core.NewDialogueEvent("", "*medics checks omar's wound*", PixelFontFace),
		core.NewDialogueEvent("Smith", "The creature you just killed, is...", PixelFontFace),
		core.NewDialogueEvent("medic", "(inject Omar with sedative)", PixelFontFace),
		core.NewDialogueEvent("Omar", "Nggggh...", PixelFontFace),
		core.NewDialogueEvent("Smith", "....Just your nightmare", PixelFontFace),
		core.NewDialogueEvent("Smith", "Now sleep", PixelFontFace),
		&core.ComplexEvent{
			Events: []core.Event{
				&core.StopBgmEvent{},
				core.NewCharacterRemoveEvent("Smith"),
				core.NewCharacterRemoveEvent("Omar"),
				core.NewDialogueEvent("", "", PixelFontFace),
				core.NewBgChangeEvent(End, core.MoveParam{Sx: 0, Sy: 0}, nil),
			},
		},
	}
	scene.TxtBg = ebiten.NewImage(1000-128, 128)
	scene.TxtBg.Fill(color.RGBA{R: 0x4f, G: 0x8f, B: 0xba, A: 255})
	pp, err := core.NewDefaultAudioInterfacer()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	scene.AudioInterface = pp
	return scene
}
