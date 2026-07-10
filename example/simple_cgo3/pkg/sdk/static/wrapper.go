package static

/*
#cgo CFLAGS: -I${SRCDIR}/../../../include
#cgo LDFLAGS: -L${SRCDIR}/../../../lib/static -lsdk
#include <stdlib.h>
#include <api.h>
*/
import "C"
import "unsafe"

func Add(a, b int32) int32 {
	return int32(C.add(C.int(a), C.int(b)))
}

func Greet(name string) {
	nameCStr := C.CString(name)
	defer C.free(unsafe.Pointer(nameCStr))
	C.greet(nameCStr)
}
