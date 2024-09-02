package main

import "fmt"

type Direction int
const (
	None Direction = iota
	North
	East
	South
	West
)

func DirectionToString(d Direction) string {
	switch d {
	case North:
		return "North"
	case East:
		return "East"
	case South:
		return "South"
	case West:
		return "West"
	default:
		return "None"
	}
}

func GetDirection(angle float64) Direction {
	switch {
	case angle >= 315, (0 <= angle && angle < 45):
		return North
	case angle >= 45 && angle < 135:
		return East
	case angle >= 135 && angle < 225:
		return South
	case angle >= 225 && angle < 315:
		return West
	default:
		return None
	}
}

func GetAngle() float64 {
	var angle float64

	_, err := fmt.Scan(&angle)
	if err != nil {
		fmt.Println(err)
		return -1
	}

	return angle
}

func main() {
	fmt.Println(DirectionToString(GetDirection(GetAngle())))
}
