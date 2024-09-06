package main

import "fmt"

func main() {
	str := "NEVER SELL YOUR $BTC"
	for _, v := range str {
		fmt.Printf(" 타입:%T	값:%d	문자값:%c\n", v,v,v)
	}	
}