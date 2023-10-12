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
#include "samure/backends/raw.h"
*/
import "C"
import (
	"sync"
	"unsafe"
)

var globalBackends map[int]Backend
var globalBackendsMtx sync.Mutex

type Backend interface {
	Init(ctx Context) error
	OnLayerSurfaceConfigure(ctx Context, layerSurface LayerSurface, width, height int)
	RenderStart(ctx Context, layerSurface LayerSurface)
	RenderEnd(ctx Context, layerSurface LayerSurface)
	Destroy(ctx Context)
	AssociateLayerSurface(ctx Context, layerSurface LayerSurface) uint64
	UnassociateLayerSurface(ctx Context, layerSurface LayerSurface)
}

func AddGlobalBackend(bak Backend) int {
	globalBackendsMtx.Lock()
	defer globalBackendsMtx.Unlock()

	if globalBackends == nil {
		globalBackends = make(map[int]Backend)
	}

	idx := len(globalBackends)
	globalBackends[idx] = bak
	return idx
}

func GetGlobalBackend(idx int) Backend {
	globalBackendsMtx.Lock()
	defer globalBackendsMtx.Unlock()

	return globalBackends[idx]
}

type RawBackend struct {
	handle *C.struct_samure_backend_raw
}

func (raw *RawBackend) Init(ctx Context) error {
	r_rs := C.samure_init_backend_raw(ctx.Handle)
	if r_rs.error != ErrorNone {
		return NewError(uint64(r_rs.error))
	}
	raw.handle = r_rs.result

	return nil
}

func (raw *RawBackend) OnLayerSurfaceConfigure(ctx Context, layerSurface LayerSurface, width, height int) {
	C.samure_backend_raw_on_layer_surface_configure(ctx.Handle, layerSurface.Handle, C.int32_t(width), C.int32_t(height))
}

func (raw *RawBackend) RenderStart(ctx Context, layerSurface LayerSurface) {

}

func (raw *RawBackend) RenderEnd(ctx Context, layerSurface LayerSurface) {
	C.samure_backend_raw_render_end(ctx.Handle, (*C.struct_samure_layer_surface)(layerSurface.Handle))
}

func (raw *RawBackend) Destroy(ctx Context) {
	C.samure_destroy_backend_raw(ctx.Handle)
}

func (raw *RawBackend) AssociateLayerSurface(ctx Context, layerSurface LayerSurface) uint64 {
	err := C.samure_backend_raw_associate_layer_surface(ctx.Handle, (*C.struct_samure_layer_surface)(layerSurface.Handle))
	if err != ErrorNone {
		return uint64(err)
	}
	return ErrorNone
}

func (raw *RawBackend) UnassociateLayerSurface(ctx Context, layerSurface LayerSurface) {
	C.samure_backend_raw_unassociate_layer_surface(ctx.Handle, (*C.struct_samure_layer_surface)(layerSurface.Handle))
}

func GetRaw(sfc LayerSurface) RawSurface {
	return RawSurface{(*C.struct_samure_raw_surface)(sfc.BackendData())}
}

type RawSurface struct {
	Handle *C.struct_samure_raw_surface
}

func (r RawSurface) Data() unsafe.Pointer {
	return r.Handle.buffer.data
}

func (r RawSurface) Buffer() SharedBuffer {
	return SharedBuffer{r.Handle.buffer}
}
