package users

import (
	"time"
)

type User struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Nickname  string    `json:"nickname"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
