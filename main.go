package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	Done  bool
	Title string
}

func main() {

	var tasks []Task
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
			for i, t := range tasks {
				check := "[ ]"
				if t.Done {
					check = "[x]"
				}
				fmt.Printf("%d. %s %s\n", i+1, check, t.Title)
			}

		// the rest of the input after "add " is the task
		case strings.HasPrefix(cmd, "add "):
			title := input[4:] // slice off "add " (4 characters)
			newTask := Task{Title: title, Done: false}
			tasks = append(tasks, newTask)
			fmt.Println("Added:", title)

		default:
			fmt.Println("Unknown command:", input)
		}

	}

}
