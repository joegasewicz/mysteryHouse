package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

const (
	JUMP_ASCEND          = "HERO_JUMP_ASCEND"
	JUMP_DESCEND         = "HERO_JUMP_DESCEND"
	JUMP_MAX_HEIGHT      = 40
	JUMP_DIRECTION_UP    = "JUMP_DIRECTION_UP"
	JUMP_DIRECTION_LEFT  = "JUMP_DIRECTION_LEFT"
	JUMP_DIRECTION_RIGHT = "JUMP_DIRECTION_RIGHT"
)

type Jump struct {
	Started bool
	// State:
	// - JUMP_ASCEND
	// - JUMP_DESCEND
	State string
	// Type:
	// - JUMP_DIRECTION_UP
	// - JUMP_DIRECTION_LEFT
	// - JUMP_DIRECTION_RIGHT
	Type   string
	StartY float64
	hero   *Hero
}

func (j *Jump) Start(direction string) {
	j.Type = direction
	if j.hero == nil {
		panic("no Hero pointer passed to Jump Object")
	}
	//                 ceiling: +50
	//      	   		    - -
	//    			 	 -       -
	// JumpStartY = h.Y -         - finish h.Y+50
	// If the jump hasn't yet started then set the Y axis
	if j.State == "" {
		// Start the jump
		j.StartY = j.hero.Y
		j.State = JUMP_ASCEND
	}
	// Check if jump is ceiling height then descend or just check we are descending
	if j.State == JUMP_DESCEND || j.reachedJumpCeiling() {
		j.descend()
		// Check if we are ascending & jumping
	} else if j.State == JUMP_ASCEND {
		j.ascend()
	}

}

func (j *Jump) Continue() {
	// Check if jump is ceiling height then descend or just check we are descending
	if j.State == JUMP_DESCEND || j.reachedJumpCeiling() {
		j.descend()
		// Check if we are ascending & jumping
	} else if j.State == JUMP_ASCEND && !j.reachedJumpCeiling() {
		j.ascend()
	}
	// Check if jump is complete & reset
	if j.State == JUMP_DESCEND && j.StartY == j.hero.Y {
		// End the jump
		j.StartY = 0
		j.State = ""
	}
}

func (j *Jump) ascend() {
	switch j.Type {
	case JUMP_DIRECTION_UP:
		j.hero.Y -= 1
	case JUMP_DIRECTION_RIGHT:
		j.hero.Y -= 1
		j.hero.X += 1
	case JUMP_DIRECTION_LEFT:
		j.hero.Y -= 1
		j.hero.X -= 1
	default:
		panic("Jump type doesn't matter")
	}
}

func (j *Jump) descend() {
	log.Println("DESCENDING...")
	if j.State == JUMP_ASCEND {
		j.State = JUMP_DESCEND
	}
	switch j.Type {
	case JUMP_DIRECTION_UP:
		j.hero.Y += 1
	case JUMP_DIRECTION_RIGHT:
		j.hero.Y += 1
		j.hero.X += 1
	case JUMP_DIRECTION_LEFT:
		j.hero.Y += 1
		j.hero.X -= 1
	default:
		panic("Jump type doesn't matter")
	}
}

func (j *Jump) reachedJumpCeiling() bool {
	return j.hero.Y <= j.StartY-JUMP_MAX_HEIGHT
}

type Hero struct {
	X        float64
	Y        float64
	HasMoved bool
	Img      *ebiten.Image
	Jump     *Jump
}

func (h *Hero) Run(xPos float64) {
	heroYPos := PLATFORM_HEIGHT + HERO_HEIGHT
	h.Y = float64(SCREEN_HEIGHT - heroYPos)
	h.X += xPos
}

func (h *Hero) LogPosition() {
	log.Printf("Hero X : %d - Y: %d\n", int(h.X), int(h.Y))
	log.Println(h)
}
