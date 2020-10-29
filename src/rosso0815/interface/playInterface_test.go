package playinterface

import (
	"log"
	"testing"
)

type Car interface {
	GetBrand() string
	GetPrice() int
}

type VW struct {
	name  string
	price int
}

func newVWCar(name string, price int) *VW {
	car := VW{name: name, price: price}
	return &car
}

// implement the interface fuctions
func (vw *VW) GetBrand() string {
	return vw.name
}

func (vw *VW) GetPrice() int {
	return vw.price
}

func Test_Cars(t *testing.T) {
	log.Println("@@@ Test_Car")

	vwPolo := newVWCar("polo", 1000)
	vwGolf := newVWCar("golf", 2000)

	var car Car

	car = vwPolo

	log.Println(car.GetBrand(), car.GetPrice())

	car = vwGolf

	log.Println(car.GetBrand(), car.GetPrice())
}

func Test_Cars_Array(t *testing.T) {
	log.Println("@@@ Test_Car")

	var cars []VW

	cars = append(cars, *newVWCar("polo", 1000))
	cars = append(cars, *newVWCar("golf", 2000))

	for i := range cars {
		var iCar Car
		iCar = &cars[i]
		log.Println(iCar.GetBrand(), iCar.GetPrice())
	}

}
