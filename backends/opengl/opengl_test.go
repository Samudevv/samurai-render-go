/***********************************************************************************
 *                         This file is part of samurai-render-go
 *                    https://github.com/Samudevv/samurai-render-go
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

package opengl

import (
	"testing"

	samure "github.com/Samudevv/samurai-render-go"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/stretchr/testify/assert"
)

type TestApp struct {
}

func (a *TestApp) OnEvent(ctx samure.Context, event interface{}) {
	switch e := event.(type) {
	case samure.EventPointerButton:
		if e.Button == samure.ButtonRight && e.State == samure.StateReleased {
			ctx.SetRunning(false)
		}
	}
}

func (a *TestApp) OnRender(ctx samure.Context, layerSurface samure.LayerSurface, o samure.Rect) {
	gl.ClearColor(1.0, 0.0, 0.0, 0.5)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(0.0, float64(o.W), float64(o.H), 0.0, 0.01, 100.0)

	gl.Disable(gl.DEPTH_TEST)
	gl.Enable(gl.DEPTH_CLAMP)

	x := float32(o.W / 2)
	y := float32(o.H / 2)
	s := float32(200)

	gl.Begin(gl.QUADS)
	gl.Color3f(0.0, 1.0, 0.0)
	gl.Vertex2f(x-s/2.0, y-s/2.0)
	gl.Vertex2f(x+s/2.0, y-s/2.0)
	gl.Vertex2f(x+s/2.0, y+s/2.0)
	gl.Vertex2f(x-s/2.0, y+s/2.0)
	gl.End()
}

func (a *TestApp) OnUpdate(ctx samure.Context, deltaTime float64) {

}

func TestOpenGL(t *testing.T) {
	cfg := samure.CreateContextConfig(&TestApp{})
	cfg.PointerInteraction = true

	var bak Backend

	ctx, err := samure.CreateContextWithBackend(cfg, &bak)
	if !assert.Nil(t, err) {
		return
	}
	defer ctx.Destroy()

	if err := bak.InitBindings(ctx, gl.Init); !assert.Nil(t, err) {
		return
	}

	ctx.Run()
}
