package main

type orderStatus string

const (
	Pending   orderStatus = "pending" //
	Confirmed             = "Confirmed"
	Delivered             = "Delivered"
	Canceled              = "Canceled"
)

func setOrderstatus(status orderStatus) {

}
func main() {
	setOrderstatus(Confirmed)
}
