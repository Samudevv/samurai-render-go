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
#include "samure/backend.h"
#include "samure/events.h"
#include "samure/output.h"
*/
import "C"
import (
	"unsafe"
)

//export globalOnEvent
func globalOnEvent(ctx *C.struct_samure_context, event *C.struct_samure_event, userData unsafe.Pointer) {
	idx := int(uintptr(userData))
	a := GetGlobalApp(idx)
	a.OnEvent(
		Context{ctx},
		cEventToGoEvent(event),
	)
}

//export globalOnRender
func globalOnRender(ctx *C.struct_samure_context, layerSurface *C.struct_samure_layer_surface, outputGeo C.struct_samure_rect, deltaTime C.double, userData unsafe.Pointer) {
	idx := int(uintptr(userData))
	a := GetGlobalApp(idx)

	a.OnRender(
		Context{ctx},
		LayerSurface{layerSurface},
		Rect{
			int(outputGeo.x),
			int(outputGeo.y),
			int(outputGeo.w),
			int(outputGeo.h),
		},
		float64(deltaTime),
	)
}

//export globalOnUpdate
func globalOnUpdate(ctx *C.struct_samure_context, deltaTime C.double, userData unsafe.Pointer) {
	idx := int(uintptr(userData))
	a := GetGlobalApp(idx)

	a.OnUpdate(
		Context{ctx},
		float64(deltaTime),
	)
}

//export wrapper_backend_on_layer_surface_configure
func wrapper_backend_on_layer_surface_configure(ctx *C.struct_samure_context, layer_surface *C.struct_samure_layer_surface, width, height C.int32_t) {
	bak := (*C.struct_wrapper_backend)(unsafe.Pointer(ctx.backend))
	idx := int(bak.go_backend)
	b := GetGlobalBackend(idx)

	b.OnLayerSurfaceConfigure(
		Context{ctx},
		LayerSurface{layer_surface},
		int(width),
		int(height),
	)
}

//export wrapper_backend_render_start
func wrapper_backend_render_start(ctx *C.struct_samure_context, layer_surface *C.struct_samure_layer_surface) {
	bak := (*C.struct_wrapper_backend)(unsafe.Pointer(ctx.backend))
	idx := int(bak.go_backend)
	b := GetGlobalBackend(idx)
	b.RenderStart(
		Context{ctx},
		LayerSurface{layer_surface},
	)
}

//export wrapper_backend_render_end
func wrapper_backend_render_end(ctx *C.struct_samure_context, layer_surface *C.struct_samure_layer_surface) {
	bak := (*C.struct_wrapper_backend)(unsafe.Pointer(ctx.backend))
	idx := int(bak.go_backend)
	b := GetGlobalBackend(idx)
	b.RenderEnd(
		Context{ctx},
		LayerSurface{layer_surface},
	)
}

//export wrapper_backend_destroy
func wrapper_backend_destroy(ctx *C.struct_samure_context) {
	bak := (*C.struct_wrapper_backend)(unsafe.Pointer(ctx.backend))
	idx := int(bak.go_backend)
	b := GetGlobalBackend(idx)
	b.Destroy(Context{ctx})
	C.free(unsafe.Pointer(ctx.backend))
}

//export wrapper_backend_associate_layer_surface
func wrapper_backend_associate_layer_surface(ctx *C.struct_samure_context, layer_surface *C.struct_samure_layer_surface) C.samure_error {
	bak := (*C.struct_wrapper_backend)(unsafe.Pointer(ctx.backend))
	idx := int(bak.go_backend)
	b := GetGlobalBackend(idx)
	return C.samure_error(b.AssociateLayerSurface(
		Context{ctx},
		LayerSurface{layer_surface},
	))
}

//export wrapper_backend_unassociate_layer_surface
func wrapper_backend_unassociate_layer_surface(ctx *C.struct_samure_context, layer_surface *C.struct_samure_layer_surface) {
	bak := (*C.struct_wrapper_backend)(unsafe.Pointer(ctx.backend))
	idx := int(bak.go_backend)
	b := GetGlobalBackend(idx)
	b.UnassociateLayerSurface(
		Context{ctx},
		LayerSurface{layer_surface},
	)
}
