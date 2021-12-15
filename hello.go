package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	vk "github.com/vulkan-go/vulkan"
	. "octadev.ru/GoLearn/parallel"
)

var window *sdl.Window
var err error

func nameCounterRoutine(process *MProcess) {
	for i := 0; process.GetGate(); i++ {
		window.SetTitle(fmt.Sprintf("Seconds passed: %d", i))
		//time.Sleep(1 * time.Second)
	}
}

func main() {
	fmt.Printf("\nНачало работы\n")
	sdl.Init(sdl.INIT_EVERYTHING)

	window, err = sdl.CreateWindow("Go window", 500, 250, 800, 600, 0)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}

	window.Show()

	err := vk.SetDefaultGetInstanceProcAddr()
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	err = vk.Init()
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}

	for event := sdl.WaitEvent(); event.GetType() != sdl.QUIT; event = sdl.WaitEvent() {

	}

	sdl.Quit()
	fmt.Printf("Работа завершена\n")
}
