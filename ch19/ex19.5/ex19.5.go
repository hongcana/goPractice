package main

import "fmt"

func main() {

	i := 0

	f := func() {
		i += 10
	}
	
	f()
	i++
	fmt.Println(i)
}