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

type Hero struct {
	X          float64
	Y          float64
	HasMoved   bool
	EventType  string
	JumpStartY float64
	Img        *ebiten.Image
}

func (h *Hero) Jump(direction string) {
	log.Printf("Jump Direction: %s", direction)
	//                 ceiling: +50
	//      	   		    - -
	//    			 	 -       -
	// JumpStartY = h.Y -         - finish h.Y+50
	// If the jump hasn't yet started then set the Y axis
	if h.EventType != JUMP_ASCEND && h.EventType != JUMP_ASCEND {
		// Start the jump
		h.JumpStartY = h.Y
		h.EventType = JUMP_ASCEND
	}
	// Check if jump is ceiling height
	if h.EventType == JUMP_ASCEND && h.reachedJumpCeiling() {
		h.X -= 1
		h.Y -= 1
		// Check if we are ascending & jumping
	} else if h.EventType == JUMP_ASCEND && h.JumpStartY < h.Y {
		h.X += 1
		h.Y += 1
		// Check if jump is complete & reset
	} else if h.EventType == JUMP_DESCEND && h.JumpStartY == h.Y {
		// End the jump
		h.JumpStartY = 0
		h.EventType = ""
	}
}

func (h *Hero) Run(xPos float64) {
	heroYPos := MUD_HEIGHT + HERO_HEIGHT
	h.Y = float64(SCREEN_HEIGHT - heroYPos)
	h.X += xPos
}

func (h *Hero) reachedJumpCeiling() bool { // TODO Move to actions struct
	return h.JumpStartY+JUMP_MAX_HEIGHT >= h.Y
}

func (h *Hero) LogPosition() {
	log.Printf("Hero X : %d - Y: %d\n", int(h.X), int(h.Y))
	log.Println(h)
}
