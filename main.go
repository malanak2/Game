package main

import (
	"flag"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/malanak2/Game/App/types"
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}
func main() {
	profile := flag.Bool("profile", false, "Enable cpu profiling")
	flag.Parse()
	if *profile {
		f, _ := os.Create("cpu.prof")
		defer f.Close()

		// Start profiling
		if err := pprof.StartCPUProfile(f); err != nil {
			panic(err)
		}
		defer pprof.StopCPUProfile()
	}

	pathPtr := flag.String("path", "./config.ini", "The path to program config")
	// Initialize engine
	err := types.InitApp(pathPtr)
	if err != nil {
		panic(err)
	}

	// Run app
	types.Run()
}
