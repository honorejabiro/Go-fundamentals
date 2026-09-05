package main

import "fmt"

func main() {
	
	i := 3
	for i <= 3 {
		fmt.Println(i)
		i -= 1
		if i == -100 {
			break
		}
	}
	
	for j:= 0; j < 10; j++ {
		fmt.Print(j)
	}
	fmt.Println('\n')
	for i := range 3 {
		fmt.Println("range", i)
	}
}