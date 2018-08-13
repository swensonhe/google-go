package mock

import "github.com/swensonhe/google-go"

// Service is a mock interface for google.Service.
type Service struct {
	GetMeFn      func(code string) (*google.User, error)
	GetMeInvoked bool
}

// NewService returns a default implementation of google.Service.
func NewService() *Service {
	return &Service{
		GetMeFn: func(code string) (*google.User, error) {
			return &google.User{}, nil
		},
	}
}

func (s *Service) GetTokenInfo(code string) (*google.User, error) {
	s.GetMeInvoked = true
	return s.GetMeFn(code)
}
