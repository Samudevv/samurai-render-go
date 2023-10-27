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

package cairo

/*
#include "../../samure/backends/cairo.h"
#include "../../samure/context.h"
*/
import "C"
import (
	"unsafe"

	samure "github.com/PucklaJ/samurai-render-go"
	"github.com/gotk3/gotk3/cairo"
)

type Backend struct {
	Handle *C.struct_samure_backend_cairo
}

func (c *Backend) Init(ctx samure.Context) error {
	c_rs := C.samure_init_backend_cairo((*C.struct_samure_context)(unsafe.Pointer(ctx.Handle)))
	if c_rs.error != samure.ErrorNone {
		return samure.NewError(uint64(c_rs.error))
	}

	c.Handle = c_rs.result
	return nil
}

func (c *Backend) OnLayerSurfaceConfigure(ctx samure.Context, layerSurface samure.LayerSurface, width, height int) {
	C.samure_backend_cairo_on_layer_surface_configure((*C.struct_samure_context)(unsafe.Pointer(ctx.Handle)), (*C.struct_samure_layer_surface)(unsafe.Pointer(layerSurface.Handle)), C.int32_t(width), C.int32_t(height))
}

func (c *Backend) RenderStart(ctx samure.Context, layerSurface samure.LayerSurface) {
}

func (c *Backend) RenderEnd(ctx samure.Context, layerSurface samure.LayerSurface) {
	C.samure_backend_cairo_render_end((*C.struct_samure_context)(unsafe.Pointer(ctx.Handle)), (*C.struct_samure_layer_surface)(unsafe.Pointer(layerSurface.Handle)))
}

func (c *Backend) Destroy(ctx samure.Context) {
	C.samure_destroy_backend_cairo((*C.struct_samure_context)(unsafe.Pointer(ctx.Handle)))
}

func (c *Backend) AssociateLayerSurface(ctx samure.Context, layerSurface samure.LayerSurface) uint64 {
	return uint64(C.samure_backend_cairo_associate_layer_surface((*C.struct_samure_context)(unsafe.Pointer(ctx.Handle)), (*C.struct_samure_layer_surface)(unsafe.Pointer(layerSurface.Handle))))
}

func (c *Backend) UnassociateLayerSurface(ctx samure.Context, layerSurface samure.LayerSurface) {
	C.samure_backend_cairo_unassociate_layer_surface((*C.struct_samure_context)(unsafe.Pointer(ctx.Handle)), (*C.struct_samure_layer_surface)(unsafe.Pointer(layerSurface.Handle)))
}

func Get(sfc samure.LayerSurface) *cairo.Context {
	cairoSfc := (*C.struct_samure_cairo_surface)(sfc.BackendData())
	return cairo.WrapContext(uintptr(unsafe.Pointer(cairoSfc.cairo)))
}

func GetSurface(sfc samure.LayerSurface) *cairo.Surface {
	cairoSfc := (*C.struct_samure_cairo_surface)(sfc.BackendData())
	return cairo.WrapSurface(uintptr(unsafe.Pointer(cairoSfc.cairo_surface)))
}
