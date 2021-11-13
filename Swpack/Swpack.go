package Swpack

// #cgo LDFLAGS:  -l SDL2
// #include <SDL2/SDL.h>
import "C"

func Mfoo() string {
	return "mfoo"
}

func SDLInit() {
	C.SDL_Init()
}
