package samure

/*
#include "samure/output.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

type Output struct {
	Handle *C.struct_samure_output
}

type Rect struct {
	X, Y, W, H int
}

func (o Rect) convertToC() C.struct_samure_rect {
	return C.struct_samure_rect{
		x: C.int32_t(o.X),
		y: C.int32_t(o.Y),
		w: C.int32_t(o.W),
		h: C.int32_t(o.H),
	}
}

func (o Rect) CircleInOutput(cx, cy, r int) bool {
	return C.samure_circle_in_output(o.convertToC(), C.int32_t(cx), C.int32_t(cy), C.int32_t(r)) != 0
}

func (o Rect) RectInOutput(x, y, w, h int) bool {
	return C.samure_rect_in_output(o.convertToC(), C.int32_t(x), C.int32_t(y), C.int32_t(w), C.int32_t(h)) != 0
}

func (o Rect) SquareInOutput(x, y, size int) bool {
	return C.samure_square_in_output(o.convertToC(), C.int32_t(x), C.int32_t(y), C.int32_t(size)) != 0
}

func (o Rect) PointInOutput(x, y int) bool {
	return C.samure_point_in_output(o.convertToC(), C.int32_t(x), C.int32_t(y)) != 0
}

func (o Rect) TriangleInOutput(x1, y1, x2, y2, x3, y3 int) bool {
	return C.samure_triangle_in_output(o.convertToC(), C.int32_t(x1), C.int32_t(y1), C.int32_t(x2), C.int32_t(y2), C.int32_t(x3), C.int32_t(y3)) != 0
}

func (o Rect) RelX(x float64) float64 {
	return x - float64(o.X)
}

func (o Rect) RelY(y float64) float64 {
	return y - float64(o.Y)
}

func (o Output) CircleInOutput(cx, cy, r int) bool {
	return C.samure_circle_in_output(o.Handle.geo, C.int32_t(cx), C.int32_t(cy), C.int32_t(r)) != 0
}

func (o Output) RectInOutput(x, y, w, h int) bool {
	return C.samure_rect_in_output(o.Handle.geo, C.int32_t(x), C.int32_t(y), C.int32_t(w), C.int32_t(h)) != 0
}

func (o Output) SquareInOutput(x, y, size int) bool {
	return C.samure_square_in_output(o.Handle.geo, C.int32_t(x), C.int32_t(y), C.int32_t(size)) != 0
}

func (o Output) PointInOutput(x, y int) bool {
	return C.samure_point_in_output(o.Handle.geo, C.int32_t(x), C.int32_t(y)) != 0
}

func (o Output) TriangleInOutput(x1, y1, x2, y2, x3, y3 int) bool {
	return C.samure_triangle_in_output(o.Handle.geo, C.int32_t(x1), C.int32_t(y1), C.int32_t(x2), C.int32_t(y2), C.int32_t(x3), C.int32_t(y3)) != 0
}

func (o Output) SetPointerInteraction(ctx Context, enable bool) {
	var cEnable C.int
	if enable {
		cEnable = 1
	} else {
		cEnable = 0
	}
	C.samure_output_set_pointer_interaction(ctx.Handle, o.Handle, cEnable)
}

func (o Output) SetInputRegions(ctx Context, rs []Rect) {
	var cRs *C.struct_samure_rect
	cRs = (*C.struct_samure_rect)(C.malloc(C.size_t(unsafe.Sizeof(*cRs) * uintptr(len(rs)))))
	if cRs == nil {
		return
	}

	for i, r := range rs {
		cR := (*C.struct_samure_rect)(unsafe.Pointer(uintptr(unsafe.Pointer(cRs)) + unsafe.Sizeof(*cRs)*uintptr(i)))
		cR.x = C.int32_t(r.X)
		cR.y = C.int32_t(r.Y)
		cR.w = C.int32_t(r.W)
		cR.h = C.int32_t(r.H)
	}

	C.samure_output_set_input_regions(ctx.Handle, o.Handle, cRs, C.size_t(len(rs)))
	C.free(unsafe.Pointer(cRs))
}

func (o Output) SetKeyboardInteraction(enable bool) {
	var cEnable C.int
	if enable {
		cEnable = 1
	} else {
		cEnable = 0
	}
	C.samure_output_set_keyboard_interaction(o.Handle, cEnable)
}

func (o Output) AttachLayerSurface(sfc LayerSurface) {
	C.samure_output_attach_layer_surface(o.Handle, sfc.Handle)
}

func (o Output) Screenshot(ctx Context) (SharedBuffer, error) {
	buf_rs := C.samure_output_screenshot(ctx.Handle, o.Handle)
	if buf_rs.error != ErrorNone {
		return SharedBuffer{}, NewError(uint64(buf_rs.error))
	}

	return SharedBuffer{buf_rs.result}, nil
}

func (o Output) LenSurfaces() int {
	return int(o.Handle.num_sfc)
}

func (o Output) Surface(idx int) LayerSurface {
	return LayerSurface{
		(*C.struct_samure_layer_surface)(unsafe.Pointer(uintptr(unsafe.Pointer(o.Handle.sfc)) + unsafe.Sizeof(*o.Handle.sfc)*uintptr(idx))),
	}
}

func (o Output) Name() string {
	return C.GoString(o.Handle.name)
}

func (o Output) Geo() Rect {
	return Rect{
		int(o.Handle.geo.x),
		int(o.Handle.geo.y),
		int(o.Handle.geo.w),
		int(o.Handle.geo.h),
	}
}
