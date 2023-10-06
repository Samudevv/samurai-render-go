package samure

/*
#include "samure/layer_surface.h"
*/
import "C"
import "unsafe"

type LayerSurface struct {
	Handle *C.struct_samure_layer_surface
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
