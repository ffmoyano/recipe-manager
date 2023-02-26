package handler

import (
	"encoding/json"
	"github.com/ffmoyano/gofer/logger"
	"net/http"
	"recipes/src/database"
	"recipes/src/entity"
	"recipes/src/response"
)

type UserHandler struct {
}

func (handler *UserHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	switch {
	case request.Method == http.MethodGet:
		handler.get(writer, request)
		return
	case request.Method == http.MethodPost:
		handler.insert(writer, request)
		return
	default:
		response.Send(writer, http.StatusMethodNotAllowed, map[string]string{"Error": "Method " + request.Method + " not allowed."})
		return
	}
}

func (handler *UserHandler) get(writer http.ResponseWriter, _ *http.Request) {
	users, err := database.GetUsers()
	if err != nil {
		response.Send(writer, http.StatusInternalServerError,
			map[string]string{"Error": "Ha ocurrido un error al tratar de recuperar los usuarios"})
	} else {
		response.Send(writer, http.StatusOK, users)
	}
}

func (handler *UserHandler) insert(writer http.ResponseWriter, request *http.Request) {

	var user entity.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		logger.Error(err.Error())
		response.Send(writer, http.StatusInternalServerError,
			map[string]string{"Error": "Ha ocurrido un error al tratar de insertar el usuario"})
	}

}
