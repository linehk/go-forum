package model

import (
	"testing"
)

func TestUser(t *testing.T) {
	want := &User{Username: "test_username", Email: "test_email@gmail.com", Admin: "test_admin"}
	if err := want.Create(); err != nil {
		t.Error(err)
	}

	gotByUsername := &User{Username: want.Username}
	if err := gotByUsername.ReadByUsername(); err != nil {
		t.Error(err)
	}
	if !compareUser(gotByUsername, want) {
		t.Errorf("got %v, want %v", gotByUsername, want)
	}

	gotByID := gotByUsername
	if err := gotByID.ReadByID(); err != nil {
		t.Error(err)
	}
	if !compareUser(gotByID, want) {
		t.Errorf("got %v, want %v", gotByID, want)
	}

	gotByEmail := &User{Email: want.Email}
	if err := gotByEmail.ReadByEmail(); err != nil {
		t.Error(err)
	}
	if !compareUser(gotByEmail, want) {
		t.Errorf("got %v, want %v", gotByEmail, want)
	}

	gotByAdmin := &User{Admin: want.Admin}
	gotByAdmins, err := gotByAdmin.ReadByAdmin()
	if err != nil {
		t.Error(err)
	}
	if !compareUser(&gotByAdmins[0], want) {
		t.Errorf("got %v, want %v", &gotByAdmins[0], want)
	}

	gotByUpdate := &User{
		ID:       gotByUsername.ID,
		Username: "test_update",
		Email:    "test_update@gmail.com",
		Admin:    "test_update",
	}
	if err := gotByUpdate.Update(); err != nil {
		t.Error(err)
	}
	if compareUser(gotByUpdate, want) {
		t.Errorf("got %v, want %v", gotByUpdate, want)
	}

	if err := gotByUsername.Delete(); err != nil {
		t.Error(err)
	}
}

func compareUser(u1, u2 *User) bool {
	if u1.Username != u2.Username && u1.Email != u2.Email && u1.Admin != u2.Admin {
		return false
	}
	return true
}
