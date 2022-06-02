package users

import (
	"context"
	"time"

	"github.com/google/uuid"
)

//go:generate moq -out ./mocks/service.go -pkg mocks  . Service
type Service interface {
	List(context.Context, *Meta) (*Meta, []*User, error)
	Create(context.Context, *User) (*User, error)
	Update(context.Context, *User, []string) (*User, error)
	Delete(ctx context.Context, id string) error
}

func New(c *Config, store Store, timeFunc func() time.Time, uuidFunc func() uuid.UUID) Service {
	return &service{
		config:   c,
		store:    store,
		timeFunc: timeFunc,
		uuidFunc: uuidFunc,
	}
}

type service struct {
	config   *Config
	store    Store
	timeFunc func() time.Time
	uuidFunc func() uuid.UUID
}
