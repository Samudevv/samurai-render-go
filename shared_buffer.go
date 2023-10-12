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
#include "samure/shared_memory.h"
*/
import "C"
import "unsafe"

type SharedBuffer struct {
	Handle *C.struct_samure_shared_buffer
}

func (b SharedBuffer) Data() unsafe.Pointer {
	return b.Handle.data
}

func (b SharedBuffer) Fd() int {
	return int(b.Handle.fd)
}

func (b SharedBuffer) Width() int {
	return int(b.Handle.width)
}

func (b SharedBuffer) Height() int {
	return int(b.Handle.height)
}

func (b SharedBuffer) Format() int {
	return int(b.Handle.format)
}

func (b SharedBuffer) Destroy() {
	C.samure_destroy_shared_buffer(b.Handle)
}

func (b SharedBuffer) Copy(src SharedBuffer) error {
	err := C.samure_shared_buffer_copy(b.Handle, src.Handle)
	if err != ErrorNone {
		return NewError(uint64(err))
	}
	return nil
}
