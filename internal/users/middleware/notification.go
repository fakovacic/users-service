package middleware

import (
	"context"

	"github.com/fakovacic/users-service/internal/users"
)

type notificationMiddleware struct {
	next     users.Service
	notifier users.Notifier
}

func NewNotificationMiddleware(next users.Service, notifier users.Notifier) users.Service {
	m := notificationMiddleware{
		next:     next,
		notifier: notifier,
	}

	return &m
}

func (m *notificationMiddleware) List(ctx context.Context, meta *users.Meta) (*users.Meta, []*users.User, error) {
	meta, list, err := m.next.List(ctx, meta)

	return meta, list, err
}

func (m *notificationMiddleware) Create(ctx context.Context, input *users.User) (*users.User, error) {
	model, err := m.next.Create(ctx, input)

	return model, err
}

func (m *notificationMiddleware) Update(ctx context.Context, id string, input *users.User, fields []string) (*users.User, error) {
	model, err := m.next.Update(ctx, id, input, fields)

	if err == nil {
		m.notifier.Send(ctx)
	}

	return model, err
}

func (m *notificationMiddleware) Delete(ctx context.Context, id string) error {
	err := m.next.Delete(ctx, id)

	return err
}
