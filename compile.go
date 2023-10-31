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

/*
#cgo pkg-config: wayland-client wayland-cursor

#include "samure/backend.c"
#include "samure/backends/raw.c"
#include "samure/callbacks.c"
#include "samure/context.c"
#include "samure/cursors.c"
#include "samure/frame_timer.c"
#include "samure/layer_surface.c"
#include "samure/output.c"
#include "samure/rect.c"
#include "samure/seat.c"
#include "samure/shared_memory.c"
#include "samure/wayland/cursor-shape.c"
#include "samure/wayland/tablet.c"
#include "samure/wayland/layer-shell.c"
#include "samure/wayland/screencopy.c"
#include "samure/wayland/xdg-output.c"
#include "samure/wayland/xdg-shell.c"
*/
import "C"
