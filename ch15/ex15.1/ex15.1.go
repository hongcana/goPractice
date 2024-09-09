package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// deprecated: rand.Seed(time.Now().UnixNano())
	source := rand.NewSource(time.Now().UnixNano())
    n := rand.New(source)

	fmt.Println(n.Intn(100))
}