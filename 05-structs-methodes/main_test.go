package models

import (
	"testing"
)

func TestUser(t *testing.T) {
	u := NewUser(1, "bob", "bob@example.com")

	t.Run("Check Fields", func(t *testing.T) {
		if u.ID != 1 || u.Username != "bob" || u.Email != "bob@example.com" {
			t.Errorf("NewUser failed, fields not set correctly: %+v", u)
		}
	})

	t.Run("GetInfo", func(t *testing.T) {
		got := u.GetInfo()
		want := "ID: 1, Username: bob"
		if got != want {
			t.Errorf("GetInfo() = %q; want %q", got, want)
		}
	})

	t.Run("UpdateEmail", func(t *testing.T) {
		u.UpdateEmail("new@example.com")
		if u.Email != "new@example.com" {
			t.Errorf("UpdateEmail failed, got %q; want \"new@example.com\"", u.Email)
		}
	})
}
