package sdk

/*
#cgo CFLAGS: -I${SRCDIR}/../../include
#cgo LDFLAGS: -L${SRCDIR}/../../lib -lsdk -lstdc++
#include <stdlib.h>
#include <wrapper.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type Sdk struct {
	h C.CalculatorHandle
}

func NewSdk() *Sdk {
	return &Sdk{
		h: C.Calculator_new(),
	}
}

func (sdk *Sdk) Add(a, b int32) int32 {
	if sdk.h == nil {
		panic("Sdk not initialized Or Free")
	}
	return int32(C.Calculator_add(sdk.h, C.int(a), C.int(b)))
}

func (sdk *Sdk) Greet(name string) {
	if sdk.h == nil {
		panic("Sdk not initialized Or Free")
	}
	nameCStr := C.CString(name)
	defer C.free(unsafe.Pointer(nameCStr))
	C.Calculator_greet(sdk.h, nameCStr)
}

func (sdk *Sdk) Free() {
	C.Calculator_delete(sdk.h)
	sdk.h = nil
	fmt.Println("Sdk Free")
}
