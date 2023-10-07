package samure

/*
#include "samure/context.h"
#include "wrappers.h"
*/
import "C"
import "unsafe"

type Context struct {
	Handle *C.struct_samure_context
}

type ContextConfig struct {
	PointerInteraction           bool
	KeyboardInteraction          bool
	TouchInteraction             bool
	MaxFPS                       int
	NotCreateOutputLayerSurfaces bool

	App App
}

func (cfg ContextConfig) convertToC() C.struct_samure_context_config {
	var c C.struct_samure_context_config
	c.backend = C.SAMURE_BACKEND_NONE
	if cfg.PointerInteraction {
		c.pointer_interaction = 1
	}
	if cfg.KeyboardInteraction {
		c.keyboard_interaction = 1
	}
	if cfg.TouchInteraction {
		c.touch_interaction = 1
	}
	c.max_fps = C.uint32_t(cfg.MaxFPS)
	if cfg.NotCreateOutputLayerSurfaces {
		c.not_create_output_layer_surfaces = 1
	}

	c.event_callback = C.samure_event_callback(C.globalOnEvent)
	c.render_callback = C.samure_render_callback(C.globalOnRender)
	c.update_callback = C.samure_update_callback(C.globalOnUpdate)
	c.user_data = unsafe.Pointer(uintptr(AddGlobalApp(cfg.App)))

	return c
}

func CreateContextWithBackend(cfg *ContextConfig, bak Backend) (Context, error) {
	var c C.struct_samure_context_config
	if cfg != nil {
		c = cfg.convertToC()
	} else {
		c = C.samure_default_context_config()
		c.backend = C.SAMURE_BACKEND_NONE
	}

	not_create_output_layer_surfaces := c.not_create_output_layer_surfaces
	c.not_create_output_layer_surfaces = 1

	ctx_rs := C.samure_create_context(&c)
	if ctx_rs.error != ErrorNone {
		return Context{}, NewError(uint64(ctx_rs.error))
	}

	ctx := Context{ctx_rs.result}

	if err := bak.Init(ctx); err != nil {
		ctx.Destroy()
		return Context{nil}, err
	}

	bakIdx := AddGlobalBackend(bak)
	wrapBak := C.create_wrapper_backend(C.int(bakIdx))

	ctx.Handle.backend = &wrapBak.base

	if not_create_output_layer_surfaces == 0 {
		ctx.Handle.config.not_create_output_layer_surfaces = 0
		if err := C.samure_context_create_output_layer_surfaces(ctx.Handle); err != ErrorNone {
			ctx.Destroy()
			return Context{nil}, NewError(uint64(err))
		}
	}

	return Context{ctx_rs.result}, nil
}

func (ctx Context) Destroy() {
	C.samure_destroy_context(ctx.Handle)
}

func (ctx Context) LenOutputs() int {
	return int(ctx.Handle.num_outputs)
}

func (ctx Context) Output(idx int) Output {
	return Output{
		(*C.struct_samure_output)(unsafe.Pointer(uintptr(unsafe.Pointer(ctx.Handle.outputs)) + unsafe.Sizeof(*ctx.Handle.outputs)*uintptr(idx))),
	}
}

func (ctx Context) LenSeats() int {
	return int(ctx.Handle.num_seats)
}

func (ctx Context) Seat(idx int) Seat {
	return Seat{
		(*C.struct_samure_seat)(unsafe.Pointer(uintptr(unsafe.Pointer(ctx.Handle.seats)) + unsafe.Sizeof(*ctx.Handle.seats)*uintptr(idx))),
	}
}

func (ctx Context) Run() {
	C.samure_context_run(ctx.Handle)
}

func (ctx Context) SetRunning(v bool) {
	if v {
		ctx.Handle.running = 1
	} else {
		ctx.Handle.running = 0
	}
}
