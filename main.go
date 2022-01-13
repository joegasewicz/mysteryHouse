package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
	"log"
)

var (
	hero        *Hero
	heroImg     *ebiten.Image
	mud         *ebiten.Image
	heroOptions *ebiten.DrawImageOptions
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
	hero = &Hero{
		HasMoved:        false,
		CurrentPosition: 10,
	}
	heroImg, _, err = ebitenutil.NewImageFromFile("assets/hero-1.png")
	mud, _, err = ebitenutil.NewImageFromFile("assets/mud.png")
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct{}

func (g *Game) Update() error {

	return nil
}

func drawPlatform(screen *ebiten.Image, length int, yPos float64) {
	op := &ebiten.DrawImageOptions{}
	// Starting position
	op.GeoM.Translate(SCREEN_WIDTH, yPos)
	for i := 1; i <= length; i++ {
		op.GeoM.Translate(-8, 0)
		screen.DrawImage(mud, op)
	}
}

func drawHero(screen *ebiten.Image, xPos float64) {
	heroPos := MUD_HEIGHT + HERO_HEIGHT
	heroOptions = &ebiten.DrawImageOptions{}
	hero.CurrentPosition += xPos
	heroOptions.GeoM.Translate(hero.CurrentPosition, float64(SCREEN_HEIGHT-heroPos))
	screen.DrawImage(heroImg, heroOptions)
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw room #1
	drawPlatform(screen, 20, 150)
	drawPlatform(screen, 40, 250)
	drawPlatform(screen, 60, SCREEN_HEIGHT-8)
	// Draw heroImg
	if !hero.HasMoved {
		drawHero(screen, 50)
		hero.HasMoved = true
	} else {
		switch true {
		case ebiten.IsKeyPressed(ebiten.KeyArrowRight):
			drawHero(screen, 1)
		case ebiten.IsKeyPressed(ebiten.KeyArrowLeft):
			drawHero(screen, -1)
		default:
			drawHero(screen, 0)
		}
	}
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
