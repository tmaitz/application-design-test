package controller

import (
	"applicationDesignTest/service"
	"net/http"
)

func GetOrders(r *http.Request) (int, any, error) {
	// get params from request and do validation
	query := r.URL.Query()
	email, err := getValidValue(query.Get("email"), "email")
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	// call service logic
	orders, err := service.OrderService.GetOrders(email)
	// handle service's error
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	// return result
	return http.StatusOK, orders, nil
}
