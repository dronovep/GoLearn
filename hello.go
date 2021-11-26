package main

import "fmt"

func main() {

	family := []string{"mother", "father", "brother", "me"}
	mbuffer := make([]rune, 50, 100)
	famsize := len(family)
	i, j, k := 0, 0, 0
mcycle:
	if i < famsize {
		member := family[i]
		mlen := len(member)
	lcycle:
		if j < mlen {
			letter := member[j]
			mbuffer[k] = rune(letter)
			j++
			k++
			goto lcycle
		}
		j = 0
		k++
		mbuffer[k] = ' '
		k++

		i++
		goto mcycle
	}

	k++
	mbuffer[k] = 0

	fmt.Printf("Result string:\n%s", string(mbuffer))
}
