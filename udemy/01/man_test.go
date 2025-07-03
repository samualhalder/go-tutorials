package main

import (
	"log"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("Load and Unload cargos", func(t *testing.T) {
		nt := &NormalTruck{id: "nt-1"}
		et := &ElectricTruck{id: "et-1"}
		if err := processTruck(nt); err != nil {
			log.Fatal(err)
		}
		if err := processTruck(et); err != nil {
			log.Fatal(err)
		}
		if nt.cargo != 0 {
			t.Fatal("truck should have 0 cargos")
		}
		if et.battery != -2 {
			t.Fatalf("e-Truck should have -2 battery")
		}
	})
}
