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
