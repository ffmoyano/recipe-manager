package response

import (
	"encoding/json"
	"github.com/ffmoyano/gofer/logger"
	"net/http"
)

func Send(writer http.ResponseWriter, statusCode int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		logger.Error(err.Error())
	}
	writer.WriteHeader(statusCode)
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("X-XSS-Protection", "1; mode=block")
	writer.Header().Set("X-Frame-Options", "deny")
	_, err = writer.Write(response)
	if err != nil {
		logger.Error(err.Error())
	}
}
