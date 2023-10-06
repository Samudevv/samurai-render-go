package samure

/*
#include "samure/output.h"
*/
import "C"

type Output struct {
	Handle *C.struct_samure_output
}

type Rect struct {
	X, Y, W, H int
}
