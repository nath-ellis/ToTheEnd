package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/nath-ellis/ToTheEnd/data"
	"github.com/solarlune/resolv"
)

type PlayerData struct {
	Obj          *resolv.Object
	XSpeed       float64
	YSpeed       float64
	Gravity      float64
	Falling      bool
	IsLeft       bool
	IsWalking    bool
	WalkingLeft  []*ebiten.Image
	WalkingRight []*ebiten.Image
}

var Player PlayerData

func Init() {
	Player.Obj = resolv.NewObject(50, 50, 75, 90, "player")
	Player.XSpeed = 3.0
	Player.YSpeed = 10.0
	Player.Gravity = 5.0
	Player.Falling = false
	Player.IsLeft = true
	Player.IsWalking = false

	data.Space.Add(Player.Obj)

	walkingOneL, _, _ := ebitenutil.NewImageFromFile("assets/player/walk/left/1.png")
	walkingTwoL, _, _ := ebitenutil.NewImageFromFile("assets/player/walk/left/2.png")
	walkingThreeL, _, _ := ebitenutil.NewImageFromFile("assets/player/walk/left/3.png")
	walkingFourL, _, _ := ebitenutil.NewImageFromFile("assets/player/walk/left/4.png")
	walkingFiveL, _, _ := ebitenutil.NewImageFromFile("assets/player/walk/left/5.png")
	walkingSixL, _, _ := ebitenutil.NewImageFromFile("assets/player/walk/left/6.png")
	walkingSevenL, _, _ := ebitenutil.NewImageFromFile("assets/player/walk/left/7.png")
	walkingEightL, _, _ := ebitenutil.NewImageFromFile("assets/player/walk/left/8.png")

	Player.WalkingLeft = append(Player.WalkingLeft, walkingOneL)
	Player.WalkingLeft = append(Player.WalkingLeft, walkingTwoL)
	Player.WalkingLeft = append(Player.WalkingLeft, walkingThreeL)
	Player.WalkingLeft = append(Player.WalkingLeft, walkingFourL)
	Player.WalkingLeft = append(Player.WalkingLeft, walkingFiveL)
	Player.WalkingLeft = append(Player.WalkingLeft, walkingSixL)
	Player.WalkingLeft = append(Player.WalkingLeft, walkingSevenL)
	Player.WalkingLeft = append(Player.WalkingLeft, walkingEightL)

	walkingOneR, _, _ := ebitenutil.NewImageFromFile("assets/player/walk/right/1.png")
	walkingTwoR, _, _ := ebitenutil.NewImageFromFile("assets/player/walk/right/2.png")
	walkingThreeR, _, _ := ebitenutil.NewImageFromFile("assets/player/walk/right/3.png")
	walkingFourR, _, _ := ebitenutil.NewImageFromFile("assets/player/walk/right/4.png")
	walkingFiveR, _, _ := ebitenutil.NewImageFromFile("assets/player/walk/right/5.png")
	walkingSixR, _, _ := ebitenutil.NewImageFromFile("assets/player/walk/right/6.png")
	walkingSevenR, _, _ := ebitenutil.NewImageFromFile("assets/player/walk/right/7.png")
	walkingEightR, _, _ := ebitenutil.NewImageFromFile("assets/player/walk/right/8.png")

	Player.WalkingRight = append(Player.WalkingRight, walkingOneR)
	Player.WalkingRight = append(Player.WalkingRight, walkingTwoR)
	Player.WalkingRight = append(Player.WalkingRight, walkingThreeR)
	Player.WalkingRight = append(Player.WalkingRight, walkingFourR)
	Player.WalkingRight = append(Player.WalkingRight, walkingFiveR)
	Player.WalkingRight = append(Player.WalkingRight, walkingSixR)
	Player.WalkingRight = append(Player.WalkingRight, walkingSevenR)
	Player.WalkingRight = append(Player.WalkingRight, walkingEightR)
}
