package main

import (
	"fmt"
)

func foo() {
	fmt.Printf("Inside foo\n")
}

func bar() {
	fmt.Printf("Inside bar\n")
}

func finish() {
	fmt.Printf("Work finished\n")
}

func buyCar(carchan chan bool) {
	var options_chan = make(chan bool)
	fmt.Printf("Buying a car\n")

	for gate := false; !gate; gate = <-carchan {
	}
	go addDopOptions(options_chan)
	options_chan <- true

	fmt.Printf("Car is yours\n")
	carchan <- true
}

func addDopOptions(options_chan chan bool) {
	<-options_chan
	fmt.Printf("Adding dop options\n")
	fmt.Printf("Dop options added\n")
}

func selectEngine(enginechan chan bool) {
	fmt.Printf("Selecting engine\n")
	fmt.Printf("Engine selected\n")
}

const (
	motherHello  = iota
	fatherHello  = iota
	sisterHello  = iota
	brotherHello = iota
	numHello     = iota
)

var start = make(chan bool, numHello)
var stop = make(chan bool, numHello)

var tasks = [numHello]func(start <-chan bool, stop chan<- bool){
	motherHello: func(start <-chan bool, stop chan<- bool) {
		<-start
		fmt.Printf("Hello from mother!\n")
		stop <- true
	},
	fatherHello: func(start <-chan bool, stop chan<- bool) {
		<-start
		fmt.Printf("Hello from father!\n")
		stop <- true
	},
	sisterHello: func(start <-chan bool, stop chan<- bool) {
		<-start
		fmt.Printf("Hello from sister!\n")
		stop <- true
	},
	brotherHello: func(start <-chan bool, stop chan<- bool) {
		<-start
		fmt.Printf("Hello from brother!\n")
		stop <- true
	},
}

func main() {
	defer finish()
	fmt.Printf("Starting main\n")
	// простейший пример
	//go foo()
	//go bar()
	//
	//time.Sleep(2 * time.Second)	// без засыпания foo и bar так и не будут вызваны
	//
	//fmt.Printf("Ending main\n")

	// пример с вложенными горутинами
	//carchan := make(chan bool)
	//
	//go buyCar(carchan)
	//carchan<- true
	//time.Sleep(3 * time.Second)

	// пример с набором из нескольких горутин-рабочих
	for i := 0; i < numHello; i++ {
		go tasks[i](start, stop)
		start <- true
	}

	for i := 0; i < numHello; i++ {
		<-stop
	}
}

/**
По факту при использовании оператора go, выполнение новых горутин не начинается немедленно, а только пока
не заблокируется родительская горутина их создавшая. Порядок запуска горутин случаен
*/
