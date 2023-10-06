package samure

import (
	"sync"
	"unsafe"
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
	OnEvent(ctx Context, event unsafe.Pointer)
	OnRender(ctx Context, layerSurface unsafe.Pointer, outputGeo unsafe.Pointer, deltaTime float64)
	OnUpdate(ctx Context, deltaTime float64)
}
