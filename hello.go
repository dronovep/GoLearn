package main

import (
	"fmt"
)

func main() {
	p := Point{0, 0, StandardSetCoords}

	p.SetCoords(12, 24)

	fmt.Printf("p = %v\n", p)
	p.SetCoordsAlg(&p, 35, 84)
	fmt.Printf("p = %v\n", p)
}
