package main

import (
	"bytes"
	"log"

	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

//go:embed PixelOperator8.ttf
var PixelFontTTF []byte

//go:embed images/warehouse.png
var warehouse []byte

//go:embed images/office_day.png
var office_day []byte

//go:embed images/office_night.png
var office_night []byte

//go:embed images/omarov.png
var omarov []byte

//go:embed images/omar_sprite_1.png
var omar_sprite_1 []byte

//go:embed images/omar_sprite_2.png
var omar_sprite_2 []byte

//go:embed images/monster_sprite_1.png
var monster_sprite_1 []byte

//go:embed images/monster_sprite_2.png
var monster_sprite_2 []byte

//go:embed images/monster_sprite_3.png
var monster_sprite_3 []byte

//go:embed images/titlescreen.png
var titlescreen []byte

//go:embed images/heart.png
var heart []byte

//go:embed images/combat_bg.png
var combat_bg []byte

//go:embed images/heart_empty.png
var heart_empty []byte

//go:embed images/menu_btn_bg.png
var menu_btn_bg []byte

//go:embed images/lobby.png
var lobby []byte

//go:embed audio/bass_slap_2.mp3
var BassMp3 []byte

//go:embed audio/water-pouring-80316.mp3
var Pouring []byte

//go:embed audio/unease_melody.mp3
var Unease []byte

//go:embed audio/snow-storm-wind-ambience-272426.mp3
var SnowStorm []byte

//go:embed audio/sound-effect-old-phone-191761_2.mp3
var PhoneRing []byte

//go:embed audio/creepy_bell.mp3
var CreepyBell []byte

var PixelFont *text.GoTextFaceSource
var PixelFontFace *text.GoTextFace

var WarehouseBg *ebiten.Image
var Omarov *ebiten.Image
var OfficeDay *ebiten.Image
var OfficeNight *ebiten.Image
var Lobby *ebiten.Image
var OmarSprite1 *ebiten.Image
var OmarSprite2 *ebiten.Image

var MonsterSprite1 *ebiten.Image
var MonsterSprite2 *ebiten.Image
var MonsterSprite3 *ebiten.Image
var CombatBg *ebiten.Image
var MenuButtonBg *ebiten.Image
var TitleScreen *ebiten.Image

var Heart *ebiten.Image
var HeartEmpty *ebiten.Image

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(PixelFontTTF))
	if err != nil {
		log.Fatal(err)
	}
	PixelFont = s
	PixelFontFace = &text.GoTextFace{
		Source: PixelFont,
		Size:   24,
	}
	if WarehouseBg == nil {
		imgReader := bytes.NewReader(warehouse)
		WarehouseBg, _, _ = ebitenutil.NewImageFromReader(imgReader)
	}
	if Omarov == nil {
		imgReader := bytes.NewReader(omarov)
		Omarov, _, _ = ebitenutil.NewImageFromReader(imgReader)
	}
	if OfficeDay == nil {
		imgReader := bytes.NewReader(office_day)
		OfficeDay, _, _ = ebitenutil.NewImageFromReader(imgReader)
	}
	if OfficeNight == nil {
		imgReader := bytes.NewReader(office_night)
		OfficeNight, _, _ = ebitenutil.NewImageFromReader(imgReader)
	}
	if Lobby == nil {
		imgReader := bytes.NewReader(lobby)
		Lobby, _, _ = ebitenutil.NewImageFromReader(imgReader)
	}
	if OmarSprite1 == nil {
		imgReader := bytes.NewReader(omar_sprite_1)
		OmarSprite1, _, _ = ebitenutil.NewImageFromReader(imgReader)
	}
	if OmarSprite2 == nil {
		imgReader := bytes.NewReader(omar_sprite_2)
		OmarSprite2, _, _ = ebitenutil.NewImageFromReader(imgReader)
	}
	if MonsterSprite1 == nil {
		imgReader := bytes.NewReader(monster_sprite_1)
		MonsterSprite1, _, _ = ebitenutil.NewImageFromReader(imgReader)
	}
	if MonsterSprite2 == nil {
		imgReader := bytes.NewReader(monster_sprite_2)
		MonsterSprite2, _, _ = ebitenutil.NewImageFromReader(imgReader)
	}
	if MonsterSprite3 == nil {
		imgReader := bytes.NewReader(monster_sprite_3)
		MonsterSprite3, _, _ = ebitenutil.NewImageFromReader(imgReader)
	}
	if CombatBg == nil {
		imgReader := bytes.NewReader(combat_bg)
		CombatBg, _, _ = ebitenutil.NewImageFromReader(imgReader)
	}
	if MenuButtonBg == nil {
		imgReader := bytes.NewReader(menu_btn_bg)
		MenuButtonBg, _, _ = ebitenutil.NewImageFromReader(imgReader)
	}
	if Heart == nil {
		imgReader := bytes.NewReader(heart)
		Heart, _, _ = ebitenutil.NewImageFromReader(imgReader)
	}
	if TitleScreen == nil {
		imgReader := bytes.NewReader(titlescreen)
		TitleScreen, _, _ = ebitenutil.NewImageFromReader(imgReader)
	}
	if HeartEmpty == nil {
		imgReader := bytes.NewReader(heart_empty)
		HeartEmpty, _, _ = ebitenutil.NewImageFromReader(imgReader)
	}
}
