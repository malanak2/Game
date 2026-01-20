package main

import (
	"flag"
	"runtime"

	"github.com/malanak2/Game/App/types"
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}
func main() {
	pathPtr := flag.String("path", "./config.ini", "The path to program config")
	// Initialize engine
	err := types.InitApp(pathPtr)
	if err != nil {
		panic(err)
	}

	// Run app
	types.Run()
}
