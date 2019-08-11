package user

type Service struct {
	user ServiceInterface
}

func (s *Service) Close() {}
