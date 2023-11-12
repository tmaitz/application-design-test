package service

import "time"

const (
	economy  = "economy"
	standard = "standard"
	lux      = "lux"
)

type OrderServiceType struct {
	ActualOrders   []Order
	AvailableRooms map[string]bool
}

// OrderDtoRequest comes from http request
type OrderDtoRequest struct {
	Email string
	Room  string
	From  string
	To    string
}

// OrderDto contains prepared data which comes from client, and should be visible for client
// TODO: maybe, it's good to use Input and Output types
type OrderDto struct {
	Email string
	Room  string
	From  time.Time
	To    time.Time
}

// Order internal representation
type Order struct {
	UserEmail string
	Room      string
	From      time.Time
	To        time.Time
}

// AvailableRoomTypes contains all available room types -> needs for room param validation
var AvailableRoomTypes = map[string]bool{
	economy:  true,
	standard: true,
	lux:      true,
}
