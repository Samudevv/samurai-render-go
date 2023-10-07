package main

import (
	"fmt"
	"math"
	"os"

	samure "github.com/PucklaJ/samurai-render-go"
	samureCairo "github.com/PucklaJ/samurai-render-go/backends/cairo"
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
	}
}

func (s *ScreenshotDraw) OnRender(ctx samure.Context, layerSurface samure.LayerSurface, o samure.Rect, deltaTime float64) {
	c := samureCairo.Get(layerSurface)
	c.SetOperator(cairo.OPERATOR_SOURCE)

	if s.isDrawing && o.PointInOutput(int(s.x), int(s.y)) {
		c.SetSourceRGBA(1.0, 0.0, 0.0, 1.0)
		c.Arc(o.RelX(s.x), o.RelY(s.y), 20.0, 0.0, 2.0*math.Pi)
		c.Fill()
	}

	c.SelectFontFace("sans-serif", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_BOLD)
	c.SetFontSize(50)

	str := fmt.Sprint("FPS: ", int(1.0/deltaTime))

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
	cfg.MaxFPS = 60

	ctx, err := samure.CreateContextWithBackend(cfg, &samureCairo.Backend{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create context: %v\n", err)
		os.Exit(1)
	}
	defer ctx.Destroy()

	for i := 0; i < ctx.LenOutputs(); i++ {
		o := ctx.Output(i)
		shot, err := o.Screenshot(ctx)
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
