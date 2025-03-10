package main

// #include "autoexports.h"
import "C"
import (
	"github.com/untrustedmodders/go-plugify"
	"unsafe"
)

// Exported methods

//export __MakePrint
func __MakePrint(count int32, message *C.String) {
	MakePrint(count, plugify.GetStringData((*plugify.PlgString)(unsafe.Pointer(message))))
}
