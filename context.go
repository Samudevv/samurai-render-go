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
#include "samure/context.h"
#include "samure/backends/opengl.h"
#include "wrappers.h"
*/
import "C"
import (
	"unsafe"
)

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
	MaxUpdateFrequency           int
	NotCreateOutputLayerSurfaces bool
	NotRequestFrame              bool
	GL                           OpenGLConfig

	App App
}

func CreateContextConfig(a App) *ContextConfig {
	return &ContextConfig{
		GL:  DefaultOpenGLConfig(),
		App: a,
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
	c.max_update_frequency = C.uint32_t(cfg.MaxUpdateFrequency)
	if cfg.NotCreateOutputLayerSurfaces {
		c.not_create_output_layer_surfaces = 1
	}
	if cfg.NotRequestFrame {
		c.not_request_frame = 1
	}

	c.on_event = C.samure_event_callback(C.globalOnEvent)
	c.on_render = C.samure_render_callback(C.globalOnRender)
	c.on_update = C.samure_update_callback(C.globalOnUpdate)
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
	C.samure_context_set_render_state(ctx.Handle, uint32(v))
}

func (ctx Context) CreateOutputLayerSurfaces() error {
	err := C.samure_context_create_output_layer_surfaces(ctx.Handle)
	if err != ErrorNone {
		return NewError(uint64(err))
	}
	return nil
}

func (ctx Context) SetPointerInteraction(enable bool) {
	var cEnable C.int
	if enable {
		cEnable = 1
	}
	C.samure_context_set_pointer_interaction(ctx.Handle, cEnable)
}

func (ctx Context) SetInputRegions(rs []Rect) {
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

	C.samure_context_set_input_regions(ctx.Handle, cRs, C.size_t(len(rs)))
	C.free(unsafe.Pointer(cRs))
}

func (ctx Context) SetKeyboardInteraction(enable bool) {
	var cEnable C.int
	if enable {
		cEnable = 1
	}

	C.samure_context_set_keyboard_interaction(ctx.Handle, cEnable)
}

func (ctx Context) ProcessEvents() {
	C.samure_context_process_events(ctx.Handle)
}

func (ctx Context) RenderLayerSurface(sfc LayerSurface, o Rect) {
	C.samure_context_render_layer_surface(ctx.Handle, sfc.Handle, o.convertToC())
}

func (ctx Context) RenderOutput(o Output) {
	C.samure_context_render_output(ctx.Handle, o.Handle)
}

func (ctx Context) Update(deltaTime float64) {
	C.samure_context_update(ctx.Handle, C.double(deltaTime))
}

func (ctx Context) SetPointerShape(shape int) {
	C.samure_context_set_pointer_shape(ctx.Handle, C.uint32_t(shape))
}

func (ctx Context) Flush() {
	C.wl_display_flush(ctx.Handle.display)
}

func (ctx Context) Display() unsafe.Pointer {
	return unsafe.Pointer(ctx.Handle.display)
}

func (ctx Context) Shm() unsafe.Pointer {
	return unsafe.Pointer(ctx.Handle.shm)
}

func (ctx Context) Compositor() unsafe.Pointer {
	return unsafe.Pointer(ctx.Handle.compositor)
}

func (ctx Context) LayerShell() unsafe.Pointer {
	return unsafe.Pointer(ctx.Handle.layer_shell)
}

func (ctx Context) OutputManager() unsafe.Pointer {
	return unsafe.Pointer(ctx.Handle.output_manager)
}

func (ctx Context) ScreencopyManager() unsafe.Pointer {
	return unsafe.Pointer(ctx.Handle.screencopy_manager)
}

func (ctx Context) GetOutputRect() Rect {
	r := C.samure_context_get_output_rect(ctx.Handle)
	return Rect{
		X: int(r.x),
		Y: int(r.y),
		W: int(r.w),
		H: int(r.h),
	}
}
