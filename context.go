package samure

/*
#include "samure/context.h"
#include "wrappers.h"
*/
import "C"
import "unsafe"

type Context struct {
	handle *C.struct_samure_context
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

	bakIdx := AddGlobalBackend(bak)
	wrapBak := C.create_wrapper_backend(C.int(bakIdx))

	ctx_rs := C.samure_create_context_with_backend(&c, &wrapBak.base)
	if ctx_rs.error != ErrorNone {
		return Context{}, NewError(ctx_rs.error)
	}

	return Context{ctx_rs.result}, nil
}

func (ctx Context) Destroy() {
	C.samure_destroy_context(ctx.handle)
}
