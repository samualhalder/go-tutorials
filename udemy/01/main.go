package main

import (
	"errors"
	"fmt"
	"log"
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

	err := truck.LoadTruck()
	if err != nil {
		return fmt.Errorf("error loading cargo: %w", err)
	}
	err = truck.UnloadTruck()
	if err != nil {
		return fmt.Errorf("error loading cargo: %w", err)
	}
	return nil
}

func main() {
	nt := &NormalTruck{id: "nt-1"}
	et := &ElectricTruck{id: "et-1"}
	if err := processTruck(nt); err != nil {
		log.Fatal(err)
	}
	if err := processTruck(et); err != nil {
		log.Fatal(err)
	}
	log.Println(nt.cargo, et.battery)
}
