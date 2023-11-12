package service

func (OrderServiceType) GetOrders(email string) ([]Order, error) {
	orders := []Order{}
	for _, item := range OrderService.actualOrders {
		if item.userEmail == email {
			orders = append(orders, item)
		}
	}
	return orders, nil
}
