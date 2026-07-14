package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var tasks []string
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()
		cmd := strings.ToLower(input)

		switch {
		case cmd == "quit":
			return

		case cmd == "list":
			for i, task := range tasks {
				fmt.Printf("%d. %s\n", i+1, task)
			}

		// the rest of the input after "add " is the task
		case strings.HasPrefix(cmd, "add "):
			task := input[4:] // slice off "add " (4 characters)
			tasks = append(tasks, task)
			fmt.Println("Added:", task)

		default:
			fmt.Println("Unknown command:", input)
		}

	}

}
