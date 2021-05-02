package engine

import (
	"github.com/veandco/go-sdl2/sdl"
)

type SceneContext struct {
	ScreenWidth  float64
	ScreenHeight float64
}

type SceneCreator func(ctx *SceneContext) Scene

type Scene interface {
	UpdateScene(ctx *SceneContext, timeDelta uint32) SceneCreator
	RenderScene(ctx *SceneContext, surface *sdl.Surface)
	HandleEvent(event sdl.Event)
}
