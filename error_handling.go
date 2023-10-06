package samure

/*
#include "samure/error_handling.h"
*/
import "C"
import "unsafe"

const (
	ErrorNone                    = C.SAMURE_ERROR_NONE
	ErrorFailed                  = C.SAMURE_ERROR_FAILED
	ErrorNotImplemented          = C.SAMURE_ERROR_NOT_IMPLEMENTED
	ErrorDisplayConnect          = C.SAMURE_ERROR_DISPLAY_CONNECT
	ErrorNoOutputs               = C.SAMURE_ERROR_NO_OUTPUTS
	ErrorNoXDGOutputManager      = C.SAMURE_ERROR_NO_XDG_OUTPUT_MANAGER
	ErrorNoLayerShell            = C.SAMURE_ERROR_NO_LAYER_SHELL
	ErrorNoSHM                   = C.SAMURE_ERROR_NO_SHM
	ErrorNoCompositor            = C.SAMURE_ERROR_NO_COMPOSITOR
	ErrorNoCursorShapeManager    = C.SAMURE_ERROR_NO_CURSOR_SHAPE_MANAGER
	ErrorNoScreencopyManager     = C.SAMURE_ERROR_NO_SCREENCOPY_MANAGER
	ErrorBackendInit             = C.SAMURE_ERROR_BACKEND_INIT
	ErrorNoBackendSupport        = C.SAMURE_ERROR_NO_BACKEND_SUPPORT
	ErrorLayerSurfaceInit        = C.SAMURE_ERROR_LAYER_SURFACE_INIT
	ErrorMemory                  = C.SAMURE_ERROR_MEMORY
	ErrorSharedBufferInit        = C.SAMURE_ERROR_SHARED_BUFFER_INIT
	ErrorOpenGLLoadProc          = C.SAMURE_ERROR_OPENGL_LOAD_PROC
	ErrorOpenGLDisplayConnect    = C.SAMURE_ERROR_OPENGL_DISPLAY_CONNECT
	ErrorOpenGLInitialize        = C.SAMURE_ERROR_OPENGL_INITIALIZE
	ErrorOpenGLConfig            = C.SAMURE_ERROR_OPENGL_CONFIG
	ErrorOpenGLBindAPI           = C.SAMURE_ERROR_OPENGL_BIND_API
	ErrorOpenGLContextInit       = C.SAMURE_ERROR_OPENGL_CONTEXT_INIT
	ErrorOpenGLwl_egl_windowInit = C.SAMURE_ERROR_OPENGL_WL_EGL_WINDOW_INIT
	ErrorOpenGLSurfaceInit       = C.SAMURE_ERROR_OPENGL_SURFACE_INIT
	ErrorSharedBufferFdInit      = C.SAMURE_ERROR_SHARED_BUFFER_FD_INIT
	ErrorSharedBufferTruncate    = C.SAMURE_ERROR_SHARED_BUFFER_TRUNCATE
	ErrorSharedBufferMmap        = C.SAMURE_ERROR_SHARED_BUFFER_MMAP
	ErrorSharedBufferPoolInit    = C.SAMURE_ERROR_SHARED_BUFFER_POOL_INIT
	ErrorSharedBufferBufferInit  = C.SAMURE_ERROR_SHARED_BUFFER_BUFFER_INIT
	ErrorFrameInit               = C.SAMURE_ERROR_FRAME_INIT
	ErrorCairoSurfaceInit        = C.SAMURE_ERROR_CAIRO_SURFACE_INIT
	ErrorCairoInit               = C.SAMURE_ERROR_CAIRO_INIT
	ErrorSurfaceInit             = C.SAMURE_ERROR_SURFACE_INIT
	ErrorOutputInit              = C.SAMURE_ERROR_OUTPUT_INIT
)

type Error struct {
	errorCode C.samure_error
}

func NewError(errorCode C.samure_error) *Error {
	return &Error{
		errorCode: errorCode,
	}
}

func (err *Error) Error() string {
	cErr := C.samure_build_error_string(err.errorCode)
	goErr := C.GoString(cErr)
	C.free(unsafe.Pointer(cErr))
	return goErr
}
