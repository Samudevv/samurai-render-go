#include "wrappers.h"
#include "samure/context.h"
#include <assert.h>
#include <stdlib.h>

struct wrapper_backend *create_wrapper_backend(int go_backend) {
  struct wrapper_backend *bak = malloc(sizeof(struct wrapper_backend));
  assert(bak != NULL);

  bak->go_backend = go_backend;

  bak->base.on_layer_surface_configure =
      wrapper_backend_on_layer_surface_configure;
  bak->base.render_start = wrapper_backend_render_start;
  bak->base.render_end = wrapper_backend_render_end;
  bak->base.destroy = wrapper_backend_destroy;
  bak->base.associate_layer_surface = wrapper_backend_associate_layer_surface;
  bak->base.unassociate_layer_surface =
      wrapper_backend_unassociate_layer_surface;

  return bak;
}
