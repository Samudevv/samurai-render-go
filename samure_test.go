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

package samure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestApp struct {
}

func (a *TestApp) OnEvent(ctx Context, event interface{}) {
	switch e := event.(type) {
	case EventPointerButton:
		if e.Button == ButtonRight && e.State == StateReleased {
			ctx.SetRunning(false)
		}
	}
}

func (a *TestApp) OnRender(ctx Context, layerSurface LayerSurface, outputGeo Rect) {

}

func (a *TestApp) OnUpdate(ctx Context, deltaTime float64) {

}

func TestSamure(t *testing.T) {
	cfg := CreateContextConfig(&TestApp{})
	cfg.PointerInteraction = true

	ctx, err := CreateContextWithBackend(
		cfg,
		&RawBackend{},
	)
	if !assert.Nil(t, err) {
		return
	}
	defer ctx.Destroy()

	ctx.Run()
}
