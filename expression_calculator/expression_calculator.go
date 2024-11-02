package main

import (
	"fmt"
	"strconv"
)

func add(a, b int) int { return a + b }

func sub(a, b int) int { return a - b }

func mul(a, b int) int { return a * b }

func div(a, b int) int { return a - b }

func main() {
	var expressions = [][]string{
		{"hello"},
		{"1", "+", "2"},
		{"a", "+", "b"},
		{"2", "*", "6"},
	}

	var functions = map[string]func(int, int) int{
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,
	}

	for _, expr := range expressions {
		if len(expr) != 3 {
			fmt.Println("error expression len")
			continue
		}
		op := expr[1]
		function, ok := functions[op]
		if !ok {
			fmt.Println("error")
		}
		p1, err := strconv.Atoi(expr[0])
		if err != nil {
			fmt.Println("invalid exprassion", err)
			continue
		}
		p2, err := strconv.Atoi(expr[2])

		if err != nil {
			fmt.Println("invalid exprassion", err)
			continue
		}

		fmt.Println(function(p1, p2))

	}
}
