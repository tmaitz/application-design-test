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

func addRoute(serveMux *http.ServeMux, path string, methodHandlerMap map[string]RequestHandler) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if handle, ok := methodHandlerMap[r.Method]; !ok {
			http.Error(w, "Method Not Allowed", 405)
			return
		} else {
			status, responseBody, err := handle(r)
			if err != nil {
				logger.Error("error while handling %s %s method: %s", r.Method, path, err)
				http.Error(w, http.StatusText(status), status)
				return
			}
			w.WriteHeader(status)

			if responseBody != nil {
				byteBody, err := json.Marshal(responseBody)
				if err != nil {
					logger.Error("error while parsing responseBody for %s %s method: %s", r.Method, path, err)
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
					return
				}
				_, err = w.Write(byteBody)
				if err != nil {
					logger.Error("error while writing response for %s %s method: %s", r.Method, path, err)
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
					return
				}
			}
			logger.Info("Method %s %s was successfully done", r.Method, path)
		}
	}
	serveMux.HandleFunc(path, fn)
}
