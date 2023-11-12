package service

import (
	"errors"
)

func (OrderServiceType) CreateOrder(orderDto OrderDto) (Order, error) {
	if !isRoomAvailable(orderDto) {
		return EmptyOrderStub, errors.New("there is no available room")
	}

	order := ToEntity(orderDto)

	OrderService.actualOrders = append(OrderService.actualOrders, order)

	return order, nil
}

func isRoomAvailable(orderDto OrderDto) bool {
	return OrderService.availableRooms[orderDto.Room]
}
