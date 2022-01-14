package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
	"log"
)

var (
	hero    *Hero
	heroImg *ebiten.Image
	mudImg  *ebiten.Image
)

const (
	SCREEN_WIDTH  = 320 * 1.5
	SCREEN_HEIGHT = 240 * 1.5
	HERO_HEIGHT   = 32
	HERO_WIDTH    = 9
	MUD_HEIGHT    = 8
)

func init() {
	var err error
	heroImg, _, err = ebitenutil.NewImageFromFile("assets/hero-1.png")
	mudImg, _, err = ebitenutil.NewImageFromFile("assets/mud.png")
	if err != nil {
		log.Fatal(err)
	}
	hero = &Hero{
		X:          HERO_WIDTH,
		Y:          MUD_HEIGHT,
		HasMoved:   false,
		EventType:  "",
		JumpStartY: 0,
		Img:        heroImg,
	}
}

type Game struct{}

func (g *Game) Update() error {
	// Run heroImg
	if !hero.HasMoved {
		hero.Run(50)
		hero.HasMoved = true
	} else {
		switch true {
		case ebiten.IsKeyPressed(ebiten.KeySpace) && ebiten.IsKeyPressed(ebiten.KeyArrowRight):
			hero.Jump(JUMP_DIRECTION_RIGHT)
		case ebiten.IsKeyPressed(ebiten.KeySpace) && ebiten.IsKeyPressed(ebiten.KeyArrowLeft):
			hero.Jump(JUMP_DIRECTION_LEFT)
		case ebiten.IsKeyPressed(ebiten.KeySpace):
			hero.Jump(JUMP_DIRECTION_UP)
		case ebiten.IsKeyPressed(ebiten.KeyArrowRight):
			hero.Run(1)
		case ebiten.IsKeyPressed(ebiten.KeyArrowLeft):
			hero.Run(-1)
		default:
			hero.Run(0)
		}
	}
	return nil
}

func drawPlatform(screen *ebiten.Image, length int, yPos float64) {
	op := &ebiten.DrawImageOptions{}
	// Starting position
	op.GeoM.Translate(SCREEN_WIDTH, yPos)
	for i := 1; i <= length; i++ {
		op.GeoM.Translate(-8, 0)
		screen.DrawImage(mudImg, op)
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw Room
	drawPlatform(screen, 20, 150)
	drawPlatform(screen, 40, 250)
	drawPlatform(screen, 60, SCREEN_HEIGHT-8)
	// Draw Hero
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(hero.X, hero.Y)
	screen.DrawImage(heroImg, op)
	//hero.LogPosition()
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screeHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func main() {
	ebiten.SetWindowSize(SCREEN_WIDTH*2, SCREEN_HEIGHT*2)
	ebiten.SetWindowTitle("Mystery House")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
