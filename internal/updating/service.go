package updating

import "strconv"

//Service interface for updating users
type Service interface {
	UpdateUser(string, User) error
}

//Repository used to interact with the data layer.
//This is platform agnostic
type Repository interface {
	Update(int64, User) error
}

type service struct {
	usrRepo Repository
}

//NewService returns a service for updating a user
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) UpdateUser(id string, u User) error {
	//Validation should be performed on the input

	idLong, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	err = s.usrRepo.Update(idLong, u)
	if err != nil {
		return err
	}
	return nil
}
