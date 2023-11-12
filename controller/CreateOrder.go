package controller

import (
	"applicationDesignTest/service"
	"encoding/json"
	"errors"
	"net/http"
)

func CreateOrder(r *http.Request) (int, any, error) {
	// get params from request body and do validation
	var orderDtoRequest service.OrderDtoRequest
	err := json.NewDecoder(r.Body).Decode(&orderDtoRequest)
	if err != nil {
		return http.StatusBadRequest, nil, errors.New("can't parse request body")
	}

	email, err1 := getValidValue(orderDtoRequest.Email, "email")
	room, err2 := getRoomTypeValidValue(orderDtoRequest.Room, "room")
	from, err3 := getDateValidValue(orderDtoRequest.From, "from")
	to, err4 := getDateValidValue(orderDtoRequest.To, "to")
	if !(from.Before(to)) {
		return http.StatusBadRequest, nil, &UrlParamValidationError{"from", "should be before 'to' value"}
	}
	if err := errors.Join(err1, err2, err3, err4); err != nil {
		return http.StatusBadRequest, nil, err
	}

	// to Dto
	orderDto := service.OrderDto{email, room, from, to}

	// call service logic
	createdOrder, err := service.OrderService.CreateOrder(orderDto)
	// handle service's error
	if err != nil {
		// TODO: status code depends on service error type???
		return http.StatusBadRequest, nil, err
	}

	return http.StatusCreated, service.ToDto(createdOrder), nil
}
