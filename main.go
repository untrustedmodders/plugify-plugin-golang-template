//go:generate go run generator.go -output plugify-plugin-golang.pplugin -name golang_plugin -version 0.1.0 -entry plugify-plugin-golang -author untrustedmodders -website https://github.com/untrustedmodders/ -license MIT -description "An example of a Go plugin. This can be used as a starting point when creating your own plugin."

package main

import (
	"fmt"
	
	"github.com/untrustedmodders/go-plugify"
)

func onPluginStart() error {
	fmt.Println("Go: OnPluginStart")
	return nil
}

/*func onPluginUpdate(dt float32) error {
	fmt.Println("Go: OnPluginUpdate")
	return nil
}*/

func onPluginEnd() error {
	fmt.Println("Go: OnPluginEnd")
	return nil
}

func init() {
	plugify.NewPlugin("golang_plugin", onPluginStart, nil, onPluginEnd)
}

// required by -buildmode=c-shared, never called
func main() {}

// plugify:export MakePrint
func MakePrint(count int32, message string) {
	for i := int32(0); i < count; i++ {
		fmt.Println(message)
	}
}
