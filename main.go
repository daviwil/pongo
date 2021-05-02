package main

import (
	"github.com/daviwil/pongo/engine"
	"github.com/daviwil/pongo/scenes"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	// Initialize SDL2
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	screenWidth := 800.0
	screenHeight := 600.0

	window, err :=
		sdl.CreateWindow(
			"PonGo",
			sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
			int32(screenWidth), int32(screenHeight), sdl.WINDOW_SHOWN,
		)

	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	sceneContext := &engine.SceneContext{
		ScreenWidth:  screenWidth,
		ScreenHeight: screenHeight,
	}

	// Create the initial scene
	var currentScene engine.Scene = game.CreateGameScene(sceneContext)

	// Store the last update time for use later
	lastUpdateTime := sdl.GetTicks()

	running := true
	var nextSceneCreator engine.SceneCreator = nil
	for running {
		// Check the event loop
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quitting...")
				running = false
				break

			default:
				currentScene.HandleEvent(event)
			}
		}

		// Update the scene
		currentTime := sdl.GetTicks()
		currentDelta := currentTime - lastUpdateTime
		nextSceneCreator = currentScene.UpdateScene(sceneContext, currentDelta)

		surface, err := window.GetSurface()
		if err != nil {
			panic(err)
		}

		// Clear the screen and draw the current scene
		surface.FillRect(nil, 0)
		currentScene.RenderScene(sceneContext, surface)
		window.UpdateSurface()

		// Store the last update time
		lastUpdateTime = currentTime

		// Change to a new scene?
		if nextSceneCreator != nil {
			// Create the next scene
			currentScene = nextSceneCreator(sceneContext)

			if currentScene == nil {
				println("Quitting due to nil scene")
				running = false
			}
		}
	}
}
