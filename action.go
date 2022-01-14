package main

import "log"

const (
	JUMP_ASCEND          = "HERO_JUMP_ASCEND"
	JUMP_DESCEND         = "HERO_JUMP_DESCEND"
	JUMP_MAX_HEIGHT      = 50
	JUMP_DIRECTION_UP    = "JUMP_DIRECTION_UP"
	JUMP_DIRECTION_LEFT  = "JUMP_DIRECTION_LEFT"
	JUMP_DIRECTION_RIGHT = "JUMP_DIRECTION_RIGHT"
)

type Action struct {
	MaxHeight float64
	Type      string
	StartX    float64
	StartY    float64
	EndX      float64
	EndY      float64
}

// Jump e Is the entity type (hero, enemies etc.)
func (a *Action) Jump(e Hero, direction string) (float64, float64) {
	log.Printf("Jump Direction: %s", direction)
	var X, Y float64
	//                 ceiling: +50
	//      	   		    - -
	//    			 	 -       -
	// JumpStartY = h.Y -         - finish h.Y+50
	// If the jump hasn't yet started then set the Y axis
	if e.EventType != JUMP_ASCEND && e.EventType != JUMP_ASCEND {
		// Start the jump
		a.StartY = e.Y
		a.Type = JUMP_ASCEND
	}
	// Check if jump is ceiling height
	if e.EventType == JUMP_ASCEND && e.reachedJumpCeiling() {
		X -= 1
		Y -= 1
		// Check if we are ascending & jumping
	} else if e.EventType == JUMP_ASCEND && e.JumpStartY < e.Y {
		X += 1
		Y += 1
		// Check if jump is complete & reset
	} else if e.EventType == JUMP_DESCEND && e.JumpStartY == e.Y {
		// End the jump
		e.JumpStartY = 0
		e.EventType = ""
	}
	return X, Y
}
