package adding

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockRepository struct {
	id  int64
	err error
}

func (m *mockRepository) Add(u User) (int64, error) {
	if m.err != nil {
		return -1, m.err
	}
	return m.id, nil
}

func TestAddUser(t *testing.T) {
	user := User{
		"Name",
		"email@example.com",
		"1231231234",
	}

	service := NewService(&mockRepository{
		1,
		nil,
	})

	actualUserID, _ := service.AddUser(user)

	assert.Equal(t, int64(1), actualUserID.ID)
}
