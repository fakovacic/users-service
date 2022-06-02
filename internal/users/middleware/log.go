package middleware

import (
	"context"

	"github.com/fakovacic/users-service/internal/users"
)

type loggingMiddleware struct {
	next    users.Service
	config  *users.Config
	service string
}

// NewLoggingMiddleware creates a new logging middleware.
func NewLoggingMiddleware(next users.Service, config *users.Config) users.Service {
	m := loggingMiddleware{
		next:    next,
		config:  config,
		service: "users",
	}

	return &m
}

func (m *loggingMiddleware) List(ctx context.Context, meta *users.Meta) (*users.Meta, []*users.User, error) {
	reqID := users.GetCtxStringVal(ctx, users.ContextKeyRequestID)

	m.config.Log.Info().
		Str("reqID", reqID).
		Str("service", m.service).
		Str("method", "List").
		Interface("meta", meta).
		Msg("service request")

	meta, list, err := m.next.List(ctx, meta)

	m.config.Log.Info().
		Str("reqID", reqID).
		Str("service", m.service).
		Str("method", "List").
		Interface("meta", meta).
		Interface("models", list).
		Err(err).
		Msg("service response")

	return meta, list, err
}

func (m *loggingMiddleware) Create(ctx context.Context, input *users.User) (*users.User, error) {
	reqID := users.GetCtxStringVal(ctx, users.ContextKeyRequestID)

	m.config.Log.Info().
		Str("reqID", reqID).
		Str("service", m.service).
		Str("method", "Create").
		Interface("input", input).
		Msg("service request")

	model, err := m.next.Create(ctx, input)

	m.config.Log.Info().
		Str("reqID", reqID).
		Str("service", m.service).
		Str("method", "Create").
		Interface("model", model).
		Err(err).
		Msg("service response")

	return model, err
}

func (m *loggingMiddleware) Update(ctx context.Context, id string, input *users.User, fields []string) (*users.User, error) {
	reqID := users.GetCtxStringVal(ctx, users.ContextKeyRequestID)

	m.config.Log.Info().
		Str("reqID", reqID).
		Str("service", m.service).
		Str("method", "Update").
		Str("id", id).
		Interface("input", input).
		Interface("fields", fields).
		Msg("service request")

	model, err := m.next.Update(ctx, id, input, fields)

	m.config.Log.Info().
		Str("reqID", reqID).
		Str("service", m.service).
		Str("method", "Update").
		Interface("model", model).
		Err(err).
		Msg("service response")

	return model, err
}

func (m *loggingMiddleware) Delete(ctx context.Context, id string) error {
	reqID := users.GetCtxStringVal(ctx, users.ContextKeyRequestID)

	m.config.Log.Info().
		Str("reqID", reqID).
		Str("service", m.service).
		Str("method", "Delete").
		Str("id", id).
		Msg("service request")

	err := m.next.Delete(ctx, id)

	m.config.Log.Info().
		Str("reqID", reqID).
		Str("service", m.service).
		Str("method", "Delete").
		Err(err).
		Msg("service response")

	return err
}
