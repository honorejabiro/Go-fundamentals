package main

import "fmt"

func main() {
	i := 5
	j := 10

	if i * j == 50 {
		fmt.Println("You are correct")
	}

	if i > j {
		fmt.Println("Here we go")
	} else {
		fmt.Println("Finish line")
	}
}