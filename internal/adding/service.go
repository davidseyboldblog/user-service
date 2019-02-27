package adding

//Service interface for registering users
type Service interface {
	AddUser(User) (UserID, error)
}

//Repository used to interact with the data layer.
//This is platform agnostic
type Repository interface {
	Add(User) (int64, error)
}

type service struct {
	usrRepo Repository
}

//NewService returns a service for registering a user
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) AddUser(u User) (UserID, error) {

	//Validation should be done on the input

	userID, err := s.usrRepo.Add(u)
	if err != nil {
		return UserID{}, err
	}
	return UserID{userID}, nil
}
