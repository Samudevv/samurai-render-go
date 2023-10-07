package samure

/*
#include "samure/layer_surface.h"
*/
import "C"
import "unsafe"

const (
	AnchorTop    = C.ZWLR_LAYER_SURFACE_V1_ANCHOR_TOP
	AnchorBottom = C.ZWLR_LAYER_SURFACE_V1_ANCHOR_BOTTOM
	AnchorLeft   = C.ZWLR_LAYER_SURFACE_V1_ANCHOR_LEFT
	AnchorRight  = C.ZWLR_LAYER_SURFACE_V1_ANCHOR_RIGHT
	AnchorFill   = AnchorTop | AnchorBottom | AnchorLeft | AnchorRight

	LayerBackground = C.ZWLR_LAYER_SHELL_V1_LAYER_BACKGROUND
	LayerBottom     = C.ZWLR_LAYER_SHELL_V1_LAYER_BOTTOM
	LayerTop        = C.ZWLR_LAYER_SHELL_V1_LAYER_TOP
	LayerOverlay    = C.ZWLR_LAYER_SHELL_V1_LAYER_OVERLAY
)

type LayerSurface struct {
	Handle *C.struct_samure_layer_surface
}

func CreateLayerSurface(ctx Context, o *Output, layer, anchor int, keyboardInteraction, pointerInteraction, backendAssociation bool) (LayerSurface, error) {
	var cO *C.struct_samure_output
	if o != nil {
		cO = o.Handle
	}

	var cKi, cPi, cBa C.int
	if keyboardInteraction {
		cKi = 1
	}
	if pointerInteraction {
		cPi = 1
	}
	if backendAssociation {
		cBa = 1
	}

	sfc_rs := C.samure_create_layer_surface(ctx.Handle, cO, C.uint32_t(layer), C.uint32_t(anchor), cKi, cPi, cBa)
	if sfc_rs.error != ErrorNone {
		return LayerSurface{}, NewError(uint64(sfc_rs.error))
	}

	return LayerSurface{sfc_rs.result}, nil
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

func (sfc LayerSurface) Destroy(ctx Context) {
	C.samure_destroy_layer_surface(ctx.Handle, sfc.Handle)
}

func (sfc LayerSurface) DrawBuffer(b SharedBuffer) {
	C.samure_layer_surface_draw_buffer(sfc.Handle, b.Handle)
}
