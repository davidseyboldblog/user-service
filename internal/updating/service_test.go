package updating

import "testing"

type mockRepository struct {
	err error
}

func (m *mockRepository) Update(id int64, u User) error {
	if m.err != nil {
		return m.err
	}
	return nil
}

func TestUpdateUser(t *testing.T) {
	user := User{
		"Name",
		"Email",
		"Phone",
	}

	service := NewService(&mockRepository{
		nil,
	})

	err := service.UpdateUser("1", user)
	if err != nil {
		t.Fatalf("Expected success but instead got error: %v", err)
	}
}
