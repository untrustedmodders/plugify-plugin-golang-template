package main

import "C"
import (
	"fmt"
	"github.com/untrustedmodders/go-plugify"
)

func init() {
	plugify.OnPluginStart(func() error {
		fmt.Println("Go: OnPluginStart")
		return nil
	})

	/*plugify.OnPluginUpdate(func(dt float32) error {
		fmt.Println("Go: OnPluginUpdate")
		return nil
	})*/

	plugify.OnPluginEnd(func() error {
		fmt.Println("Go: OnPluginEnd")
		return nil
	})
}

func main() {
}

func MakePrint(count int32, message string) {
	for i := int32(0); i < count; i++ {
		fmt.Println(message)
	}
}
