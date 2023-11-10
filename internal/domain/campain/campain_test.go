package campain

import (
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

// Constants
var (
	fake    = faker.New()
	name    = "Campain X"
	content = "content"
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

func Test_NewCampain_MustValidateNameMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampain("", content, emails)

	assert.Equal("name is required with min 5", err.Error())

}

func Test_NewCampain_MustValidateNameMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampain(fake.Lorem().Text(25), content, emails)

	assert.Equal("name is required with max 24", err.Error())

}

func Test_NewCampain_MustValidateContentMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampain(name, "", emails)

	assert.Equal("content is required with min 5", err.Error())

}

func Test_NewCampain_MustValidateContentMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampain(name, fake.Lorem().Text(2000), emails)

	assert.Equal("content is required with max 1024", err.Error())

}

func Test_NewCampain_MustValidateContactsMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampain(name, content, []string{})

	assert.Equal("contacts is required with min 1", err.Error())

}

func Test_NewCampain_MustValidateContacts(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampain(name, content, []string{"email_invalid"})

	assert.Equal("email is invalid", err.Error())

}
