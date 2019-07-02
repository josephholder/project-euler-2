package main

import (
	"fmt"
)

func main() {

	sum := 0

	beforeLastInSequence := 1
	lastInSequence := 2
	sum += 2 // adding the first even number

	for lastInSequence < 4000000 {

		value := beforeLastInSequence + lastInSequence
		beforeLastInSequence = lastInSequence
		lastInSequence = value

		if value%2 == 0 {
			sum += value
		}
	}

	fmt.Println(sum)
}
