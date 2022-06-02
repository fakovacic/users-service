package users

import (
	"context"
)

type Store interface {
	List(context.Context, *Meta) (int64, []*User, error)
	Get(context.Context, string) (*User, error)
	Create(context.Context, *User) (*User, error)
	Update(context.Context, *User, []string) (*User, error)
	Delete(context.Context, string) error
}
