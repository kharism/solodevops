package main

import (
	"fmt"
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kharism/hanashi/core"
)

func OpeningScene(layouter core.GetLayouter) *core.Scene {
	scene := core.NewScene()
	scene.SetLayouter(layouter)
	scene.FontFace = PixelFontFace
	scene.Characters = []*core.Character{
		core.NewCharacterImage("Omar", Omarov),
	}
	// spiaShader, _ := core.GetShaderPool().GetShader(core.SEPIA_SHADER)
	scene.Events = []core.Event{
		&core.ComplexEvent{
			Events: []core.Event{
				&core.PlayBgmEvent{
					Audio: &BassMp3,
					Type:  core.TypeMP3,
				},
				core.NewBgChangeEvent(WarehouseBg, core.MoveParam{Sx: 0, Sy: 0}, nil),
				core.NewCharacterAddEvent("Omar",
					&core.MoveParam{Sy: 120, Sx: 600, Tx: 460, Ty: 120, Speed: 10},
					&core.ScaleParam{Sx: 4, Sy: 4},
				),
				core.NewDialogueEvent("Omar", "(*Sneeze*..what a weather)", PixelFontFace),
			},
		},
		core.NewDialogueEvent("Omar", "(it is colder than I expect)", PixelFontFace),
		core.NewDialogueEvent("Omar", "(Who even suspect 'Remote\nsoftware developer' job meaning\nsoftware developer works in rural\noffice out of nowhere)", PixelFontFace),
		core.NewDialogueEvent("Omar", "(but the pay's good so no\ncomplaint there)", PixelFontFace),
		&core.ComplexEvent{
			Events: []core.Event{
				&core.PlaySfxEvent{Audio: &PhoneRing, Type: core.TypeMP3},
				core.NewDialogueEvent("", "", PixelFontFace),
			},
		},
		core.NewDialogueEvent("Omar", "Yes father?", PixelFontFace),
		core.NewDialogueEvent("Omar", "No, I won't be home at the\nend of the year. I need the\nextra money for my wedding.", PixelFontFace),
		core.NewDialogueEvent("Omar", "I'll do my best to go home\nnext Eid.", PixelFontFace),
		core.NewDialogueEvent("Omar", "No. There's no cryptid\nmurder here. You should stop\nwatching conspiracy videos.", PixelFontFace),
		core.NewDialogueEvent("Omar", "bye dad.", PixelFontFace),
		&core.ComplexEvent{
			Events: []core.Event{
				core.NewBgChangeEvent(Lobby, core.MoveParam{Sx: 0, Sy: 0}, nil),
				core.NewDialogueEvent("Security", "Still working on holiday omar?", PixelFontFace),
			},
		},

		core.NewDialogueEvent("Omar", "Yes, Everyone left on christmas\neve ", PixelFontFace),
		core.NewDialogueEvent("Omar", "but there's still some\njob that needs to be done.", PixelFontFace),
		core.NewDialogueEvent("Omar", "And the someone needs to\nbe close at the server.", PixelFontFace),
		core.NewDialogueEvent("Security", "Tom the new guy will\nstay with you.", PixelFontFace),
		core.NewDialogueEvent("Security", "There might be some\nserial killer on the loose lately.", PixelFontFace),
		core.NewDialogueEvent("Omar", "Mmm Okay, thanks.", PixelFontFace),
		core.NewDialogueEvent("Omar", "(Another murder?)", PixelFontFace),
		core.NewDialogueEvent("Omar", "(Am I the one out of\ntouch with the news???)", PixelFontFace),
		/*&core.ComplexEvent{
			Events: []core.Event{
				core.NewBgChangeEvent(OfficeDay, core.MoveParam{Sx: 0, Sy: 0}, &core.ShaderParam{Shader: spiaShader}),
				&core.CharacterRemoveEvent{Name: "Omar"},
				core.NewDialogueEvent("", "Few days ago....", PixelFontFace),
			},
		},
		&core.ComplexEvent{
			Events: []core.Event{
				core.NewCharacterAddEvent("Omar",
					&core.MoveParam{Sy: 120, Sx: 600, Tx: 460, Ty: 120, Speed: 10},
					&core.ScaleParam{Sx: 4, Sy: 4},
				),
				core.NewCharacterAddShaderEvent("Omar", &core.ShaderParam{Shader: spiaShader}),
				core.NewDialogueEvent("Boss", "we need to clear up these high\npriority ticket before new\nyear's eve,", PixelFontFace),
			},
		},
		core.NewDialogueEvent("Boss", "and someone needs to stand by\nnear the server for the\nremote worker.", PixelFontFace),
		core.NewDialogueEvent("Omar", "You have nothing to do on\nchristmas eve don't you omar?", PixelFontFace),
		core.NewDialogueEvent("Boss", "You live alone, no girlfriend,\nand you're a muslim.", PixelFontFace),
		core.NewDialogueEvent("Omar", "Just because I'm a muslim does\nnot mean I can work alone\non christmas eve.", PixelFontFace),
		core.NewDialogueEvent("Boss", "do it or not. There will be\n1.5 times overtime bonus", PixelFontFace),
		core.NewDialogueEvent("Omar", "*smile* Sweet, I'll take it", PixelFontFace),
		&core.ComplexEvent{
			Events: []core.Event{
				core.NewBgChangeEvent(OfficeNight, core.MoveParam{Sx: 0, Sy: 0}, nil),
				&core.CharacterRemoveEvent{Name: "Omar"},
				core.NewCharacterAddShaderEvent("Omar", nil),
				core.NewDialogueEvent("", "Current day....", PixelFontFace),
			},
		},
		&core.ComplexEvent{
			Events: []core.Event{
				&core.PlayBgmEvent{
					Audio: &SnowStorm,
					Type:  core.TypeMP3,
				},
				core.NewCharacterAddEvent("Omar",
					&core.MoveParam{Sy: 120, Sx: 600, Tx: 460, Ty: 120, Speed: 10},
					&core.ScaleParam{Sx: 4, Sy: 4},
				),
				core.NewCharacterAddShaderEvent("Omar", &core.ShaderParam{Shader: nil}),
				core.NewDialogueEvent("Omar", "(Alright, this time for some\nticket clean up and some\ntroublehooting. This shouldn't\ntake long)", PixelFontFace),
			},
		},
		core.NewDialogueEvent("", "*Few Hours Later....*", PixelFontFace),
		core.NewDialogueEvent("Omar", "(*groan* FINALLY,Troubleshooting\nis easy, as long as the trouble\ndoes not shoot back)", PixelFontFace),
		core.NewDialogueEvent("Omar", "(WHICH IS CLEARLY NOT THE CASE\nHERE)", PixelFontFace),
		core.NewDialogueEvent("Omar", "(now create a merge request and\nthat should be it for tonight)", PixelFontFace),
		core.NewDialogueEvent("Omar", "(The snowstorm has just begin.\nTime for night snack\n*pull out cup noodle from drawer*)", PixelFontFace),
		core.NewDialogueEvent("Omar", "(Hmm, Never try this brand before\n.\"Mie Goreng Extra spicy\")", PixelFontFace),
		*/
		core.NewDialogueEvent("Omar", "(Hmm, Never try this brand before\n.\"Mie Goreng Extra spicy\")", PixelFontFace),
		core.NewDialogueEvent("Omar", "(How spicy does it mean?\nIndonesian spicy or\nmalaysian spicy?)", PixelFontFace),
		core.NewDialogueEvent("Omar", "(should I mix the\nchili powder?)", PixelFontFace),
		core.NewOptionSelectEvent("Spice", "put chilly in?", "yes", "no"),
		&core.PlaySfxEvent{
			Audio: &Pouring,
			Type:  core.TypeMP3,
		},
		core.NewDialogueEvent("Omar", "(now we should wait, I should get\nsome coffee and maybe talk with Tom)", PixelFontFace),
		&core.ComplexEvent{
			Events: []core.Event{
				&core.PlayBgmEvent{
					Audio: &Unease,
					Type:  core.TypeMP3,
				},
				core.NewBgChangeEvent(Lobby, core.MoveParam{Sx: 0, Sy: 0}, nil),
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
	scene.Events[0].Execute(scene)
	return scene
}
