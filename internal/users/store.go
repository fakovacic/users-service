package users

import (
	"context"
)

//go:generate moq -out ./mocks/store.go -pkg mocks  . Store
type Store interface {
	List(context.Context, *Meta) (int64, []*User, error)
	Get(context.Context, string) (*User, error)
	Create(context.Context, *User) error
	Update(context.Context, string, *User) error
	Delete(context.Context, string) error
}
