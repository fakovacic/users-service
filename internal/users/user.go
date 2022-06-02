package users

import (
	"time"
)

type User struct {
	ID        string
	FirstName string
	LastName  string
	Nickname  string
	Password  string
	Email     string
	Country   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

var UserField = struct {
	ID        string
	FirstName string
	LastName  string
	Nickname  string
	Password  string
	Email     string
	Country   string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	FirstName: "first_name",
	LastName:  "last_name",
	Nickname:  "nickname",
	Password:  "password",
	Email:     "email",
	Country:   "country",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var UserUpdateFields = map[string]bool{
	UserField.ID:        false,
	UserField.FirstName: true,
	UserField.LastName:  true,
	UserField.Nickname:  true,
	UserField.Password:  true,
	UserField.Email:     true,
	UserField.Country:   true,
	UserField.CreatedAt: false,
	UserField.UpdatedAt: false,
}

var UserFilterFields = []string{
	UserField.Country,
	UserField.FirstName,
	UserField.LastName,
}
