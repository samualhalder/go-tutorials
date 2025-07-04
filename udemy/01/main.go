package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var (
	ErrorNotImplemented = errors.New("not implemented")
	ErrorWrongTruck     = errors.New("wrong truck")
)

type Truck interface {
	LoadTruck() error
	UnloadTruck() error
	GetId() string
}

type NormalTruck struct {
	id    string
	cargo int
}

func (t *NormalTruck) LoadTruck() error {
	t.cargo++
	return nil
}
func (t *NormalTruck) UnloadTruck() error {
	t.cargo = 0
	return nil
}
func (t *NormalTruck) GetId() string {
	return t.id
}

type ElectricTruck struct {
	id      string
	cargo   int
	battery float32
}

func (t *ElectricTruck) LoadTruck() error {
	t.cargo++
	t.battery--
	return nil
}
func (t *ElectricTruck) UnloadTruck() error {
	t.cargo = 0
	t.battery--
	return nil
}
func (t *ElectricTruck) GetId() string {
	return t.id
}

func processTruck(truck Truck) error {

	fmt.Printf("Processing truck %s\n", truck.GetId())
	time.Sleep(time.Second)

	err := truck.LoadTruck()
	if err != nil {
		return fmt.Errorf("error loading cargo: %w", err)
	}
	err = truck.UnloadTruck()
	if err != nil {
		return fmt.Errorf("error loading cargo: %w", err)
	}
	fmt.Printf("Processed truck %s\n", truck.GetId())
	return nil
}

func processFleet(trucks []Truck) error {
	var wg sync.WaitGroup
	for _, truck := range trucks {
		wg.Add(1)
		go func() {
			processTruck(truck)
			wg.Done()
		}()
	}
	wg.Wait()
	return nil
}

func main() {
	fleet := []Truck{
		&NormalTruck{id: "nt-1", cargo: 10},
		&ElectricTruck{id: "et-1", cargo: 11, battery: 100},
		&NormalTruck{id: "nt-2", cargo: 10},
		&ElectricTruck{id: "et-2", cargo: 20, battery: 100},
	}
	processFleet(fleet)

}
