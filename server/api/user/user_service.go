package user

type Service interface {
	Get() string
	InsertOne() string
}

type service struct {
	userRepository Repository
}

func NewService(userRepository Repository) Service {
	return &service{
		userRepository: userRepository,
	}
}

func (s *service) Get() string {
	return s.userRepository.Get()
}

func (s *service) InsertOne() string {
	return s.userRepository.InsertOne()
}
