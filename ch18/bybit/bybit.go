package bybit

import "fmt"

type BybitSender struct{
}

func (bybit *BybitSender) Send(signal string) {
	fmt.Printf("Bybit sends %v signal\n", signal)
}