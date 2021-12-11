package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

var window *sdl.Window
var err error

var mutex chan bool
var gate bool = true

func getGate() bool {
	mutex <- true
	result := gate
	<-mutex
	return result
}

func openGate() {
	mutex <- true
	gate = false
	<-mutex
}

func nameCounterRoutine(start <-chan bool) {
	defer (func() {
		fmt.Printf("Выходим из nameCounterRoutine\n")
		<-start
	})()
	<-start
	for i := 0; getGate(); i++ {
		window.SetTitle(fmt.Sprintf("Seconds passed: %d", i))
		time.Sleep(1 * time.Second)
	}
	fmt.Printf("В конце nameCounterRoutine\n")
}

func main() {
	fmt.Printf("\nНачало работы\n")
	mutex = make(chan bool, 1)
	sdl.Init(sdl.INIT_EVERYTHING)

	window, err = sdl.CreateWindow("Go window", 500, 250, 800, 600, 0)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}

	window.Show()

	start := make(chan bool)
	//stop := make(chan bool, 1)
	//stop <- true

	go nameCounterRoutine(start)
	start <- true

	for event := sdl.WaitEvent(); event.GetType() != sdl.QUIT; event = sdl.WaitEvent() {

	}
	openGate()
	start <- false
	sdl.Quit()
	fmt.Printf("Работа завершена\n")
}
