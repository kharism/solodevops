package main

import (
	"fmt"
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kharism/hanashi/core"
)

type dummyScene struct{}

func (d *dummyScene) Execute(scene *core.Scene) {

}
func OpeningScene(layouter core.GetLayouter) *core.Scene {
	scene := core.NewScene()
	scene.SetLayouter(layouter)
	scene.FontFace = PixelFontFace
	scene.Characters = []*core.Character{
		core.NewCharacterImage("Omar", Omarov),
		core.NewCharacterImage("Monster", Monster),
	}
	spiaShader, _ := core.GetShaderPool().GetShader(core.SEPIA_SHADER)
	scene.Events = []core.Event{
		// &dummyScene{},
		&core.ComplexEvent{
			Events: []core.Event{
				&core.PlayBgmEvent{
					Audio: &BassMp3,
					Type:  core.TypeMP3,
				},
				core.NewBgChangeEvent(WarehouseBg, core.MoveParam{Sx: 0, Sy: 0}, nil),
				core.NewCharacterAddEvent("Omar",
					core.MoveParam{Sy: 120, Sx: 600, Tx: 460, Ty: 120, Speed: 10},
					&core.ScaleParam{Sx: 4, Sy: 4},
				),
				core.NewDialogueEvent("Omar", "(*Sneeze*..what a weather)", PixelFontFace),
			},
		},
		&core.ComplexEvent{
			Events: []core.Event{

				core.NewDialogueEvent("Omar", "(it is colder than I expect)", PixelFontFace),
			},
		},
		core.NewDialogueEvent("Omar", "(Who even suspect 'Remote\nsoftware developer' job meaning\nsoftware developer works in rural\noffice out of nowhere)", PixelFontFace),
		core.NewDialogueEvent("Omar", "(but the pay's good so no\ncomplaint there)", PixelFontFace),
		&core.ComplexEvent{
			Events: []core.Event{
				&core.StopBgmEvent{},
				&core.PlaySfxEvent{Audio: &PhoneRing, Type: core.TypeMP3, Name: "ring1"},
				core.NewDialogueEvent("", "", PixelFontFace),
			},
		},
		&core.ComplexEvent{
			Events: []core.Event{
				&core.StopSfxEvent{Name: "ring1"},
				core.NewDialogueEvent("Omar", "Yes father?", PixelFontFace),
			},
		},

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
		&core.ComplexEvent{
			Events: []core.Event{
				core.NewBgChangeEvent(OfficeDay, core.MoveParam{Sx: 0, Sy: 0}, &core.ShaderParam{Shader: spiaShader}),
				&core.CharacterRemoveEvent{Name: "Omar"},
				core.NewDialogueEvent("", "Few days ago....", PixelFontFace),
			},
		},
		&core.ComplexEvent{
			Events: []core.Event{
				core.NewCharacterAddEvent("Omar",
					core.MoveParam{Sy: 120, Sx: 600, Tx: 460, Ty: 120, Speed: 10},
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
					core.MoveParam{Sy: 120, Sx: 600, Tx: 460, Ty: 120, Speed: 10},
					&core.ScaleParam{Sx: 4, Sy: 4},
				),
				core.NewCharacterAddShaderEvent("Omar", &core.ShaderParam{Shader: nil}),
				core.NewDialogueEvent("Omar", "(Alright, this time for some\nticket clean up and some\ntroublehooting. This shouldn't\ntake long)", PixelFontFace),
			},
		},
		core.NewDialogueEvent("", "*Few Hours Later....*", PixelFontFace),
		core.NewDialogueEvent("Omar", "(WHAT'S WRONG WITH THE BUG,\nIts like a game of whack-a-mole)", PixelFontFace),
		core.NewDialogueEvent("Omar", "(*smash face into table*)", PixelFontFace),
		core.NewDialogueEvent("Tom", "Everything alright sir?", PixelFontFace),
		core.NewDialogueEvent("Omar", "NO!!", PixelFontFace),
		core.NewDialogueEvent("Tom", "Business as usual I see", PixelFontFace),
		core.NewDialogueEvent("Tom", "*leaves*", PixelFontFace),
		core.NewDialogueEvent("", "*Few Hours Later....*", PixelFontFace),
		core.NewDialogueEvent("Omar", "(*groan* FINALLY,Troubleshooting\nis easy, as long as the trouble\ndoes not shoot back)", PixelFontFace),
		core.NewDialogueEvent("Omar", "(WHICH IS CLEARLY NOT THE CASE\nHERE)", PixelFontFace),
		core.NewDialogueEvent("Omar", "(I don't know how to test\nsome functions properly.\nSo this will do)", PixelFontFace),
		core.NewDialogueEvent("Omar", "(Worst case scenario it will\ncrash the server and I\nneed to restart it physically)", PixelFontFace),
		core.NewDialogueEvent("Omar", "(The snowstorm has just begin.\nTime for night snack\n*pull out cup noodle from drawer*)", PixelFontFace),
		core.NewDialogueEvent("Omar", "(Hmm, Never try this brand before\n.\"Mie Goreng Extra spicy\")", PixelFontFace),
		core.NewDialogueEvent("Omar", "(How spicy does it mean?\nIndonesian spicy or\nmalaysian spicy?)", PixelFontFace),
		core.NewDialogueEvent("Omar", "(should I mix the\nchili powder?)", PixelFontFace),
		core.NewOptionSelectEvent("Spice", "put chilly in?", "yes", "no"),
		&core.ComplexEvent{
			Events: []core.Event{
				&core.PlaySfxEvent{
					Audio: &Pouring,
					Type:  core.TypeMP3,
				},
				core.NewDialogueEvent("", "", PixelFontFace),
			},
		},
		&core.ComplexEvent{
			Events: []core.Event{
				&core.StopBgmEvent{},
				core.NewDialogueEvent("Omar", "(now we should wait, I should get\nsome coffee and maybe apoligize\nto Tom)", PixelFontFace),
			},
		},

		&core.ComplexEvent{
			Events: []core.Event{
				&core.PlayBgmEvent{
					Audio: &Unease,
					Type:  core.TypeMP3,
				},
				core.NewBgChangeEvent(Lobby, core.MoveParam{Sx: 0, Sy: 0}, nil),
				core.NewDialogueEvent("Omar", "(What's this smell?)", PixelFontFace),
			},
		},
		core.NewDialogueEvent("Omar", "(it smells like blood)", PixelFontFace),
		&core.ComplexEvent{
			Events: []core.Event{
				core.NewBgChangeEvent(SeveredHand, core.MoveParam{Sx: 0, Sy: 0}, nil),
				core.NewDialogueEvent("Omar", "(*gasp* whose hand is that)", PixelFontFace),
			},
		},
		&core.ComplexEvent{
			Events: []core.Event{
				core.NewCharacterAddEvent("Monster",
					core.MoveParam{Sy: 120, Sx: -300, Tx: 0, Ty: 120, Speed: 10},
					&core.ScaleParam{Sx: 4, Sy: 4},
				),
				core.NewDialogueEvent("", "*the monster gnaw at tom's body*", PixelFontFace),
			},
		},
		core.NewDialogueEvent("Omar", "(I need to back to my room)", PixelFontFace),
		core.NewDialogueEvent("Omar", "(hope it does not follow me)", PixelFontFace),
		core.NewDialogueEvent("Omar", "*walk back slowly*", PixelFontFace),
		core.NewDialogueEvent("", "*the monster begin chasing*", PixelFontFace),
		core.NewDialogueEvent("Omar", "*run to room*", PixelFontFace),
		&core.ComplexEvent{
			Events: []core.Event{
				&core.PlaySfxEvent{
					Audio: &DoorSlam,
					Type:  core.TypeMP3,
				},
				core.NewCharacterRemoveEvent("Monster"),
				core.NewBgChangeEvent(OfficeNight, core.MoveParam{Sx: 0, Sy: 0}, nil),
			},
		},
		&core.ComplexEvent{
			Events: []core.Event{
				&core.PlaySfxEvent{
					Audio: &DoorLock,
					Type:  core.TypeMP3,
				},
				core.NewDialogueEvent("Omar", "(This should be it)", PixelFontFace),
			},
		},
		core.NewDialogueEvent("Omar", "(Need to call police)", PixelFontFace),
		&core.ComplexEvent{
			Events: []core.Event{
				&core.PlaySfxEvent{
					Audio: &PhoneDial,
					Type:  core.TypeMP3,
				},
				core.NewDialogueEvent("", "", PixelFontFace),
			},
		},
		core.NewDialogueEvent("Omar", "Hello, there's a murder in\nboostlab", PixelFontFace),
		core.NewDialogueEvent("Omar", "And there's monster here,\nplease hurry", PixelFontFace),
		&core.ComplexEvent{
			Events: []core.Event{
				&core.PlaySfxEvent{
					Audio: &DoorBang,
					Type:  core.TypeMP3,
				},
				core.NewDialogueEvent("Omar", "(The door won't last)", PixelFontFace),
			},
		},
		core.NewDialogueEvent("Omar", "(I need to jump down)", PixelFontFace),
		&core.ComplexEvent{
			Events: []core.Event{
				&core.PlaySfxEvent{
					Audio: &Thud,
					Type:  core.TypeMP3,
				},
				core.NewBgChangeEvent(WarehouseBg, core.MoveParam{Sx: 0, Sy: 0}, nil),
				core.NewDialogueEvent("Omar", "(the thick snow helped\nwith my fall)", PixelFontFace),
			},
		},
		core.NewDialogueEvent("Omar", "(doubt it will help with me driving\nthough)", PixelFontFace),
		core.NewDialogueEvent("Omar", "(I need some kind of weapon\nto defend myself)", PixelFontFace),
		core.NewDialogueEvent("Omar", "(this crowbar will do, need to\nchannel my inner gordon freeman)", PixelFontFace),
		&core.ComplexEvent{
			Events: []core.Event{
				&core.PlaySfxEvent{Audio: &PhoneRing, Type: core.TypeMP3},
				core.NewDialogueEvent("", "", PixelFontFace),
			},
		},
		core.NewDialogueEvent("Omar", "Yes boss?", PixelFontFace),
		core.NewDialogueEvent("Boss", "The CI/CD server has crashed\nCan you restart them?", PixelFontFace),
		core.NewDialogueEvent("Omar", "There's a monster in the office\nand It killed Tom", PixelFontFace),
		core.NewDialogueEvent("Boss", "Oh...Can you still restart the\nserver?", PixelFontFace),
		core.NewDialogueEvent("Omar", "There's man-eating monster in the\noffice", PixelFontFace),
		core.NewDialogueEvent("Boss", "The server can't restart itself", PixelFontFace),
		core.NewDialogueEvent("Boss", "And make sure that\nmonster does not reach\nserver room", PixelFontFace),
		core.NewDialogueEvent("Boss", "Or else, you don't get your bonus", PixelFontFace),
		core.NewDialogueEvent("", "*hangs up*", PixelFontFace),
		core.NewDialogueEvent("Omar", "Holy-C, what a boss", PixelFontFace),
		core.NewDialogueEvent("Omar", "(Need to get going\nor my bonus goes bust)", PixelFontFace),
		&core.ComplexEvent{
			Events: []core.Event{
				core.NewBgChangeEvent(Lobby, core.MoveParam{Sx: 0, Sy: 0}, nil),
				core.NewDialogueEvent("Omar", "(Just be carefull not to\nget caught by that....)", PixelFontFace),
			},
		},
		&core.ComplexEvent{
			Events: []core.Event{
				core.NewCharacterAddEvent("Monster",
					core.MoveParam{Sy: 120, Sx: -300, Tx: 0, Ty: 120, Speed: 10},
					&core.ScaleParam{Sx: 4, Sy: 4}),
				core.NewDialogueEvent("Omar", "MOTHERFORKER!!!", PixelFontFace),
				&core.StopBgmEvent{},
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
