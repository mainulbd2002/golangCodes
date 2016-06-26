package main

import "fmt"

func main() {

	for i := 0; i < 201; i++ {
		fmt.Printf("%d - %b - %o - %x\n", i, i, i, i)
	}

}
