package campain

import (
	"emailn/internal/contract"
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
		Content: "Body",
		Emails:  []string{"test1@test.com"},
	}

	service = Service{}
)

func Test_Create_Campain(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(nil)
	service.Repository = repositoryMock
	id, err := service.Create(newCampain)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_ValidateComainError(t *testing.T) {
	assert := assert.New(t)
	newCampain.Name = ""
	_, err := service.Create(newCampain)

	assert.NotNil(err)
	assert.Equal("name is required", err.Error())
}

func Test_Create_SaveCampain(t *testing.T) {

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
	assert := assert.New(t)
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(errors.New("error to save on database"))

	service.Repository = repositoryMock
	_, err := service.Create(newCampain)

	assert.Equal("error to save on database", err.Error())
}
