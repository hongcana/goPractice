package main

import (
	"bufio" // io를 담당하는 패키지
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)

	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)

	if err != nil {
		fmt.Println(n, err)
		stdin.ReadString('\n') // 표준 입력 스트림 지우기? -> '\n' 문자가 나올때까지 읽는다.
	} else {
		fmt.Println(n,a,b)
	}
	n, err = fmt.Scanln(&a,&b)
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n,a,b)
	}
}