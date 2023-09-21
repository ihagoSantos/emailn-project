package campain

import "testing"

func TestNewCampain(t *testing.T) {
	name := "Campain X"
	content := "Body"
	emails := []string{"email1@e.com", "email2@e.com"}

	campain := NewCampain(name, content, emails)

	if campain.ID != "1" {
		t.Errorf("expected 1")
	} else if campain.Name != name {
		t.Errorf("expected correct name")
	} else if campain.Content != content {
		t.Errorf("expected correct content")
	} else if len(campain.Contacts) != len(emails) {
		t.Errorf("expected correct contacts")
	}
}
