/***********************************************************************************
 *                         This file is part of samurai-render-go
 *                    https://github.com/Samudevv/samurai-render-go
 ***********************************************************************************
 * Copyright (c) 2026 Kassandra Pucher
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
#include <samure/backends/cairo.h>
#include <samure/context.h>
#include "../../wrappers.h"
*/
import "C"
import (
	"unsafe"

	samure "github.com/Samudevv/samurai-render-go"
	"github.com/gotk3/gotk3/cairo"
)

type Backend struct {
	Handle *C.struct_samure_backend
}

func (c *Backend) Init(ctx samure.Context) error {
	libname := C.CString("libsamurai-render-backend-cairo.so")
	depname := C.CString("libcairo.so")
	defer C.free(unsafe.Pointer(libname))
	defer C.free(unsafe.Pointer(depname))

	c_rs := C.samure_create_backend_from_lib((*C.struct_samure_context)(unsafe.Pointer(ctx.Handle)), libname, depname)
	if c_rs.error != samure.ErrorNone {
		return samure.NewError(uint64(c_rs.error))
	}

	c.Handle = c_rs.result
	return nil
}

func (c *Backend) OnLayerSurfaceConfigure(ctx samure.Context, layerSurface samure.LayerSurface, width, height int) {
	C.wrapper_backend_fptr_on_layer_surface_configure(c.Handle.on_layer_surface_configure, (*C.struct_samure_context)(unsafe.Pointer(ctx.Handle)), (*C.struct_samure_layer_surface)(unsafe.Pointer(layerSurface.Handle)), C.int32_t(width), C.int32_t(height))
}

func (c *Backend) RenderStart(ctx samure.Context, layerSurface samure.LayerSurface) {
}

func (c *Backend) RenderEnd(ctx samure.Context, layerSurface samure.LayerSurface) {
	C.wrapper_backend_fptr_render_end(c.Handle.render_end, (*C.struct_samure_context)(unsafe.Pointer(ctx.Handle)), (*C.struct_samure_layer_surface)(unsafe.Pointer(layerSurface.Handle)))
}

func (c *Backend) Destroy(ctx samure.Context) {
	C.wrapper_backend_fptr_destroy(c.Handle.destroy, (*C.struct_samure_context)(unsafe.Pointer(ctx.Handle)))
}

func (c *Backend) AssociateLayerSurface(ctx samure.Context, layerSurface samure.LayerSurface) uint64 {
	return uint64(C.wrapper_backend_fptr_associate_layer_surface(c.Handle.associate_layer_surface, (*C.struct_samure_context)(unsafe.Pointer(ctx.Handle)), (*C.struct_samure_layer_surface)(unsafe.Pointer(layerSurface.Handle))))
}

func (c *Backend) UnassociateLayerSurface(ctx samure.Context, layerSurface samure.LayerSurface) {
	C.wrapper_backend_fptr_unassociate_layer_surface(c.Handle.unassociate_layer_surface, (*C.struct_samure_context)(unsafe.Pointer(ctx.Handle)), (*C.struct_samure_layer_surface)(unsafe.Pointer(layerSurface.Handle)))
}

func Get(sfc samure.LayerSurface) *cairo.Context {
	cairoSfc := (*C.struct_samure_cairo_surface)(sfc.BackendData())
	return cairo.WrapContext(uintptr(unsafe.Pointer(cairoSfc.cairo)))
}

func GetSurface(sfc samure.LayerSurface) *cairo.Surface {
	cairoSfc := (*C.struct_samure_cairo_surface)(sfc.BackendData())
	return cairo.WrapSurface(uintptr(unsafe.Pointer(cairoSfc.cairo_surface)))
}
