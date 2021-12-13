package mtypes

import (
	"fmt"
	"io"
)

type Wiper struct {
	Value string
}

func (p *Wiper) Scan(state fmt.ScanState, verb rune) error {
	for runa, _, err := state.ReadRune(); err != io.EOF; runa, _, err = state.ReadRune() {
		fmt.Printf("%c\n", runa)
	}
	return nil
}
