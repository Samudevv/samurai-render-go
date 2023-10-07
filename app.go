package samure

import (
	"sync"
)

var globalApps map[int]App
var globalAppsMtx sync.Mutex

func AddGlobalApp(a App) int {
	globalAppsMtx.Lock()
	defer globalAppsMtx.Unlock()

	if globalApps == nil {
		globalApps = make(map[int]App)
	}

	idx := len(globalApps)
	globalApps[idx] = a
	return idx
}

func GetGlobalApp(idx int) App {
	globalAppsMtx.Lock()
	defer globalAppsMtx.Unlock()

	return globalApps[idx]
}

type App interface {
	OnEvent(ctx Context, event interface{})
	OnRender(ctx Context, layerSurface LayerSurface, outputGeo Rect, deltaTime float64)
	OnUpdate(ctx Context, deltaTime float64)
}
