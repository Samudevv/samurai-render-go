/***********************************************************************************
 *                         This file is part of samurai-render-go
 *                    https://github.com/Samudevv/samurai-render-go
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
#include <samure/seat.h>
#include <samure/wayland/cursor-shape.h>
*/
import "C"
import "unsafe"

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

func (s Seat) TouchFocus() Focus {
	return Focus{s.Handle.touch_focus}
}

func (s Seat) Name() string {
	return C.GoString(s.Handle.name)
}

func (s Seat) Seat() unsafe.Pointer {
	return unsafe.Pointer(s.Handle.seat)
}

func (s Seat) Pointer() unsafe.Pointer {
	return unsafe.Pointer(s.Handle.pointer)
}

func (s Seat) Keyboard() unsafe.Pointer {
	return unsafe.Pointer(s.Handle.keyboard)
}

func (s Seat) Touch() unsafe.Pointer {
	return unsafe.Pointer(s.Handle.touch)
}
