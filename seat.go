package samure

/*
#include "samure/seat.h"
#include "samure/wayland/cursor-shape-v1-client-protocol.h"
*/
import "C"

const (
	CursorShapeDefault      = C.WP_CURSOR_SHAPE_DEVICE_V1_SHAPE_DEFAULT
	CursorShapeContextMenu  = C.WP_CURSOR_SHAPE_DEVICE_V1_SHAPE_CONTEXT_MENU
	CursorShapeHelp         = C.WP_CURSOR_SHAPE_DEVICE_V1_SHAPE_HELP
	CursorShapePointer      = C.WP_CURSOR_SHAPE_DEVICE_V1_SHAPE_POINTER
	CursorShapeProgress     = C.WP_CURSOR_SHAPE_DEVICE_V1_SHAPE_PROGRESS
	CursorShapeWait         = C.WP_CURSOR_SHAPE_DEVICE_V1_SHAPE_WAIT
	CursorShapeCell         = C.WP_CURSOR_SHAPE_DEVICE_V1_SHAPE_CELL
	CursorShapeCrosshair    = C.WP_CURSOR_SHAPE_DEVICE_V1_SHAPE_CROSSHAIR
	CursorShapeText         = C.WP_CURSOR_SHAPE_DEVICE_V1_SHAPE_TEXT
	CursorShapeVerticalText = C.WP_CURSOR_SHAPE_DEVICE_V1_SHAPE_VERTICAL_TEXT
	CursorShapeAlias        = C.WP_CURSOR_SHAPE_DEVICE_V1_SHAPE_ALIAS
	CursorShapeCopy         = C.WP_CURSOR_SHAPE_DEVICE_V1_SHAPE_COPY
	CursorShapeMove         = C.WP_CURSOR_SHAPE_DEVICE_V1_SHAPE_MOVE
	CursorShapeNoDrop       = C.WP_CURSOR_SHAPE_DEVICE_V1_SHAPE_NO_DROP
	CursorShapeNotAllowed   = C.WP_CURSOR_SHAPE_DEVICE_V1_SHAPE_NOT_ALLOWED
	CursorShapeGrab         = C.WP_CURSOR_SHAPE_DEVICE_V1_SHAPE_GRAB
	CursorShapeGrabbing     = C.WP_CURSOR_SHAPE_DEVICE_V1_SHAPE_GRABBING
	CursorShapeEResize      = C.WP_CURSOR_SHAPE_DEVICE_V1_SHAPE_E_RESIZE
	CursorShapeNResize      = C.WP_CURSOR_SHAPE_DEVICE_V1_SHAPE_N_RESIZE
	CursorShapeNeResize     = C.WP_CURSOR_SHAPE_DEVICE_V1_SHAPE_NE_RESIZE
	CursorShapeNwResize     = C.WP_CURSOR_SHAPE_DEVICE_V1_SHAPE_NW_RESIZE
	CursorShapeSResize      = C.WP_CURSOR_SHAPE_DEVICE_V1_SHAPE_S_RESIZE
	CursorShapeSeResize     = C.WP_CURSOR_SHAPE_DEVICE_V1_SHAPE_SE_RESIZE
	CursorShapeSwResize     = C.WP_CURSOR_SHAPE_DEVICE_V1_SHAPE_SW_RESIZE
	CursorShapeWResize      = C.WP_CURSOR_SHAPE_DEVICE_V1_SHAPE_W_RESIZE
	CursorShapeEwResize     = C.WP_CURSOR_SHAPE_DEVICE_V1_SHAPE_EW_RESIZE
	CursorShapeNsResize     = C.WP_CURSOR_SHAPE_DEVICE_V1_SHAPE_NS_RESIZE
	CursorShapeNeswResize   = C.WP_CURSOR_SHAPE_DEVICE_V1_SHAPE_NESW_RESIZE
	CursorShapeNwseResize   = C.WP_CURSOR_SHAPE_DEVICE_V1_SHAPE_NWSE_RESIZE
	CursorShapeColResize    = C.WP_CURSOR_SHAPE_DEVICE_V1_SHAPE_COL_RESIZE
	CursorShapeRowResize    = C.WP_CURSOR_SHAPE_DEVICE_V1_SHAPE_ROW_RESIZE
	CursorShapeAllScroll    = C.WP_CURSOR_SHAPE_DEVICE_V1_SHAPE_ALL_SCROLL
	CursorShapeZoomIn       = C.WP_CURSOR_SHAPE_DEVICE_V1_SHAPE_ZOOM_IN
	CursorShapeZoomOut      = C.WP_CURSOR_SHAPE_DEVICE_V1_SHAPE_ZOOM_OUT
)

type Focus struct {
	Handle C.struct_samure_focus
}

func (f Focus) Output() Output {
	return Output{f.Handle.output}
}

func (f Focus) Surface() LayerSurface {
	return LayerSurface{f.Handle.surface}
}

type Seat struct {
	Handle *C.struct_samure_seat
}

func (s Seat) PointerFocus() Focus {
	return Focus{s.Handle.pointer_focus}
}

func (s Seat) KeyboardFocus() Focus {
	return Focus{s.Handle.keyboard_focus}
}

func (s Seat) Name() string {
	return C.GoString(s.Handle.name)
}

func (s Seat) SetPointerShape(shape int) {
	C.samure_seat_set_pointer_shape(s.Handle, C.uint32_t(shape))
}
