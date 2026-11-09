package main

//#include "autoexports.h"
import "C"
import (
	"reflect"
	"unsafe"

	"github.com/untrustedmodders/go-plugify"
)

var _ = reflect.TypeOf(0)
var _ = unsafe.Sizeof(0)
var _ = plugify.ApiVersion

//export __MakePrint
func __MakePrint(count int32, message *C.String) {
	MakePrint(count, plugify.GetStringData[string]((*plugify.PlgString)(unsafe.Pointer(message))))
}
