package mutils

import (
	"fmt"
	"octadev.ru/GoLearn/mtypes"
)

func TestScanInterface() {
	wiper := mtypes.Wiper{}
	fmt.Sscanf("vaka gashira wipeout", "%v", &wiper)
}
