package main

import "fmt"

type ParkingLot struct {
	LotSize	int
}

func ParkCar(lot *ParkingLot, carSize int) {
	lot.LotSize -= carSize
}

func (lot *ParkingLot) ParkingCar (carSize int) { 
	lot.LotSize -= carSize
}

func (lot ParkingLot) GetLotSize() int {
	return lot.LotSize
}

func main() {
	lot := &ParkingLot{ 100 }
	ParkCar(lot, 10)

	lot.ParkingCar(10)
	fmt.Println(lot.GetLotSize())
}