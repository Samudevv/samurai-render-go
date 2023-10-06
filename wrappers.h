#pragma once

#include "samure/backend.h"
#include "samure/events.h"
#include "samure/output.h"

extern void globalOnEvent(struct samure_context *ctx,
                          struct samure_event *event, void *user_data);
extern void globalOnRender(struct samure_context *ctx,
                           struct samure_layer_surface *layer_surface,
                           struct samure_rect output_geo, double delta_time,
                           void *user_data);
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
