package main

import (
	"fmt"
	"time"
)

// struct is like a class equivlent of oop
type Car struct {
	model             string
	price             float32
	dateOfManufacture time.Time
	isBought          bool
}

// go does not support constructor but we may use a hack to achive the same

func newCar(model string, price float32, isBought bool) *Car {
	car := Car{
		model:             model,
		price:             price,
		dateOfManufacture: time.Now(),
		isBought:          isBought,
	}
	return &car
}

// recivers equvalent to methods in oop
func (c Car) getPrice() float32 {
	return c.price
}

// always pass the ref while updating
func (c *Car) setPrice(price float32) {
	c.price = price
}

func main() {
	// car := Car{
	// 	model:             "Tata",
	// 	price:             300000.5,
	// 	dateOfManufacture: time.Now(),
	// 	isBought:          false,
	// }
	// car2 := newCar("KIA", 100000, true)

	// // fmt.Println(car.getPrice())
	// car2.setPrice(400000)
	// // fmt.Println(car.getPrice())
	// fmt.Println(car2.getPrice())

	//short hand for declaring a struct
	c := struct {
		name  string
		price int
	}{"Mahindra", 550000}

	fmt.Println(c)
	fmt.Println(c.name)
	fmt.Println(c.price)
}
