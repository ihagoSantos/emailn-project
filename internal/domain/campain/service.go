package campain

import (
	"emailn/internal/contract"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampain contract.NewCampain) (string, error) {
	campain, _ := NewCampain(newCampain.Name, newCampain.Content, newCampain.Emails)
	return campain.ID, nil
}
