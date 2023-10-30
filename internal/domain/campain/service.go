package campain

import (
	"emailn/internal/contract"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampain contract.NewCampain) (string, error) {
	campain, err := NewCampain(newCampain.Name, newCampain.Content, newCampain.Emails)
	if err != nil {
		return "", err
	}
	err = s.Repository.Save(campain)

	if err != nil {
		return "", err
	}

	return campain.ID, nil
}
