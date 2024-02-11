package service

import (
	"context"

	"github.com/ksaritek/paper-rock-scissors/internal/domain/model"
	"github.com/ksaritek/paper-rock-scissors/internal/domain/repository"
	"github.com/ksaritek/paper-rock-scissors/internal/observability"
)

type Session interface {
	StartSession(ctx context.Context, playerId string) (*model.Session, error)
	GetSession(ctx context.Context, sessionId string) (*model.Session, error)
	DeleteSession(ctx context.Context, sessionId string) (*model.Session, error)
	Move(ctx context.Context, move model.Move, fn ComputerMove) (*model.Session, model.Choice, error)
}

type sessionService struct {
	sessionRepo repository.Session
}

func NewSessionService(ctx context.Context, sessionRepo repository.Session) Session {
	_, span := observability.Tracer().Start(ctx, "session_service_new")
	defer span.End()

	return &sessionService{
		sessionRepo: sessionRepo,
	}
}
