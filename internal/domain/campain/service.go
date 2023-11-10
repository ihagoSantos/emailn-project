package campain

import (
	"emailn/internal/contract"
	internalerrors "emailn/internal/internalerrors"
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
		return "", internalerrors.ErrInternal
	}

	return campain.ID, nil
}
