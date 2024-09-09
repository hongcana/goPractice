package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var stdin = bufio.NewReader(os.Stdin)

func CheckInput() (int, error) {
	var foo int
	_, err := fmt.Scanln(&foo) // int 타입 값을 입력 받음
	if err != nil {
		stdin.ReadString('\n') // 만약 에러 발생시 입력 스트림을 비움(bufio.NewReader(os.Stdin)으로 만든 객체의 .ReadString('\n'))
	}
	return foo, err
}

func MakeRandomNumber() (int) {
	source := rand.NewSource(time.Now().UnixNano())
	n := rand.New(source)

	return n.Intn(100)
}

func main() {
	ranNum := MakeRandomNumber()
	cnt := 1

	for {
		fmt.Println("숫자를 입력하시요. ")
		num, err := CheckInput()
		if err != nil {
			fmt.Println("숫자만 입력하라고 했을텐데..")
		} else {
			if num > ranNum {
				fmt.Println("입력 값이 너무 커.")
			} else if num < ranNum {
				fmt.Println("입력 값이 너무 작아.")
			} else {
				fmt.Println("축하해.. 성불해. ", cnt, "번만에 맞췄어.")
				break
			}
			cnt++
		}
	}
}