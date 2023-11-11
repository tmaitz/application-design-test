package service

func (OrderServiceType) GetOrders(email string) ([]Order, error) {
	var orders []Order
	for _, item := range OrderService.ActualOrders {
		if item.UserEmail == email {
			orders = append(orders, item)
		}
	}
	return orders, nil
}
