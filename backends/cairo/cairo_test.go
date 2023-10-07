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
	var cfg samure.ContextConfig
	cfg.App = &TestCairoApp{}
	cfg.MaxFPS = 60

	ctx, err := samure.CreateContextWithBackend(&cfg, &Backend{})
	if !assert.Nil(t, err) {
		return
	}
	defer ctx.Destroy()

	ctx.Run()
}
