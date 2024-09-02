package main

import "fmt"

func getMyAge() int {
	return 29
}

func main() {
	switch age := getMyAge(); true {
	case age < 20:
		fmt.Println("Teen")
	case age >= 20 && age < 40:
		fmt.Println("adult")
	default:
		fmt.Println("Your age is ...", age)
	}
}