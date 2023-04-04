package main

import (
	"fmt"

	"github.com/ivan-gavran/go-exampleproject/morefunctions"
)

func alternativeIncrement(i int) int {
	a := i + 1
	if i > 5 {
		panic("i is too big")
	}
	return a
}

func main() {
	var a int
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			a = morefunctions.Increment(i)
		} else {
			a = alternativeIncrement(i)
		}

		fmt.Println(a)
	}
}
