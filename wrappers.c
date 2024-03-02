/***********************************************************************************
 *                         This file is part of samurai-render-go
 *                    https://github.com/Samudevv/samurai-render-go
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

#include "wrappers.h"
#include <assert.h>
#include <samure/context.h>
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
