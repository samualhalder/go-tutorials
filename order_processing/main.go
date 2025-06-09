package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	id     int
	status string
	mu     sync.Mutex
}

var (
	updatedOrders int
	updateMu      sync.Mutex
)

func main() {
	orders := createOrders(20)
	var wg sync.WaitGroup
	wg.Add(5)
	// go processOrder(orders, &wg)

	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			for _, order := range orders {
				updateOrder(order)
			}

		}()
	}

	// go printOrderReport(orders, &wg)
	wg.Wait()
	fmt.Printf("total order updated %d", updatedOrders)
}

func createOrders(count int) []*Order {

	orders := make([]*Order, count)
	for i := 0; i < count; i++ {
		order := Order{id: i + 1, status: "Pending"}
		orders[i] = &order
	}
	return orders
}
func processOrder(orders []*Order, wg *sync.WaitGroup) {
	defer wg.Done()

	for _, order := range orders {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		fmt.Printf("processing order %d\n", order.id)
	}
}
func updateOrder(order *Order) {
	order.mu.Lock()
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	status := []string{
		"Dispatched",
		"Delivared",
		"Canceled",
	}[rand.Intn(3)]
	order.status = status
	order.mu.Unlock()

	fmt.Printf("Order id %d is updated with status :%s\n", order.id, status)
	updateMu.Lock()
	updatedOrders += 1
	updateMu.Unlock()
}

func printOrderReport(orders []*Order, wg *sync.WaitGroup) {
	defer wg.Done()

	for range 5 {
		time.Sleep(1 * time.Second)
		fmt.Println("----------------Order report starts------------")
		for _, order := range orders {
			fmt.Printf("Oder id : %d, order status : %s\n", order.id, order.status)
		}
		fmt.Println("----------------Order report ends------------")

	}
}
