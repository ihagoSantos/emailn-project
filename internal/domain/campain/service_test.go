package campain

import (
	"emailn/internal/contract"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Create_Campain(t *testing.T) {
	assert := assert.New(t)

	service := Service{}
	newCampain := contract.NewCampain{
		Name:    "Test Y",
		Content: "Body",
		Emails:  []string{"teste1@test.com"},
	}

	id, err := service.Create(newCampain)

	assert.NotNil(id)
	assert.Nil(err)
}
