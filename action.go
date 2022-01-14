package main

import "log"

const (
	ACTION_TYPE_JUMP     = "ACTION_TYPE_JUMP"
	JUMP_ASCEND          = "HERO_JUMP_ASCEND"
	JUMP_DESCEND         = "HERO_JUMP_DESCEND"
	JUMP_MAX_HEIGHT      = 50
	JUMP_DIRECTION_UP    = "JUMP_DIRECTION_UP"
	JUMP_DIRECTION_LEFT  = "JUMP_DIRECTION_LEFT"
	JUMP_DIRECTION_RIGHT = "JUMP_DIRECTION_RIGHT"
)

type IAction interface {
	Init()
}

type Action struct {
	MaxHeight float64
	Type      string
	StartX    float64
	StartY    float64
	EndX      float64
	EndY      float64
}

// Jump e Is the entity type (hero, enemies etc.)
// Actions always return X, Y postions for entity
func (a *Action) Init(actionType string, eventType string, direction string, entityY float64) (float64, float64) {
	log.Printf("Jump Direction: %s", direction)
	var X, Y float64
	//                 ceiling: +50
	//      	   		    - -
	//    			 	 -       -
	// JumpStartY = h.Y -         - finish h.Y+50
	// If the jump hasn't yet started then set the Y axis
	if eventType != JUMP_ASCEND && eventType != JUMP_ASCEND {
		// Start the jump
		a.StartY = entityY
		a.Type = JUMP_ASCEND
	}
	// Check if jump is ceiling height
	if eventType == JUMP_ASCEND && a.reachedJumpCeiling(entityY) {
		X -= 1
		Y -= 1
		// Check if we are ascending & jumping
	} else if eventType == JUMP_ASCEND && a.StartY < entityY {
		X += 1
		Y += 1
		// Check if jump is complete & reset
	} else if eventType == JUMP_DESCEND && a.StartY == entityY {
		// End the jump
		a.StartY = 0
		eventType = ""
	}
	return X, Y
}

func (a *Action) reachedJumpCeiling(entityY float64) bool { // TODO Move to actions struct
	return a.StartY+JUMP_MAX_HEIGHT >= entityY
}
