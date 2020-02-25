package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"test-go-microservice/services"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userParam := req.URL.Query().Get("user_id")
	userId, err := strconv.ParseInt(userParam, 10, 64)

	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte("Need to pass an integer user id"))
		return
	}

	user, err := services.GetUser(userId)
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(err.Error()))
		return
	}

	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
