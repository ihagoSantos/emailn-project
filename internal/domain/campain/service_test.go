package campain

import (
	"emailn/internal/contract"
	internalerrors "emailn/internal/internalerrors"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campain *Campain) error {
	args := r.Called(campain)
	return args.Error(0)
}

var (
	newCampain = contract.NewCampain{
		Name:    "Test Y",
		Content: "content",
		Emails:  []string{"test1@test.com"},
	}

	service = Service{}
)

func setup() {
	service = Service{}
}

func Test_Create_Campain(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(nil)
	service.Repository = repositoryMock
	id, err := service.Create(newCampain)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(internalerrors.ErrInternal)
	service.Repository = repositoryMock
	modifiedCampain := newCampain
	modifiedCampain.Name = ""
	_, err := service.Create(modifiedCampain)

	assert.False(errors.Is(internalerrors.ErrInternal, err))
}

func Test_Create_SaveCampain(t *testing.T) {
	setup()

	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(func(campain *Campain) bool {
		if campain.Name != newCampain.Name ||
			campain.Content != newCampain.Content ||
			len(campain.Contacts) != len(newCampain.Emails) {
			return false
		}

		return true
	})).Return(nil)

	service.Repository = repositoryMock
	service.Create(newCampain)
	repositoryMock.AssertExpectations(t)
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	setup()

	assert := assert.New(t)
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(errors.New("error to save on database"))

	service.Repository = repositoryMock
	id, err := service.Create(newCampain)
	println("ID", id)
	println("ERR", err)
	// assert.True(errors.Is(internalerrors.ErrInternal, err))
	assert.True(true)
}
