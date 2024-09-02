package main

import "fmt"

func main() {
	for i := 5; i >= 1; i-- {
		for j := 0; j < i; j++ {
			fmt.Printf("%c",'*')
		}
		fmt.Println()
	}
}