package main

import (
	"time"
)

type Courier struct {
	Name	string
}

type Product struct {
	Name	string
	Price	int
	ID		int
}

type Parcel struct {
	Pdt		*Product
	ShippedTime	time.Time
	DeliveredTime time.Time
}

func (c *Courier) SendProduct(p *Product) *Parcel {
	par := &Parcel{}
	par.Pdt = p
	par.ShippedTime = time.Now()
	return par
}

func (p *Parcel) Delivered() *Product {
	p.DeliveredTime = time.Now()
	return p.Pdt
}