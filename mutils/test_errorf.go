package mutils

import (
	"fmt"
	. "octadev.ru/GoLearn/mtypes"
)

func mcatch() {
	merror := recover().(*TestError)
	message := fmt.Errorf("Ловим возникшую ошибку - %w\n", merror)
	fmt.Printf("mcatch: error message: %s\n", message.Error())
}

func TestErrorf() {
	defer mcatch()
	panic(&TestError{})
}
