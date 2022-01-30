package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Modify the echo program to also print os.Args[0], the name of the command that invoked it.
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(strings.Join(os.Args[:], " "))

	// Modify the echo program to print the index and value of each of its arguments, one per line.
	for i, v := range os.Args[:] {
		fmt.Printf("Test：%s\n", i+1, v)
	}
}
