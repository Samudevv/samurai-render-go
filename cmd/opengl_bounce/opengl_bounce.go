package main

import (
	"fmt"
	"os"

	samure "github.com/Samudevv/samurai-render-go"
	samureGL "github.com/Samudevv/samurai-render-go/backends/opengl"
	"github.com/go-gl/gl/v4.1-core/gl"
)

type App struct {
}

func (a *App) OnEvent(ctx samure.Context, event interface{}) {
	switch event.(type) {
	case samure.EventPointerButton:
		ctx.SetRunning(false)
	}
}

func (a *App) OnRender(ctx samure.Context, layerSurface samure.LayerSurface, outputGeo samure.Rect) {
}

func (a *App) OnUpdate(ctx samure.Context, deltaTime float64) {
}

func main() {
	cfg := samure.CreateContextConfig(&App{})
	cfg.PointerInteraction = true
	cfg.GL.MajorVersion = 4
	cfg.GL.MinorVersion = 1

	bkd := &samureGL.Backend{}
	ctx, err := samure.CreateContextWithBackend(cfg, bkd)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer ctx.Destroy()

	if err := bkd.InitBindings(ctx, gl.Init); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL Version:", version)

	ctx.Run()
}
