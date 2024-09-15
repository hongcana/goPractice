package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}

	defer fmt.Println("stack으로 반대로 표시됩니당.")
	defer f.Close()
	defer fmt.Println("file은 반드시 사용후 닫아야 합니다.")

	fmt.Println("파일에 Hell, World!를 씁니다.")
	fmt.Fprintln(f, "Hell, World!")
}