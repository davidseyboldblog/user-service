package main

import (
	"net/http"

	"userservice/internal/adding"
	"userservice/internal/http/rest"
	"userservice/internal/listing"
	"userservice/internal/storage"
	"userservice/internal/updating"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// Only log the Info severity or above.
	logrus.SetLevel(logrus.InfoLevel)

	storage, err := storage.New("./user.db")
	if err != nil {
		logrus.Error(err)
		panic(err)
	}

	addService := adding.NewService(storage)
	updateService := updating.NewService(storage)
	listService := listing.NewService(storage)

	router := rest.Handler(addService, listService, updateService)

	err = http.ListenAndServe(":12000", router)
	if err != nil {
		logrus.Errorf("Server exited with error: %v", err)
	}
}
