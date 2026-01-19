package main

import (
	"Game/App/types"
	"flag"
	"runtime"
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}
func main() {
	pathPtr := flag.String("path", "./config.ini", "The path to program config")
	// Initialize engine
	app, err := types.InitApp(pathPtr)
	if err != nil {
		panic(err)
	}

	// Run app
	app.Run()
}
