package listing

import "testing"

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

	actualUserList, _ := service.GetUserList()

	if len(actualUserList.Users) != 1 {
		t.Fatalf("Expected number of users to be: %v but found: %v", 1, len(actualUserList.Users))
	}

	actualUser := actualUserList.Users[0]

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
