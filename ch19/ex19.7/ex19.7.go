package main

import (
	"fmt"
	"os"
)

type Writer func(string)

func writeheroesOfTheStorm(wri Writer) {
	wri("Play heroesOfTheStorm")
}

func main() {
	f, err := os.Create("hots.txt")
	if err != nil {
		fmt.Println("Failed to create file.")
		return
	}

	defer f.Close()

	writeheroesOfTheStorm(func(msg string) {
		fmt.Fprintln(f, msg)
	})
}