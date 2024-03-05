package user

import "testing"

func TestRegister(t *testing.T) {
	u := User{Name: "Sooyeon", Email: "abc@gmail.com"}
	Register(u)

	retUser, exists := GetUser(u.Email)

	if !exists {
		t.Errorf("존재하지 않음 : %v", exists)
	}

	if retUser.Name != u.Name {
		t.Errorf("User name was incorrect, got: %s, want: %s.", retUser.Name, "Sooyeon")
	}

	if retUser.Email != u.Email {
		t.Errorf("User email was incorrect, got: %s, want: %s.", retUser.Email, "abc@gmail.com")
	}
}
