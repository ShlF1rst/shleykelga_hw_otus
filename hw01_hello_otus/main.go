package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

const stringToReverse = "Hello, OTUS!"

func main() {
	reversedString := stringutil.Reverse(stringToReverse)
	fmt.Printf("%s\n", reversedString)
}
