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

package opengl

/*
#include <samure/backends/opengl.h>
#include <samure/context.h>
#include "../../wrappers.h"
*/
import "C"
import (
	"unsafe"

	samure "github.com/Samudevv/samurai-render-go"
)

type Backend struct {
	Handle *C.struct_samure_backend
}

func (gl *Backend) Init(ctx samure.Context) error {
	libname := C.CString("libsamurai-render-backend-opengl.so")
	depname := C.CString("libEGL.so")
	cCtx := (*C.struct_samure_context)(unsafe.Pointer(ctx.Handle))
	defer C.free(unsafe.Pointer(libname))
	defer C.free(unsafe.Pointer(depname))

	gl_rs := C.samure_create_backend_from_lib(cCtx, libname, depname)
	if gl_rs.error != samure.ErrorNone {
		return samure.NewError(uint64(gl_rs.error))
	}

	gl.Handle = gl_rs.result

	return nil
}

func (gl *Backend) OnLayerSurfaceConfigure(ctx samure.Context, layerSurface samure.LayerSurface, width, height int) {
	cCtx := (*C.struct_samure_context)(unsafe.Pointer(ctx.Handle))
	cSfc := (*C.struct_samure_layer_surface)(unsafe.Pointer(layerSurface.Handle))

	bak := cCtx.backend
	cCtx.backend = gl.Handle
	C.wrapper_backend_fptr_on_layer_surface_configure(gl.Handle.on_layer_surface_configure, cCtx, cSfc, C.int32_t(width), C.int32_t(height))
	cCtx.backend = bak
}

func (gl *Backend) RenderStart(ctx samure.Context, layerSurface samure.LayerSurface) {
	cCtx := (*C.struct_samure_context)(unsafe.Pointer(ctx.Handle))
	cSfc := (*C.struct_samure_layer_surface)(unsafe.Pointer(layerSurface.Handle))
	bak := cCtx.backend
	cCtx.backend = gl.Handle
	C.wrapper_backend_fptr_render_start(gl.Handle.render_start, cCtx, cSfc)
	cCtx.backend = bak
}

func (gl *Backend) RenderEnd(ctx samure.Context, layerSurface samure.LayerSurface) {
	cCtx := (*C.struct_samure_context)(unsafe.Pointer(ctx.Handle))
	cSfc := (*C.struct_samure_layer_surface)(unsafe.Pointer(layerSurface.Handle))
	bak := cCtx.backend
	cCtx.backend = gl.Handle
	C.wrapper_backend_fptr_render_end(gl.Handle.render_end, cCtx, cSfc)
	cCtx.backend = bak
}

func (gl *Backend) Destroy(ctx samure.Context) {
	cCtx := (*C.struct_samure_context)(unsafe.Pointer(ctx.Handle))
	bak := cCtx.backend
	cCtx.backend = gl.Handle
	C.wrapper_backend_fptr_destroy(gl.Handle.destroy, cCtx)
	cCtx.backend = bak
}

func (gl *Backend) AssociateLayerSurface(ctx samure.Context, layerSurface samure.LayerSurface) uint64 {
	cCtx := (*C.struct_samure_context)(unsafe.Pointer(ctx.Handle))
	cSfc := (*C.struct_samure_layer_surface)(unsafe.Pointer(layerSurface.Handle))
	bak := cCtx.backend
	cCtx.backend = gl.Handle
	defer func() { cCtx.backend = bak }()
	return uint64(C.wrapper_backend_fptr_associate_layer_surface(gl.Handle.associate_layer_surface, cCtx, cSfc))
}

func (gl *Backend) UnassociateLayerSurface(ctx samure.Context, layerSurface samure.LayerSurface) {
	cCtx := (*C.struct_samure_context)(unsafe.Pointer(ctx.Handle))
	cSfc := (*C.struct_samure_layer_surface)(unsafe.Pointer(layerSurface.Handle))
	bak := cCtx.backend
	cCtx.backend = gl.Handle
	C.wrapper_backend_fptr_unassociate_layer_surface(gl.Handle.unassociate_layer_surface, cCtx, cSfc)
	cCtx.backend = bak
}

func (gl Backend) InitBindings(ctx samure.Context, initFunc func() error) error {
	cSfc := (*C.struct_samure_layer_surface)(unsafe.Pointer(ctx.Output(0).Surface(0).Handle))
	C.samure_backend_opengl_make_context_current(gl.Handle, cSfc)
	return initFunc()
}

func (gl Backend) MakeContextCurrent() {
	C.samure_backend_opengl_make_context_current(gl.Handle, nil)
}

// TODO: Display is not exposed
// func (gl Backend) Display() unsafe.Pointer {
// 	return unsafe.Pointer(gl.Handle.display)
// }

// TODO: EGLContext is not exposed
// func (gl Backend) Context() unsafe.Pointer {
// 	return unsafe.Pointer(gl.Handle.context)
// }
