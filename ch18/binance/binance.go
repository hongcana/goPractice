package binance

import "fmt"

type BinanceSender struct{
}

func (binance *BinanceSender) Send(signal string) {
	fmt.Printf("Binance sends %v signal\n", signal)
}