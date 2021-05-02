package game

import (
	"github.com/daviwil/pongo/engine"
	"github.com/daviwil/pongo/entities"
	"github.com/veandco/go-sdl2/sdl"
)

type GameScene struct {
	paddleLeft  *paddle.Paddle
	paddleRight *paddle.Paddle
}

func CreateGameScene(ctx *engine.SceneContext) engine.Scene {
	screenWidthPadding := ctx.ScreenWidth / 12.0
	gameScene := &GameScene{
		paddleLeft:  paddle.MakePaddle(screenWidthPadding, ctx.ScreenHeight/3, 0xff0000ff),
		paddleRight: paddle.MakePaddle(ctx.ScreenWidth-(paddle.PaddleWidth+screenWidthPadding), ctx.ScreenHeight/3, 0xffff0000),
	}

	return gameScene
}

func (game *GameScene) UpdateScene(ctx *engine.SceneContext, timeDelta uint32) engine.SceneCreator {
	paddle.UpdatePaddle(game.paddleLeft, timeDelta, ctx.ScreenHeight)
	paddle.UpdatePaddle(game.paddleRight, timeDelta, ctx.ScreenHeight)

	return nil
}

func (game *GameScene) RenderScene(ctx *engine.SceneContext, surface *sdl.Surface) {
	paddle.RenderPaddle(game.paddleLeft, surface)
	paddle.RenderPaddle(game.paddleRight, surface)
}

func (game *GameScene) HandleEvent(event sdl.Event) {
	switch t := event.(type) {
	case *sdl.KeyboardEvent:
		if t.Repeat == 0 {
			// TODO: Can only hold down 2 keys at once???
			switch t.Keysym.Scancode {
			case sdl.SCANCODE_E:
				paddle.HandlePaddleKey(game.paddleLeft, paddle.PaddleKeyUp, paddle.GetKeyState(t.Type))
				break
			case sdl.SCANCODE_D:
				paddle.HandlePaddleKey(game.paddleLeft, paddle.PaddleKeyDown, paddle.GetKeyState(t.Type))
				break
			case sdl.SCANCODE_I:
				paddle.HandlePaddleKey(game.paddleRight, paddle.PaddleKeyUp, paddle.GetKeyState(t.Type))
				break
			case sdl.SCANCODE_K:
				paddle.HandlePaddleKey(game.paddleRight, paddle.PaddleKeyDown, paddle.GetKeyState(t.Type))
				break
			}
		}
	}
}
