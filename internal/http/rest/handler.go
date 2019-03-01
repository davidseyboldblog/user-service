package rest

import (
	"encoding/json"
	"net/http"
	"os"

	"userservice/internal/adding"
	"userservice/internal/listing"
	"userservice/internal/updating"

	"github.com/gorilla/handlers"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

const (
	//ContentType header for type of content in response
	ContentType = "Content-Type"
	//ApplicationJSON value for content type in response
	ApplicationJSON = "application/json"
	//InternalServerErrorMessage message for an unexpected error
	InternalServerErrorMessage = "Internal Server Error"
	//BadRequestMessage message for a bad request sent from client
	BadRequestMessage = "Bad Request"
)

type Server struct {
	addService    adding.Service
	listService   listing.Service
	updateService updating.Service
	router        *httprouter.Router
}

// Handler initializes routes
func Handler(addService adding.Service, listService listing.Service, updateService updating.Service) http.Handler {
	router := httprouter.New()

	logrus.Info("Initalizing routes")
	router.POST("/user-service/user", addUser(addService))
	router.PUT("/user-service/user/:id", updateUser(updateService))
	router.GET("/user-service/user/:id", getUser(listService))
	router.GET("/user-service/user", getUserList(listService))

	logrus.Info("Adding Middleware")
	handler := handlers.CombinedLoggingHandler(os.Stdout, router)
	return handler
}

func addUser(s adding.Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		decoder := json.NewDecoder(r.Body)

		var user adding.User
		err := decoder.Decode(&user)
		if err != nil {
			logrus.Error(err)
			returnError(w, http.StatusBadRequest, BadRequestMessage)
			return
		}

		userID, err := s.AddUser(user)
		if err != nil {
			logrus.Error(err)
			returnError(w, http.StatusInternalServerError, InternalServerErrorMessage)
			return
		}

		w.Header().Set(ContentType, ApplicationJSON)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(userID)
	}
}

func updateUser(s updating.Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		decoder := json.NewDecoder(r.Body)

		id := ps.ByName("id")

		var user updating.User

		err := decoder.Decode(&user)
		if err != nil {
			logrus.Error(err)
			returnError(w, http.StatusBadRequest, BadRequestMessage)
			return
		}

		err = s.UpdateUser(id, user)

		w.Header().Set(ContentType, ApplicationJSON)
		w.WriteHeader(http.StatusNoContent)
	}
}

func getUser(s listing.Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		id := ps.ByName("id")

		user, err := s.GetUser(id)
		if err != nil {
			logrus.Error(err)
			returnError(w, http.StatusInternalServerError, InternalServerErrorMessage)
			return
		}

		w.Header().Set(ContentType, ApplicationJSON)
		json.NewEncoder(w).Encode(user)
	}
}

func getUserList(s listing.Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		userList, err := s.GetUserList()
		if err != nil {
			logrus.Error(err)
			returnError(w, http.StatusInternalServerError, InternalServerErrorMessage)
			return
		}

		w.Header().Set(ContentType, ApplicationJSON)
		json.NewEncoder(w).Encode(userList)
	}
}

func returnError(w http.ResponseWriter, httpCode int, message string) {
	w.Header().Set(ContentType, ApplicationJSON)
	w.WriteHeader(httpCode)

	json.NewEncoder(w).Encode(message)
}
