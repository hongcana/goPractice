package main

import "fmt"

func main() {
	var a int
	var b int

	n, err := fmt.Scan(&a,&b) // 성공적으로 입력된 값의 개수와 입력 실패시 에러를 반환
	if err != nil {
		fmt.Println(n, err) // Downteamis LG를 입력 > 0 expected integer
	}	else {
		fmt.Println(n,a,b)
	}
}