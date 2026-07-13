package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	for {
		fmt.Print("> ")
		fmt.Scanln(&input)

		if strings.ToLower(input) == "quit" {
			break
		}

	}
}
