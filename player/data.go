package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/solarlune/resolv"
)

type PlayerData struct {
	Obj       *resolv.Object
	IsLeft    bool
	IsWalking bool
	Walking   []*ebiten.Image
}

var Player PlayerData

func Init() {
	Player.Obj = resolv.NewObject(50, 50, 150, 183, "player")
	Player.IsLeft = true
	Player.IsWalking = false

	walkingOne, _, _ := ebitenutil.NewImageFromFile("assets/player/1.png")
	walkingTwo, _, _ := ebitenutil.NewImageFromFile("assets/player/2.png")
	walkingThree, _, _ := ebitenutil.NewImageFromFile("assets/player/3.png")
	walkingFour, _, _ := ebitenutil.NewImageFromFile("assets/player/4.png")
	walkingFive, _, _ := ebitenutil.NewImageFromFile("assets/player/5.png")
	walkingSix, _, _ := ebitenutil.NewImageFromFile("assets/player/6.png")
	walkingSeven, _, _ := ebitenutil.NewImageFromFile("assets/player/7.png")
	walkingEight, _, _ := ebitenutil.NewImageFromFile("assets/player/8.png")

	Player.Walking = append(Player.Walking, walkingOne)
	Player.Walking = append(Player.Walking, walkingTwo)
	Player.Walking = append(Player.Walking, walkingThree)
	Player.Walking = append(Player.Walking, walkingFour)
	Player.Walking = append(Player.Walking, walkingFive)
	Player.Walking = append(Player.Walking, walkingSix)
	Player.Walking = append(Player.Walking, walkingSeven)
	Player.Walking = append(Player.Walking, walkingEight)
}
