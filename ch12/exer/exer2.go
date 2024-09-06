package main

import "fmt"

type Actor struct {
	Name 	string
	HP		int
	Speed	float64
}

func NewActor(name string, hp int, speed float64) *Actor {
	actor := new(Actor)
	actor.Name = name
	actor.HP = hp
	actor.Speed = speed
	return actor
}

func main() {
	var actor = NewActor("Goldrabbit", 99, 100)
	fmt.Println(actor.Speed)
	fmt.Println(actor.Name)
}