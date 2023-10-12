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

package cairo

import (
	"math"
	"testing"

	samure "github.com/PucklaJ/samurai-render-go"
	"github.com/gotk3/gotk3/cairo"
	"github.com/stretchr/testify/assert"
)

type TestCairoApp struct {
	elapsedTime float64
}

func (a *TestCairoApp) OnEvent(ctx samure.Context, event interface{}) {

}

func (a *TestCairoApp) OnRender(ctx samure.Context, layerSurface samure.LayerSurface, outputGeo samure.Rect, deltaTime float64) {
	c := Get(layerSurface)

	c.SetOperator(cairo.OPERATOR_SOURCE)
	c.SetSourceRGBA(1.0, 0.0, 0.0, 0.1)
	c.Paint()

	c.SetSourceRGBA(0.0, 1.0, 0.0, 1.0)
	c.Arc(float64(outputGeo.W/2), float64(outputGeo.H/2), 200.0, 0.0, 2.0*math.Pi)
	c.Fill()

	c.SetSourceRGBA(0.0, 0.0, 1.0, 1.0)
	c.Arc(float64(outputGeo.W/2), float64(outputGeo.H/2), 200.0, 0.0, 2.0*math.Pi)
	c.Stroke()
}

func (a *TestCairoApp) OnUpdate(ctx samure.Context, deltaTime float64) {
	a.elapsedTime += deltaTime
	if a.elapsedTime > 5.0 {
		ctx.SetRunning(false)
	}
}

func TestCairo(t *testing.T) {
	ctx, err := samure.CreateContextWithBackend(samure.CreateContextConfig(&TestCairoApp{}), &Backend{})
	if !assert.Nil(t, err) {
		return
	}
	defer ctx.Destroy()

	ctx.Run()
}
