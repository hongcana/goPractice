package main

import "fmt"

func Multiple(a int, b int) int {
	return a * b
}

func main() {
	var a int
	var b int

	n, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println(n, err)
		return
	}
	c := Multiple(a, b)
	fmt.Println(c)
}