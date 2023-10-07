package samure

/*
#include "samure/context.h"
#include "samure/backends/opengl.h"
#include "wrappers.h"
*/
import "C"
import "unsafe"

const (
	RenderStateAlways = C.SAMURE_RENDER_STATE_ALWAYS
	RenderStateNone   = C.SAMURE_RENDER_STATE_NONE
	RenderStateOnce   = C.SAMURE_RENDER_STATE_ONCE
)

type Context struct {
	Handle *C.struct_samure_context
}

type ContextConfig struct {
	PointerInteraction           bool
	KeyboardInteraction          bool
	TouchInteraction             bool
	MaxFPS                       int
	NotCreateOutputLayerSurfaces bool
	GL                           OpenGLConfig

	App App
}

func CreateContextConfig(a App) *ContextConfig {
	return &ContextConfig{
		MaxFPS: 60,
		GL:     DefaultOpenGLConfig(),
		App:    a,
	}
}

type OpenGLConfig struct {
	RedSize      int
	GreenSize    int
	BlueSize     int
	AlphaSize    int
	Samples      int
	DepthSize    int
	MajorVersion int
	MinorVersion int
	ProfileMask  int
	Debug        int
	ColorSpace   int
	RenderBuffer int
}

func DefaultOpenGLConfig() OpenGLConfig {
	return OpenGLConfig{
		RedSize:      8,
		GreenSize:    8,
		BlueSize:     8,
		AlphaSize:    8,
		MajorVersion: 1,
		ProfileMask:  0x00000001,
		ColorSpace:   0x308A,
		RenderBuffer: 0x3084,
	}
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
	c.gl = cfg.GL.convertToC()

	return c
}

func (cfg OpenGLConfig) convertToC() *C.struct_samure_opengl_config {
	ptr := (*C.struct_samure_opengl_config)(C.malloc(C.size_t(unsafe.Sizeof(C.struct_samure_opengl_config{}))))
	if ptr == nil {
		return nil
	}
	ptr.red_size = C.int(cfg.RedSize)
	ptr.green_size = C.int(cfg.GreenSize)
	ptr.blue_size = C.int(cfg.BlueSize)
	ptr.alpha_size = C.int(cfg.AlphaSize)
	ptr.samples = C.int(cfg.Samples)
	ptr.depth_size = C.int(cfg.DepthSize)
	ptr.major_version = C.int(cfg.MajorVersion)
	ptr.minor_version = C.int(cfg.MinorVersion)
	ptr.profile_mask = C.int(cfg.ProfileMask)
	ptr.debug = C.int(cfg.Debug)
	ptr.color_space = C.int(cfg.ColorSpace)
	ptr.render_buffer = C.int(cfg.RenderBuffer)
	return ptr
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
		*(**C.struct_samure_output)(unsafe.Pointer(uintptr(unsafe.Pointer(ctx.Handle.outputs)) + unsafe.Sizeof(&C.struct_samure_output{})*uintptr(idx))),
	}
}

func (ctx Context) LenSeats() int {
	return int(ctx.Handle.num_seats)
}

func (ctx Context) Seat(idx int) Seat {
	return Seat{
		*(**C.struct_samure_seat)(unsafe.Pointer(uintptr(unsafe.Pointer(ctx.Handle.seats)) + unsafe.Sizeof(&C.struct_samure_seat{})*uintptr(idx))),
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

func (ctx Context) SetRenderState(v int) {
	ctx.Handle.render_state = uint32(v)
}

func (ctx Context) CreateOutputLayerSurfaces() error {
	err := C.samure_context_create_output_layer_surfaces(ctx.Handle)
	if err != ErrorNone {
		return NewError(uint64(err))
	}
	return nil
}
