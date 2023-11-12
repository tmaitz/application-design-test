package service

var OrderService OrderServiceType

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
