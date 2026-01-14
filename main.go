package main

import (
	"Game/App/types"
	"flag"
)

func main() {
	pathPtr := flag.String("path", "./config.ini", "The path to program config")
	// Initialize engine
	app := types.InitApp(pathPtr)
	// Initialize Renderer
	app.
		// Run app
		app.Run()
}
