package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	_ "image/png"
	"log"
)

var (
	hero       *Hero
	heroImg    *ebiten.Image
	mudImg     *ebiten.Image
	tilesImage *ebiten.Image
)

const (
	SCREEN_WIDTH    = 320
	SCREEN_HEIGHT   = 256
	HERO_HEIGHT     = 32
	HERO_WIDTH      = 9
	PLATFORM_HEIGHT = 32
	TILE_SIZE       = 32
	TILE_X_NUM      = 10
)

func init() {
	var err error
	heroImg, _, err = ebitenutil.NewImageFromFile("assets/hero-1.png")
	mudImg, _, err = ebitenutil.NewImageFromFile("assets/mud.png")
	tilesImage, _, err = ebitenutil.NewImageFromFile("assets/tile_map.png")
	if err != nil {
		log.Fatal(err)
	}
	hero = &Hero{
		X:        HERO_WIDTH,
		Y:        PLATFORM_HEIGHT,
		HasMoved: false,
		Img:      heroImg,
		Jump: &Jump{
			State:  "",
			StartY: 0,
		},
	}
	hero.Jump.hero = hero
}

type Game struct {
	layers [][]int
}

func (g *Game) Update() error {
	// Hero
	if !hero.HasMoved {
		hero.Run(50)
		hero.HasMoved = true
	} else {
		// Key Inputs
		if hero.Jump.State != "" {
			hero.Jump.Continue()
		} else {
			switch true {
			case ebiten.IsKeyPressed(ebiten.KeySpace) && ebiten.IsKeyPressed(ebiten.KeyArrowRight):
				hero.Jump.Start(JUMP_DIRECTION_RIGHT)
			case ebiten.IsKeyPressed(ebiten.KeySpace) && ebiten.IsKeyPressed(ebiten.KeyArrowLeft):
				hero.Jump.Start(JUMP_DIRECTION_LEFT)
			case ebiten.IsKeyPressed(ebiten.KeySpace):
				hero.Jump.Start(JUMP_DIRECTION_UP)
			case ebiten.IsKeyPressed(ebiten.KeyArrowRight):
				hero.Run(1)
			case ebiten.IsKeyPressed(ebiten.KeyArrowLeft):
				hero.Run(-1)
			default:
				hero.Run(0)
			}
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// ------------------------------
	// Tiles
	const xNum = SCREEN_WIDTH / TILE_SIZE
	for _, l := range g.layers {
		for i, t := range l {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64((i%xNum)*TILE_SIZE), float64((i/xNum)*TILE_SIZE))

			sx := (t % TILE_X_NUM) * TILE_SIZE
			sy := (t / TILE_X_NUM) * TILE_SIZE
			screen.DrawImage(tilesImage.SubImage(image.Rect(sx, sy, sx+TILE_SIZE, sy+TILE_SIZE)).(*ebiten.Image), op)
		}
	}
	// ------------------------------
	// Hero
	heroOptions := &ebiten.DrawImageOptions{}
	heroOptions.GeoM.Translate(hero.X, hero.Y)
	screen.DrawImage(heroImg, heroOptions)
	//hero.LogPosition()
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screeHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func main() {
	g := &Game{
		layers: [][]int{
			{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				1, 1, 1, 1, 0, 1, 1, 3, 1, 1,
				0, 0, 0, 0, 0, 0, 0, 2, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 2, 0, 0,
				1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			},
		},
	}

	ebiten.SetWindowSize(SCREEN_WIDTH*2, SCREEN_HEIGHT*2)
	ebiten.SetWindowTitle("Mystery House")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
