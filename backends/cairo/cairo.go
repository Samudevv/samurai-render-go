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
