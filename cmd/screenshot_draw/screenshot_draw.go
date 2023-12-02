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

package main

import (
	"fmt"
	"math"
	"os"

	samure "github.com/Samudevv/samurai-render-go"
	samureCairo "github.com/Samudevv/samurai-render-go/backends/cairo"
	"github.com/gotk3/gotk3/cairo"
)

type ScreenshotDraw struct {
	x, y      float64
	isDrawing bool
}

func (s *ScreenshotDraw) OnEvent(ctx samure.Context, event interface{}) {
	switch e := event.(type) {
	case samure.EventPointerMotion:
		s.x = e.X + float64(e.Seat.PointerFocus().Output().Geo().X)
		s.y = e.Y + float64(e.Seat.PointerFocus().Output().Geo().Y)
	case samure.EventPointerButton:
		if e.Button == samure.ButtonLeft {
			s.isDrawing = e.State == samure.StatePressed
		} else if e.Button == samure.ButtonRight && e.State == samure.StateReleased {
			ctx.SetRunning(false)
		}
	case samure.EventPointerEnter:
		ctx.SetPointerShape(samure.CursorShapeDefault)
	}
}

func (s *ScreenshotDraw) OnRender(ctx samure.Context, layerSurface samure.LayerSurface, o samure.Rect) {
	c := samureCairo.Get(layerSurface)
	c.SetOperator(cairo.OPERATOR_SOURCE)

	if s.isDrawing && o.PointInOutput(int(s.x), int(s.y)) {
		c.SetSourceRGBA(1.0, 0.0, 0.0, 1.0)
		c.Arc(o.RelX(s.x), o.RelY(s.y), 20.0, 0.0, 2.0*math.Pi)
		c.Fill()
	}

	c.SelectFontFace("sans-serif", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_BOLD)
	c.SetFontSize(50)

	str := fmt.Sprintf("FPS: %.3f", 1.0/layerSurface.FrameDeltaTime())

	te := c.TextExtents(str)
	c.SetSourceRGBA(0.0, 0.0, 0.0, 0.0)
	c.Rectangle(20.0, 20.0, te.Width, te.Height)
	c.Fill()
	c.MoveTo(20.0, 20.0+te.Height)
	c.SetSourceRGBA(1.0, 1.0, 1.0, 1.0)
	c.ShowText(str)
}

func (s *ScreenshotDraw) OnUpdate(ctx samure.Context, deltaTime float64) {

}

func main() {
	cfg := samure.CreateContextConfig(&ScreenshotDraw{})
	cfg.NotCreateOutputLayerSurfaces = true
	cfg.PointerInteraction = true
	cfg.MaxUpdateFrequency = 60

	ctx, err := samure.CreateContextWithBackend(cfg, &samureCairo.Backend{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create context: %v\n", err)
		os.Exit(1)
	}
	defer ctx.Destroy()

	for i := 0; i < ctx.LenOutputs(); i++ {
		o := ctx.Output(i)
		shot, err := o.Screenshot(ctx, false)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to take screenshot of %s: %v\n", o.Name(), err)
			continue
		}
		defer shot.Destroy()

		bg, err := samure.CreateLayerSurface(ctx, &o, samure.LayerOverlay, samure.AnchorFill, false, false, false)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create layer surface of %s: %v\n", o.Name(), err)
			continue
		}

		bg.DrawBuffer(shot)
	}

	ctx.CreateOutputLayerSurfaces()

	ctx.Run()
}
