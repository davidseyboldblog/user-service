package listing

import "strconv"

//Service interface for retrieving users
type Service interface {
	GetUser(string) (User, error)
	GetUserList() (UserList, error)
}

//Repository used to interact with the data layer.
//This is database agnostic
type Repository interface {
	Get(int64) (User, error)
	GetAll() ([]User, error)
}

type service struct {
	usrRepo Repository
}

//NewService returns a service for retrieving a user
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetUser(id string) (User, error) {

	//Validation should be done on the input
	idLong, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return User{}, err
	}
	user, err := s.usrRepo.Get(idLong)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (s *service) GetUserList() (UserList, error) {

	users, err := s.usrRepo.GetAll()
	if err != nil {
		return UserList{}, err
	}
	return UserList{users}, nil
}
