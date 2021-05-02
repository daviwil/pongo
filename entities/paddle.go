package paddle

import (
	"github.com/veandco/go-sdl2/sdl"
)

const PaddleWidth float64 = 35.0
const PaddleHeight float64 = 150.0
const PaddleMoveSpeed float64 = 400.0 // pixels per second

type Paddle struct {
	position  float64
	direction float64
	color     uint32
	rect      sdl.Rect
}

func UpdatePaddle(paddle *Paddle, timeDelta uint32, screenHeight float64) {
	paddle.position += paddle.direction * PaddleMoveSpeed * float64(timeDelta) / 1000.0
	if paddle.position < 20.0 {
		paddle.position = 20.0
	} else if (paddle.position + PaddleHeight) > (float64(screenHeight) - 20.0) {
		paddle.position = float64(screenHeight) - 20.0 - PaddleHeight
	}

	paddle.rect.Y = int32(paddle.position)
}

func RenderPaddle(paddle *Paddle, surface *sdl.Surface) {
	surface.FillRect(&paddle.rect, paddle.color)
}

func MakePaddle(posX float64, posY float64, color uint32) *Paddle {
	return &Paddle{
		position: posY,
		color:    color,
		rect: sdl.Rect{
			X: int32(posX),
			Y: int32(posY),
			W: int32(PaddleWidth),
			H: int32(PaddleHeight),
		},
	}
}

type PaddleKey uint8

const (
	PaddleKeyNone PaddleKey = iota
	PaddleKeyUp
	PaddleKeyDown
)

type PaddleKeyState uint8

const (
	PaddleKeyReleased PaddleKeyState = iota
	PaddleKeyPressed
)

func HandlePaddleKey(paddle *Paddle, key PaddleKey, keyState PaddleKeyState) {
	switch key {
	case PaddleKeyUp:
		if keyState == PaddleKeyPressed {
			paddle.direction = -1
		} else {
			if paddle.direction == -1 {
				paddle.direction = 0
			}
		}
		break
	case PaddleKeyDown:
		if keyState == PaddleKeyPressed {
			paddle.direction = 1
		} else {
			if paddle.direction == 1 {
				paddle.direction = 0
			}
		}
		break
	}
}

func GetKeyState(eventType uint32) PaddleKeyState {
	if int(eventType) == sdl.KEYDOWN {
		return PaddleKeyPressed
	}

	return PaddleKeyReleased
}
