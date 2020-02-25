package app

import (
	"log"
	"net/http"
	"test-go-microservice/controllers"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
