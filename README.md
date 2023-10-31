# SamuraiRender Go

Go bindings to [SamuraiRender](https://github.com/PucklaJ/samurai-render) which is an easy to use library to render into the layer shell of wayland compositors.

## Getting Started

```go
package main

import (
	"fmt"
	"os"

	samure "github.com/PucklaJ/samurai-render-go"
	samureCairo "github.com/PucklaJ/samurai-render-go/backends/cairo"
	"github.com/gotk3/gotk3/cairo"
)

type App struct {
}

func (*App) OnEvent(ctx samure.Context, event interface{}) {
}

func (*App) OnRender(ctx samure.Context, layerSurface samure.LayerSurface, o samure.Rect) {
	c := samureCairo.Get(layerSurface)
	c.SetOperator(cairo.OPERATOR_SOURCE)
	c.SetSourceRGBA(1.0, 1.0, 1.0, 0.2)
	c.Paint()
}

func (*App) OnUpdate(ctx samure.Context, deltaTime float64) {
}

func main() {
	cfg := samure.CreateContextConfig(&App{})

	ctx, err := samure.CreateContextWithBackend(cfg, &samureCairo.Backend{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create context: %v\n", err)
		os.Exit(1)
	}
	defer ctx.Destroy()

	ctx.Run()
}
```
