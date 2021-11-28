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

var d int = 1

func motherHello(mchan chan bool) {
	<-mchan
	fmt.Printf("Hello from mother!\n")
	d++
	if d == 4 {
		close(mchan)
	}
}

func fatherHello(mchan chan bool) {
	<-mchan
	fmt.Printf("Hello from father!\n")
	d++ // плохой вариант, возможно состояние гонки. все горутины могут одновременно попробовать увеличить d
	if d == 4 {
		close(mchan)
	}
}

func sisterHello(mchan chan bool) {
	<-mchan
	fmt.Printf("Hello from sister!\n")
	d++
	if d == 4 {
		close(mchan)
	}
}

func brotherHello(mchan chan bool) {
	<-mchan
	fmt.Printf("Hello from brother!\n")
	d++
	if d == 4 {
		close(mchan)
	}
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
	var mchan = make(chan bool, 4)
	go motherHello(mchan)
	go fatherHello(mchan)
	go sisterHello(mchan)
	go brotherHello(mchan)

	for i := 0; i < cap(mchan); i++ {
		mchan <- true
	}
	mchan <- true

	for result := range mchan {
		fmt.Printf("Result = %v\n", result)
	}
}

/**
По факту при использовании оператора go, выполнение новых горутин не начинается немедленно, а только пока
не заблокируется родительская горутина их создавшая. Порядок запуска горутин случаен
*/
