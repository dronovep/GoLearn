package main

import "fmt"

type Greeter interface {
	greet()
}

type AbstractGreeter struct {
	Greeter
	name string
}

func (o *AbstractGreeter) hello() {
	fmt.Printf("Hello, my friend!\n")
	o.greet()
	fmt.Printf("Goodbye\n")
}

type Dog struct {
	AbstractGreeter
}

func (o *Dog) greet() {
	fmt.Printf("Woof!Woof! My name is %s!\n", o.name)
}

func createDog(name string) *Dog {
	dog := new(Dog)
	dog.name = name
	dog.Greeter = dog
	return dog
}

func main() {
	var greeter *AbstractGreeter = &createDog("Pongo").AbstractGreeter
	greeter.hello()
}
