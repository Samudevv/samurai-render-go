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
