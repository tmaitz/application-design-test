package controller

type UrlParamValidationError struct {
	paramName string
	message   string
}

func (e *UrlParamValidationError) Error() string {
	return "Param '" + e.paramName + "' is not valid, because: " + e.message
}
