package main

import (
	"itsware/internal/controllers"
	"itsware/internal/repositories"
	"itsware/internal/services"
	"itsware/pkg/db"
	"net/http"
)

func main() {
	database := db.Connect()
	repo := &repositories.Device{DB: database}
	service := &services.Device{Repository: repo}
	controllers := &controllers.Device{Service: service}

	http.HandleFunc("/devices/create", controllers.Create)
	http.HandleFunc("/devices/get", controllers.Get)
	http.HandleFunc("/devices/update", controllers.Update)
	http.HandleFunc("/devices/delete", controllers.Delete)
	http.HandleFunc("/devices", controllers.GetAll)

	http.ListenAndServe(":8080", nil)
}
