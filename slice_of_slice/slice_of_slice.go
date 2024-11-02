package main

import "fmt"

func main() {
	x := []string{"a", "b", "c", "d"}
	y := x[:2]
	fmt.Println(len(x), len(y), cap(x), cap(y))
	y = append(y, "z")
	y = append(y, "z")
	y = append(y, "z")

	fmt.Println("x:", x)
	fmt.Println("y:", y)
}
