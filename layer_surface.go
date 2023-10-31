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
#include "samure/layer_surface.h"
*/
import "C"
import "unsafe"

const (
	AnchorTop    = C.ZWLR_LAYER_SURFACE_V1_ANCHOR_TOP
	AnchorBottom = C.ZWLR_LAYER_SURFACE_V1_ANCHOR_BOTTOM
	AnchorLeft   = C.ZWLR_LAYER_SURFACE_V1_ANCHOR_LEFT
	AnchorRight  = C.ZWLR_LAYER_SURFACE_V1_ANCHOR_RIGHT
	AnchorFill   = AnchorTop | AnchorBottom | AnchorLeft | AnchorRight

	LayerBackground = C.ZWLR_LAYER_SHELL_V1_LAYER_BACKGROUND
	LayerBottom     = C.ZWLR_LAYER_SHELL_V1_LAYER_BOTTOM
	LayerTop        = C.ZWLR_LAYER_SHELL_V1_LAYER_TOP
	LayerOverlay    = C.ZWLR_LAYER_SHELL_V1_LAYER_OVERLAY
)

type LayerSurface struct {
	Handle *C.struct_samure_layer_surface
}

func CreateLayerSurface(ctx Context, o *Output, layer, anchor int, keyboardInteraction, pointerInteraction, backendAssociation bool) (LayerSurface, error) {
	var cO *C.struct_samure_output
	if o != nil {
		cO = o.Handle
	}

	var cKi, cPi, cBa C.int
	if keyboardInteraction {
		cKi = 1
	}
	if pointerInteraction {
		cPi = 1
	}
	if backendAssociation {
		cBa = 1
	}

	sfc_rs := C.samure_create_layer_surface(ctx.Handle, cO, C.uint32_t(layer), C.uint32_t(anchor), cKi, cPi, cBa)
	if sfc_rs.error != ErrorNone {
		return LayerSurface{}, NewError(uint64(sfc_rs.error))
	}

	return LayerSurface{sfc_rs.result}, nil
}

func (sfc LayerSurface) BackendData() unsafe.Pointer {
	return sfc.Handle.backend_data
}

func (sfc LayerSurface) W() int {
	return int(sfc.Handle.w)
}

func (sfc LayerSurface) H() int {
	return int(sfc.Handle.h)
}

func (sfc LayerSurface) Destroy(ctx Context) {
	C.samure_destroy_layer_surface(ctx.Handle, sfc.Handle)
}

func (sfc LayerSurface) DrawBuffer(b SharedBuffer) {
	C.samure_layer_surface_draw_buffer(sfc.Handle, b.Handle)
}

func (sfc LayerSurface) FrameDeltaTime() float64 {
	return float64(sfc.Handle.frame_delta_time)
}

func (sfc LayerSurface) Surface() unsafe.Pointer {
	return unsafe.Pointer(sfc.Handle.surface)
}

func (sfc LayerSurface) LayerSurface() unsafe.Pointer {
	return unsafe.Pointer(sfc.Handle.layer_surface)
}
