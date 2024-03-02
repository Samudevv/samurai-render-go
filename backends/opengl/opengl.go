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

package opengl

/*
#include <samure/backends/opengl.h>
#include <samure/context.h>
*/
import "C"
import (
	"unsafe"

	samure "github.com/Samudevv/samurai-render-go"
)

type Backend struct {
	Handle *C.struct_samure_backend_opengl
}

func (gl *Backend) Init(ctx samure.Context) error {
	cCtx := (*C.struct_samure_context)(unsafe.Pointer(ctx.Handle))

	gl_rs := C.samure_init_backend_opengl(cCtx, cCtx.config.gl)
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
	cCtx.backend = &gl.Handle.base
	C.samure_backend_opengl_on_layer_surface_configure(cCtx, cSfc, C.int32_t(width), C.int32_t(height))
	cCtx.backend = bak
}

func (gl *Backend) RenderStart(ctx samure.Context, layerSurface samure.LayerSurface) {
	cCtx := (*C.struct_samure_context)(unsafe.Pointer(ctx.Handle))
	cSfc := (*C.struct_samure_layer_surface)(unsafe.Pointer(layerSurface.Handle))
	bak := cCtx.backend
	cCtx.backend = &gl.Handle.base
	C.samure_backend_opengl_render_start(cCtx, cSfc)
	cCtx.backend = bak
}

func (gl *Backend) RenderEnd(ctx samure.Context, layerSurface samure.LayerSurface) {
	cCtx := (*C.struct_samure_context)(unsafe.Pointer(ctx.Handle))
	cSfc := (*C.struct_samure_layer_surface)(unsafe.Pointer(layerSurface.Handle))
	bak := cCtx.backend
	cCtx.backend = &gl.Handle.base
	C.samure_backend_opengl_render_end(cCtx, cSfc)
	cCtx.backend = bak
}

func (gl *Backend) Destroy(ctx samure.Context) {
	cCtx := (*C.struct_samure_context)(unsafe.Pointer(ctx.Handle))
	bak := cCtx.backend
	cCtx.backend = &gl.Handle.base
	C.samure_destroy_backend_opengl(cCtx)
	cCtx.backend = bak
}

func (gl *Backend) AssociateLayerSurface(ctx samure.Context, layerSurface samure.LayerSurface) uint64 {
	cCtx := (*C.struct_samure_context)(unsafe.Pointer(ctx.Handle))
	cSfc := (*C.struct_samure_layer_surface)(unsafe.Pointer(layerSurface.Handle))
	bak := cCtx.backend
	cCtx.backend = &gl.Handle.base
	defer func() { cCtx.backend = bak }()
	return uint64(C.samure_backend_opengl_associate_layer_surface(cCtx, cSfc))
}

func (gl *Backend) UnassociateLayerSurface(ctx samure.Context, layerSurface samure.LayerSurface) {
	cCtx := (*C.struct_samure_context)(unsafe.Pointer(ctx.Handle))
	cSfc := (*C.struct_samure_layer_surface)(unsafe.Pointer(layerSurface.Handle))
	bak := cCtx.backend
	cCtx.backend = &gl.Handle.base
	C.samure_backend_opengl_unassociate_layer_surface(cCtx, cSfc)
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

func (gl Backend) Display() unsafe.Pointer {
	return unsafe.Pointer(gl.Handle.display)
}

func (gl Backend) Context() unsafe.Pointer {
	return unsafe.Pointer(gl.Handle.context)
}
