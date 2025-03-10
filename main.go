package main

import "C"
import (
	"fmt"
	"github.com/untrustedmodders/go-plugify"
	"runtime/debug"
)

func init() {
	plugify.OnPluginStart(func() {
		fmt.Println("Go: OnPluginStart")
	})

	/*plugify.OnPluginUpdate(func(dt float32) {
		fmt.Println("Go: OnPluginUpdate")
	})*/

	plugify.OnPluginEnd(func() {
		fmt.Println("Go: OnPluginEnd")
	})

	plugify.OnPluginPanic(func() []byte {
		return debug.Stack() // workaround for could not import runtime/debug inside plugify package
	})
}

func main() {
}

func MakePrint(count int32, message string) {
	for i := int32(0); i < count; i++ {
		fmt.Println(message)
	}
}
