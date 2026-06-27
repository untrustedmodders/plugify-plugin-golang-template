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
	plugin = plugify.NewPlugin("golang_plugin", onPluginStart, nil, onPluginEnd)
}

// plugify:export MakePrint
func MakePrint(count int32, message string) {
	for i := int32(0); i < count; i++ {
		fmt.Println(message)
	}
}
