package model

import (
	"testing"
)

func TestUser_Create(t *testing.T) {
	u := &User{Name: "2", Email: "2", Password: "1"}
	if err := u.Create(); err != nil {
		t.Fatal(err)
	}
}

func TestUser_ReadByID(t *testing.T) {
	u := &User{ID: 1}
	if err := u.ReadByID(); err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}

func TestUser_ReadByName(t *testing.T) {
	u := &User{Name: "1"}
	if err := u.ReadByName(); err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}

func TestUser_ReadByEmail(t *testing.T) {
	u := &User{Email: "1"}
	if err := u.ReadByEmail(); err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}

func TestUser_ReadAll(t *testing.T) {
	u := &User{}
	users, err := u.ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(users)
}

func TestUser_Update(t *testing.T) {
	u := &User{ID: 1, Name: "2", Email: "2", Password: "2"}
	if err := u.Update(); err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}

func TestUser_Delete(t *testing.T) {
	u := &User{ID: 1}
	if err := u.Delete(); err != nil {
		t.Fatal(err)
	}
}

func TestUser_DeleteAll(t *testing.T) {
	u := &User{}
	if err := u.DeleteAll(); err != nil {
		t.Fatal(err)
	}
}
