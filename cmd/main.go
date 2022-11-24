package main

import "fmt"

func main() {
	todos := []string{
		"Do the laundry",
		"Learn golang",
		"Have a snack",
	}

	for i, todo := range todos {
		mark := "✗"

		if i%2 == 1 {
			mark = "✓"
		}

		fmt.Printf("%s %s\n", mark, todo)
	}
}
