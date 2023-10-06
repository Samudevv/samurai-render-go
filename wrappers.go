package samure

/*
#include "samure/backend.h"
#include "samure/events.h"
#include "samure/output.h"
*/
import "C"
import "unsafe"

//export globalOnEvent
func globalOnEvent(ctx *C.struct_samure_context, event *C.struct_samure_event, userData unsafe.Pointer) {
	idx := int(uintptr(userData))
	a := GetGlobalApp(idx)
	a.OnEvent(
		Context{ctx},
		unsafe.Pointer(event),
	)
}

//export globalOnRender
func globalOnRender(ctx *C.struct_samure_context, layerSurface *C.struct_samure_layer_surface, outputGeo C.struct_samure_rect, deltaTime C.double, userData unsafe.Pointer) {
	idx := int(uintptr(userData))
	a := GetGlobalApp(idx)

	a.OnRender(
		Context{ctx},
		unsafe.Pointer(layerSurface),
		unsafe.Pointer(&outputGeo),
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
		unsafe.Pointer(layer_surface),
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
		unsafe.Pointer(layer_surface),
	)
}

//export wrapper_backend_render_end
func wrapper_backend_render_end(ctx *C.struct_samure_context, layer_surface *C.struct_samure_layer_surface) {
	bak := (*C.struct_wrapper_backend)(unsafe.Pointer(ctx.backend))
	idx := int(bak.go_backend)
	b := GetGlobalBackend(idx)
	b.RenderEnd(
		Context{ctx},
		unsafe.Pointer(layer_surface),
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
		unsafe.Pointer(layer_surface),
	))
}

//export wrapper_backend_unassociate_layer_surface
func wrapper_backend_unassociate_layer_surface(ctx *C.struct_samure_context, layer_surface *C.struct_samure_layer_surface) {
	bak := (*C.struct_wrapper_backend)(unsafe.Pointer(ctx.backend))
	idx := int(bak.go_backend)
	b := GetGlobalBackend(idx)
	b.UnassociateLayerSurface(
		Context{ctx},
		unsafe.Pointer(layer_surface),
	)
}
