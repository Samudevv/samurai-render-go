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
	OnLayerSurfaceConfigure(ctx Context, layerSurface unsafe.Pointer, width, height int)
	RenderStart(ctx Context, layerSurface unsafe.Pointer)
	RenderEnd(ctx Context, layerSurface unsafe.Pointer)
	Destroy(ctx Context)
	AssociateLayerSurface(ctx Context, layerSurface unsafe.Pointer) uint64
	UnassociateLayerSurface(ctx Context, layerSurface unsafe.Pointer)
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

func InitRawBackend() (*RawBackend, error) {
	r_rs := C.samure_init_backend_raw(nil)
	if r_rs.error != ErrorNone {
		return nil, NewError(r_rs.error)
	}

	return &RawBackend{r_rs.result}, nil
}

func (raw *RawBackend) OnLayerSurfaceConfigure(ctx Context, layerSurface unsafe.Pointer, width, height int) {

}

func (raw *RawBackend) RenderStart(ctx Context, layerSurface unsafe.Pointer) {

}

func (raw *RawBackend) RenderEnd(ctx Context, layerSurface unsafe.Pointer) {
	C.samure_backend_raw_render_end(ctx.handle, (*C.struct_samure_layer_surface)(layerSurface))
}

func (raw *RawBackend) Destroy(ctx Context) {
	C.samure_destroy_backend_raw(ctx.handle)
}

func (raw *RawBackend) AssociateLayerSurface(ctx Context, layerSurface unsafe.Pointer) uint64 {
	err := C.samure_backend_raw_associate_layer_surface(ctx.handle, (*C.struct_samure_layer_surface)(layerSurface))
	if err != ErrorNone {
		return uint64(err)
	}
	return ErrorNone
}

func (raw *RawBackend) UnassociateLayerSurface(ctx Context, layerSurface unsafe.Pointer) {
	C.samure_backend_raw_unassociate_layer_surface(ctx.handle, (*C.struct_samure_layer_surface)(layerSurface))
}
