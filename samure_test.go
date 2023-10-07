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

func (a *TestApp) OnRender(ctx Context, layerSurface LayerSurface, outputGeo Rect, deltaTime float64) {

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
