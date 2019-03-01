package listing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockRepository struct {
	user  User
	users []User
	err   error
}

func (m *mockRepository) Get(id int64) (User, error) {
	if m.err != nil {
		return User{}, m.err
	}
	return m.user, nil
}

func (m *mockRepository) GetAll() ([]User, error) {
	if m.err != nil {
		return []User{}, m.err
	}
	return m.users, nil
}

func TestGetUser(t *testing.T) {
	user := User{
		1,
		"Name",
		"email@example.com",
		"1231231234",
	}

	service := NewService(&mockRepository{
		user,
		[]User{},
		nil,
	})

	actualUser, err := service.GetUser("1")

	assert.NoError(t, err)

	assert.Equal(t, user.ID, actualUser.ID)
	assert.Equal(t, user.Name, actualUser.Name)
	assert.Equal(t, user.Email, actualUser.Email)
	assert.Equal(t, user.Phone, actualUser.Phone)
}

func TestGetUserList(t *testing.T) {
	user := User{
		1,
		"Name",
		"email@example.com",
		"1231231234",
	}

	service := NewService(&mockRepository{
		User{},
		[]User{user},
		nil,
	})

	actualUserList, err := service.GetUserList()

	assert.NoError(t, err)

	assert.Equal(t, 1, len(actualUserList.Users))

	actualUser := actualUserList.Users[0]

	assert.Equal(t, user.ID, actualUser.ID)
	assert.Equal(t, user.Name, actualUser.Name)
	assert.Equal(t, user.Email, actualUser.Email)
	assert.Equal(t, user.Phone, actualUser.Phone)

}
