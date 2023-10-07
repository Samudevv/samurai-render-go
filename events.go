package samure

/*
#include "samure/events.h"
*/
import "C"
import "unsafe"

type EventPointerButton struct {
	Seat   unsafe.Pointer
	Button int
	State  int
}

type EventPointerMotion struct {
	Seat unsafe.Pointer
	X    float64
	Y    float64
}

type EventPointerEnter struct {
	Seat   unsafe.Pointer
	Output Output
	X      float64
	Y      float64
}

type EventPointerLeave struct {
	Seat   unsafe.Pointer
	Output Output
}

type EventKeyboardKey struct {
	Seat  unsafe.Pointer
	Key   int
	State int
}

type EventKeyboardEnter struct {
	Seat   unsafe.Pointer
	Output Output
}

type EventKeyboardLeave struct {
	Seat unsafe.Pointer
}

func cEventToGoEvent(e *C.struct_samure_event) interface{} {
	switch e._type {
	case C.SAMURE_EVENT_POINTER_BUTTON:
		return EventPointerButton{
			Seat:   unsafe.Pointer(e.seat),
			Button: int(e.button),
			State:  int(e.state),
		}
	case C.SAMURE_EVENT_POINTER_MOTION:
		return EventPointerMotion{
			Seat: unsafe.Pointer(e.seat),
			X:    float64(e.x),
			Y:    float64(e.y),
		}
	case C.SAMURE_EVENT_POINTER_ENTER:
		return EventPointerEnter{
			Seat:   unsafe.Pointer(e.seat),
			Output: Output{e.output},
			X:      float64(e.x),
			Y:      float64(e.y),
		}
	case C.SAMURE_EVENT_POINTER_LEAVE:
		return EventPointerLeave{
			Seat:   unsafe.Pointer(e.seat),
			Output: Output{e.output},
		}
	case C.SAMURE_EVENT_KEYBOARD_KEY:
		return EventKeyboardKey{
			Seat:  unsafe.Pointer(e.seat),
			Key:   int(e.button),
			State: int(e.state),
		}
	case C.SAMURE_EVENT_KEYBOARD_ENTER:
		return EventKeyboardEnter{
			Seat:   unsafe.Pointer(e.seat),
			Output: Output{e.output},
		}
	case C.SAMURE_EVENT_KEYBOARD_LEAVE:
		return EventKeyboardLeave{
			Seat: unsafe.Pointer(e.seat),
		}
	default:
		return nil
	}
}
