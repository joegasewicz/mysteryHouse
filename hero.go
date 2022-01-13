package main

import "github.com/hajimehoshi/ebiten/v2"

const (
	JUMP_ASCEND     = "HERO_JUMP_ASCEND"
	JUMP_DESCEND    = "HERO_JUMP_DESCEND"
	JUMP_MAX_HEIGHT = 50
)

type Hero struct {
	X          float64
	Y          float64
	HasMoved   bool
	eventType  string
	JumpStartY float64
	Options    *ebiten.DrawImageOptions
	Img        *ebiten.Image
}

func (h *Hero) jump(screen *ebiten.Image) {
	//                 ceiling: +50
	//      	   		    - -
	//    			 	 -       -
	// JumpStartY = h.Y -         - finish h.Y+50
	// If the jump hasn't yet started then set the Y axis
	if h.eventType != JUMP_ASCEND && h.eventType != JUMP_ASCEND {
		// Start the jump
		h.JumpStartY = h.Y
		h.eventType = JUMP_ASCEND
	}
	// Check if jump is ceiling height
	if h.eventType == JUMP_ASCEND && h.ReachedJumpCeiling() {
		h.X -= 1
		h.Y -= 1
		// Check if we are ascending and & jumping
	} else if h.eventType == JUMP_ASCEND && h.JumpStartY < h.Y {
		h.X += 1
		h.Y += 1
		// Check if jump is complete & reset
	} else if h.eventType == JUMP_DESCEND && h.JumpStartY == h.Y {
		// End the jump
		h.JumpStartY = 0
		h.eventType = ""
	}

	h.Options.GeoM.Translate(h.X, h.Y)
}

func (h *Hero) ReachedJumpCeiling() bool {
	return h.JumpStartY+JUMP_MAX_HEIGHT >= h.Y
}
