package mutils

import (
	"fmt"
	. "octadev.ru/GoLearn/mtypes"
)

func errorfCatch() {
	merror := recover().(*TestError)
	message := fmt.Errorf("Ловим возникшую ошибку - %w\n", merror)
	fmt.Printf("fprintf_catch: error message: %s\n", message.Error())
}

func TestErrorf() {
	defer errorfCatch()
	panic(&TestError{})
}
