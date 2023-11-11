package controller

import (
	"applicationDesignTest/service"
	"time"
)

type UrlParamValidationError struct {
	paramName string
	message   string
}

func (e *UrlParamValidationError) Error() string {
	return "Param '" + e.paramName + "' is not valid, because: " + e.message
}

func getValidValue(value string, key string) (string, error) {
	if value == "" {
		return "", &UrlParamValidationError{key, "value must not be empty"}
	}
	return value, nil
}

func getRoomTypeValidValue(value string, key string) (string, error) {
	if _, ok := service.AvailableRoomTypes[value]; !ok {
		return "", &UrlParamValidationError{key, "unknown room type"}
	}
	return value, nil
}

func getDateValidValue(value string, key string) (time.Time, error) {
	valueTime, err := time.Parse(time.DateOnly, value)
	if err != nil {
		return time.Time{}, &UrlParamValidationError{key, "can't parse as date"}
	}
	return valueTime, nil
}
