package main

import (
	"applicationDesignTest/controller"
	"applicationDesignTest/logger"
	"applicationDesignTest/service"
	"encoding/json"
	"errors"
	"net/http"
	"os"
)

// RequestHandler /** returns status code, response and error
type RequestHandler func(*http.Request) (int, any, error)

var routes = map[string]map[string]RequestHandler{
	"/orders": {
		http.MethodPost: controller.CreateOrder,
		http.MethodGet:  controller.GetOrders,
	},
}

func main() {
	service.Init()

	mux := http.NewServeMux()
	for path, methodHandlerMap := range routes {
		addRoute(mux, path, methodHandlerMap)
	}

	err := http.ListenAndServe(":8080", mux)
	if errors.Is(err, http.ErrServerClosed) {
		logger.Info("server closed")
	} else if err != nil {
		logger.Error("error listening for server: %s", err)
		os.Exit(1)
	}
}

func addRoute(serveMux *http.ServeMux, path string, httpMethodHandlerMap map[string]RequestHandler) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		handle, containsHttpMethod := httpMethodHandlerMap[r.Method]
		if !containsHttpMethod {
			logger.Error("Method %s is not allowed for path %s", r.Method, path)
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		status, responseBody, err := handle(r)
		if err != nil {
			logger.Error("Error while handling %s %s method: %s", r.Method, path, err)
			http.Error(w, http.StatusText(status), status)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)

		if responseBody != nil {
			err := json.NewEncoder(w).Encode(responseBody)
			if err != nil {
				logger.Error("Error while writing response for %s %s method: %s", r.Method, path, err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
		}
		logger.Info("Method %s %s was successfully done", r.Method, path)
	}
	serveMux.HandleFunc(path, fn)
}
