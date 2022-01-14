package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

const (
	JUMP_ASCEND          = "HERO_JUMP_ASCEND"
	JUMP_DESCEND         = "HERO_JUMP_DESCEND"
	JUMP_MAX_HEIGHT      = 50
	JUMP_DIRECTION_UP    = "JUMP_DIRECTION_UP"
	JUMP_DIRECTION_LEFT  = "JUMP_DIRECTION_LEFT"
	JUMP_DIRECTION_RIGHT = "JUMP_DIRECTION_RIGHT"
)

type Jump struct {
	EventType string
	StartY    float64
	hero      *Hero
}

func (j *Jump) Start(direction string) {
	log.Printf("Jump Direction: %s", direction)
	if j.hero == nil {
		panic("no Hero pointer passed to Jump Object")
	}
	//                 ceiling: +50
	//      	   		    - -
	//    			 	 -       -
	// JumpStartY = h.Y -         - finish h.Y+50
	// If the jump hasn't yet started then set the Y axis
	if j.EventType != JUMP_ASCEND {
		// Start the jump
		j.StartY = j.hero.Y
		j.EventType = JUMP_ASCEND
	}
	// Check if jump is ceiling height
	if j.EventType == JUMP_ASCEND && j.reachedJumpCeiling() {
		j.hero.X -= 1
		j.hero.Y -= 1
		// Check if we are ascending & jumping
	} else if j.EventType == JUMP_ASCEND && j.StartY < j.hero.Y {
		if direction == JUMP_DIRECTION_UP {

		}
		j.hero.X += 1
		j.hero.Y += 1
		// Check if jump is complete & reset
	} else if j.EventType == JUMP_DESCEND && j.StartY == j.hero.Y {
		// End the jump
		j.StartY = 0
		j.EventType = ""
	}
}

func (j *Jump) reachedJumpCeiling() bool {
	return j.StartY+JUMP_MAX_HEIGHT >= j.hero.Y
}

type Hero struct {
	X        float64
	Y        float64
	HasMoved bool
	Img      *ebiten.Image
	Jump     *Jump
}

func (h *Hero) Run(xPos float64) {
	heroYPos := MUD_HEIGHT + HERO_HEIGHT
	h.Y = float64(SCREEN_HEIGHT - heroYPos)
	h.X += xPos
}

func (h *Hero) LogPosition() {
	log.Printf("Hero X : %d - Y: %d\n", int(h.X), int(h.Y))
	log.Println(h)
}
