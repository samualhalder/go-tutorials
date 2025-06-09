package main

import "fmt"

// interface
type paymenter interface {
	pay(amount int) bool
	getBalane() int
}

// gateway is a class that is implementing paymenter interface
type payment struct {
	gateway paymenter
}

func (p payment) makePaymant(amount int) {
	p.gateway.pay(amount)
}

type razorpay struct {
}

// if the class has all the methods that the interface have then it autometically implementing the interface
func (r razorpay) pay(amount int) bool {
	fmt.Println("making a payment of ", amount, "using razorpay")
	return true
}

func (r razorpay) getBalane() int {

	return 1000
}

func main() {
	newRazorpay := razorpay{}
	newPayment := payment{
		gateway: newRazorpay,
	}
	newPayment.makePaymant(100)
}
