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

#pragma once

#include "samure/backend.h"
#include "samure/events.h"
#include "samure/output.h"

extern void globalOnEvent(struct samure_context *ctx,
                          struct samure_event *event, void *user_data);
extern void globalOnRender(struct samure_context *ctx,
                           struct samure_layer_surface *layer_surface,
                           struct samure_rect output_geo, void *user_data);
extern void globalOnUpdate(struct samure_context *ctx, double delta_time,
                           void *user_data);

struct wrapper_backend {
  struct samure_backend base;

  int go_backend;
};

extern struct wrapper_backend *create_wrapper_backend(int go_backend);

extern void wrapper_backend_on_layer_surface_configure(
    struct samure_context *ctx, struct samure_layer_surface *layer_surface,
    int32_t width, int32_t height);

extern void
wrapper_backend_render_start(struct samure_context *ctx,
                             struct samure_layer_surface *layer_surface);

extern void
wrapper_backend_render_end(struct samure_context *ctx,
                           struct samure_layer_surface *layer_surface);

extern void wrapper_backend_destroy(struct samure_context *ctx);

extern samure_error wrapper_backend_associate_layer_surface(
    struct samure_context *ctx, struct samure_layer_surface *layer_surface);

extern void wrapper_backend_unassociate_layer_surface(
    struct samure_context *ctx, struct samure_layer_surface *layer_surface);
