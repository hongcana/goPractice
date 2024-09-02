package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("입력하시요.")
		var number int
		_, err := fmt.Scanln(&number)
		if err != nil {
			fmt.Println("정수를 입력하신게 맞습니까?")

			// 키보드 버퍼 비우기
			stdin.ReadString('\n')
			continue
		}
		fmt.Printf("입력하신 숫자는 %d입니다.\n", number)
		if number % 2 == 0 {
			break
		}
	}

	fmt.Println("bye")
}