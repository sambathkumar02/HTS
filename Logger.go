package main

import (
	"net/http"
	"time"
)

type Logger struct {
	TempDict    map[time.Time]interface{}
	LogFilePath string
	//count       int
}

func (logger Logger) Log() {

}

func (logger Logger) GetDataFromRequest(request *http.Request) {
	data := make(map[string]interface{})
	data["url"] = request.URL
	data["headers"] = request.Header
	data["body"] = request.Body
	data["method"] = request.Method
	data["client_addr"] = request.RemoteAddr

	logger.TempDict[time.Now()] = data

}
