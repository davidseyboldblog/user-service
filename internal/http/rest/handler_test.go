package rest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"userservice/internal/adding"
	"userservice/internal/listing"
	"userservice/internal/updating"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

type addService struct {
	userID adding.UserID
	err    error
}

func (s *addService) AddUser(u adding.User) (adding.UserID, error) {
	if s.err != nil {
		return adding.UserID{}, s.err
	}
	return s.userID, nil
}

type listService struct {
	user     listing.User
	userList listing.UserList
	err      error
}

func (s *listService) GetUser(id string) (listing.User, error) {
	if s.err != nil {
		return listing.User{}, s.err
	}
	return s.user, nil
}

func (s *listService) GetUserList() (listing.UserList, error) {
	if s.err != nil {
		return listing.UserList{}, s.err
	}
	return s.userList, nil
}

type updateService struct {
	err error
}

func (s *updateService) UpdateUser(id string, u updating.User) error {
	if s.err != nil {
		return s.err
	}
	return nil
}
func TestAddUser(t *testing.T) {
	router := httprouter.New()
	userID := adding.UserID{
		1,
	}
	service := &addService{
		userID,
		nil,
	}
	router.POST("/user-service/user", addUser(service))

	newUser := adding.User{
		"Name",
		"email@example.com",
		"1231231234",
	}
	b, _ := json.Marshal(newUser)

	req, _ := http.NewRequest("POST", "/user-service/user", bytes.NewBuffer(b))
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusCreated, rr.Code)

	decoder := json.NewDecoder(rr.Body)

	var actual adding.UserID
	err := decoder.Decode(&actual)
	assert.NoError(t, err)

	assert.Equal(t, userID.ID, actual.ID)
}

func TestGetUser(t *testing.T) {
	router := httprouter.New()

	user := listing.User{
		1,
		"Name",
		"email@example.com",
		"1231231234",
	}

	service := &listService{
		user,
		listing.UserList{},
		nil,
	}

	router.GET("/user-service/user/:id", getUser(service))
	req, _ := http.NewRequest("GET", "/user-service/user/1", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected HttpStatus of: %v but found: %v", http.StatusOK, status)
	}

	decoder := json.NewDecoder(rr.Body)

	var actual listing.User
	err := decoder.Decode(&actual)
	if err != nil {
		t.Fatalf("Unable to decode the response")
	}

	assert.Equal(t, user.ID, actual.ID)
	assert.Equal(t, user.Name, actual.Name)
	assert.Equal(t, user.Email, actual.Email)
	assert.Equal(t, user.Phone, actual.Phone)

}

func TestGetUserList(t *testing.T) {
	router := httprouter.New()
	user := listing.User{
		1,
		"Name",
		"email@example.com",
		"1231231234",
	}

	service := &listService{
		listing.User{},
		listing.UserList{[]listing.User{user}},
		nil,
	}

	router.GET("/user-service/user", getUserList(service))
	req, _ := http.NewRequest("GET", "/user-service/user", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected HttpStatus of: %v but found: %v", http.StatusOK, status)
	}

	decoder := json.NewDecoder(rr.Body)

	var actualList listing.UserList
	err := decoder.Decode(&actualList)
	if err != nil {
		t.Fatalf("Unable to decode the response")
	}

	assert.Equal(t, 1, len(actualList.Users))

	actualUser := actualList.Users[0]

	assert.Equal(t, user.ID, actualUser.ID)
	assert.Equal(t, user.Name, actualUser.Name)
	assert.Equal(t, user.Email, actualUser.Email)
	assert.Equal(t, user.Phone, actualUser.Phone)

}

func TestUpdateUser(t *testing.T) {
	router := httprouter.New()
	user := updating.User{
		"Name",
		"email@example.com",
		"1231231234",
	}

	service := &updateService{
		nil,
	}

	router.PUT("/user-service/user/:id", updateUser(service))
	b, _ := json.Marshal(user)

	req, _ := http.NewRequest("PUT", "/user-service/user/1", bytes.NewBuffer(b))
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusNoContent, rr.Code)

}
