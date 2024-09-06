package main

import (
	"fmt"
)

func main() {
	str := "Hello, 시우우우우 World!"
	
	for i := 0; i < len(str); i++ {
		fmt.Printf(" 타입:%T	값:%d	문자값:%c\n", str[i],str[i],str[i])
	}
}