package main

import (
	"fmt"
	"os"
)

func read_file(file string) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("file can't be opened: ", err)
	}
	defer f.Close()
	data := make([]byte, 64)
	for {

		count, err := f.Read(data)
		if err != nil {
			fmt.Println(err)
			return
		}
		// fmt.Println(string(data[:count]))
		os.Stdout.Write(data[:count])
	}
}
func main() {
	if len(os.Args) < 2 {
		fmt.Println("No file specified, Please provide file name while running the program")
		return
	}
	file := os.Args[1]
	read_file(file)
}
