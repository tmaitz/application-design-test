package service

import (
	"errors"
)

func (OrderServiceType) CreateOrder(orderDto OrderDto) (Order, error) {
	if !isRoomAvailable(orderDto) {
		return EmptyOrderStub, errors.New("there is no available room")
	}

	order := Order{
		orderDto.Room,
		orderDto.Email,
		orderDto.From,
		orderDto.To,
	}

	OrderService.ActualOrders = append(OrderService.ActualOrders, order)

	return order, nil
}

func isRoomAvailable(orderDto OrderDto) bool {
	return OrderService.AvailableRooms[orderDto.Room]
}
