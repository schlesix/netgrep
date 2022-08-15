package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Nothing to see here, please move along! ;-)")
	args := os.Args[1:]
	fmt.Println(args)
}
