/***********************************************************************************
 *                         This file is part of samurai-render-go
 *                    https://github.com/PucklaJ/samurai-render-go
 ***********************************************************************************
 * Copyright (c) 2023 Jonas Pucher
 *
 * This software is provided ‘as-is’, without any express or implied
 * warranty. In no event will the authors be held liable for any damages
 * arising from the use of this software.
 *
 * Permission is granted to anyone to use this software for any purpose,
 * including commercial applications, and to alter it and redistribute it
 * freely, subject to the following restrictions:
 *
 * 1. The origin of this software must not be misrepresented; you must not
 * claim that you wrote the original software. If you use this software
 * in a product, an acknowledgment in the product documentation would be
 * appreciated but is not required.
 *
 * 2. Altered source versions must be plainly marked as such, and must not be
 * misrepresented as being the original software.
 *
 * 3. This notice may not be removed or altered from any source
 * distribution.
 ************************************************************************************/

package samure

/*
#include "samure/events.h"
*/
import "C"

type EventPointerButton struct {
	Seat   Seat
	Button int
	State  int
}

type EventPointerMotion struct {
	Seat Seat
	X    float64
	Y    float64
}

type EventPointerEnter struct {
	Seat    Seat
	Output  Output
	Surface LayerSurface
	X       float64
	Y       float64
}

type EventPointerLeave struct {
	Seat    Seat
	Output  Output
	Surface LayerSurface
}

type EventKeyboardKey struct {
	Seat  Seat
	Key   int
	State int
}

type EventKeyboardEnter struct {
	Seat    Seat
	Output  Output
	Surface LayerSurface
}

type EventKeyboardLeave struct {
	Seat Seat
}

type EventTouchDown struct {
	Seat    Seat
	Output  Output
	Surface LayerSurface
	X       float64
	Y       float64
	TouchID int
}

type EventTouchUp struct {
	Seat    Seat
	TouchID int
}

type EventTouchMotion struct {
	Seat    Seat
	X       float64
	Y       float64
	TouchID int
}

func cEventToGoEvent(e *C.struct_samure_event) interface{} {
	switch e._type {
	case C.SAMURE_EVENT_POINTER_BUTTON:
		return EventPointerButton{
			Seat:   Seat{e.seat},
			Button: int(e.button),
			State:  int(e.state),
		}
	case C.SAMURE_EVENT_POINTER_MOTION:
		return EventPointerMotion{
			Seat: Seat{e.seat},
			X:    float64(e.x),
			Y:    float64(e.y),
		}
	case C.SAMURE_EVENT_POINTER_ENTER:
		return EventPointerEnter{
			Seat:    Seat{e.seat},
			Output:  Output{e.output},
			Surface: LayerSurface{e.surface},
			X:       float64(e.x),
			Y:       float64(e.y),
		}
	case C.SAMURE_EVENT_POINTER_LEAVE:
		return EventPointerLeave{
			Seat:    Seat{e.seat},
			Output:  Output{e.output},
			Surface: LayerSurface{e.surface},
		}
	case C.SAMURE_EVENT_KEYBOARD_KEY:
		return EventKeyboardKey{
			Seat:  Seat{e.seat},
			Key:   int(e.button),
			State: int(e.state),
		}
	case C.SAMURE_EVENT_KEYBOARD_ENTER:
		return EventKeyboardEnter{
			Seat:    Seat{e.seat},
			Output:  Output{e.output},
			Surface: LayerSurface{e.surface},
		}
	case C.SAMURE_EVENT_KEYBOARD_LEAVE:
		return EventKeyboardLeave{
			Seat: Seat{e.seat},
		}
	case C.SAMURE_EVENT_TOUCH_DOWN:
		return EventTouchDown{
			Seat:    Seat{e.seat},
			Output:  Output{e.output},
			Surface: LayerSurface{e.surface},
			X:       float64(e.x),
			Y:       float64(e.y),
			TouchID: int(e.touch_id),
		}
	case C.SAMURE_EVENT_TOUCH_UP:
		return EventTouchUp{
			Seat:    Seat{e.seat},
			TouchID: int(e.touch_id),
		}
	case C.SAMURE_EVENT_TOUCH_MOTION:
		return EventTouchMotion{
			Seat:    Seat{e.seat},
			X:       float64(e.x),
			Y:       float64(e.y),
			TouchID: int(e.touch_id),
		}
	default:
		return nil
	}
}
