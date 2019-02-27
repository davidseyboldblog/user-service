package listing

import "testing"

type mockRepository struct {
	user User
	err  error
}

func (m *mockRepository) Get(id int64) (User, error) {
	if m.err != nil {
		return User{}, m.err
	}
	return m.user, nil
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
		nil,
	})

	actualUser, _ := service.GetUser("1")

	if actualUser.ID != user.ID {
		t.Fatalf("Expected ID to equal: %v but actual: %v", user.ID, actualUser.ID)
	} else if actualUser.Name != user.Name {
		t.Fatalf("Expected Name to equal: %v but actual: %v", user.Name, actualUser.Name)
	} else if actualUser.Email != user.Email {
		t.Fatalf("Expected Email to equal: %v but actual: %v", user.Email, actualUser.Email)
	} else if actualUser.Phone != user.Phone {
		t.Fatalf("Expected Phone to equal: %v but actual: %v", user.Phone, actualUser.Phone)
	}
}
