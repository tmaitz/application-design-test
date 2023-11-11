package service

import "time"

const (
	economy  = "economy"
	standard = "standard"
	lux      = "lux"
)

var AvailableRoomTypes = map[string]bool{
	economy:  true,
	standard: true,
	lux:      true,
}

type OrderServiceType struct {
	ActualOrders   []Order
	AvailableRooms map[string]bool
}

var OrderService OrderServiceType

type OrderDtoRequest struct {
	Email string
	Room  string
	From  string
	To    string
}

type OrderDto struct {
	Email string
	Room  string
	From  time.Time
	To    time.Time
}

type Order struct {
	Room      string
	UserEmail string
	From      time.Time
	To        time.Time
}

var EmptyOrderStub = Order{}

func Init() {
	// here we can include logic of predefining of ActualOrders, AvailableRooms and etc. from DB or other services
	OrderService = OrderServiceType{
		[]Order{},
		map[string]bool{
			economy:  true,
			standard: true,
			lux:      true,
		},
	}
}
