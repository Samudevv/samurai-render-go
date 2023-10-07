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
