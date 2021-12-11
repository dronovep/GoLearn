package parallel

import "fmt"

type MProcess struct {
	mutex      chan bool
	Control    chan bool
	gate       bool
	routine    func(process *MProcess)
	needToWait bool
}

func (o *MProcess) run() {
	defer (func() {
		fmt.Printf("Выходим из run\n")
	})()
	<-o.Control
	o.routine(o)
	if o.GetNeedToWait() {
		o.Control <- true
	}
	fmt.Printf("В конце run\n")
}

func (o *MProcess) Start() {
	go o.run()
	o.Control <- true
}

func (o *MProcess) Stop() {
	o.mutex <- true
	o.gate = false
	<-o.mutex
}

func (o *MProcess) WaitForStop() {
	o.mutex <- true
	o.gate = false
	o.needToWait = true
	<-o.mutex
	<-o.Control
}

func (o *MProcess) GetGate() bool {
	o.mutex <- true
	result := o.gate
	<-o.mutex

	return result
}

func (o *MProcess) GetNeedToWait() bool {
	o.mutex <- true
	result := o.needToWait
	<-o.mutex

	return result
}

func CreateMProcess(routine func(process *MProcess)) *MProcess {
	process := new(MProcess)
	process.mutex = make(chan bool, 1)
	process.Control = make(chan bool)
	process.gate = true
	process.routine = routine
	process.needToWait = false

	return process
}
