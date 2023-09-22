package campain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Constants
var (
	name    = "Campain X"
	content = "Body"
	emails  = []string{"email1@e.com", "email2@e.com"}
)

func Test_NewCampain_CreateNewCampain(t *testing.T) {
	assert := assert.New(t)

	campain, _ := NewCampain(name, content, emails)

	assert.Equal(campain.Name, name)
	assert.Equal(campain.Content, content)
	assert.Equal(len(campain.Contacts), len(emails))

}

func Test_NewCampain_IDIsNotNil(t *testing.T) {
	assert := assert.New(t)

	campain, _ := NewCampain(name, content, emails)

	assert.NotNil(campain.ID)
}

func Test_NewCampain_CreatedOnMustBeNow(t *testing.T) {
	assert := assert.New(t)

	now := time.Now().Add(-time.Minute)

	campain, _ := NewCampain(name, content, emails)

	assert.NotNil(campain.CreatedOn)
	assert.Greater(campain.CreatedOn, now)

}

func Test_NewCampain_MustValidateName(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampain("", content, emails)

	assert.Equal("name is required", err.Error())

}

func Test_NewCampain_MustValidateContent(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampain(name, "", emails)

	assert.Equal("content is required", err.Error())

}

func Test_NewCampain_MustValidateContacts(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampain(name, content, []string{})

	assert.Equal("contacts is required", err.Error())

}
