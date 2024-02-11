package repository

import (
	"context"

	"github.com/ksaritek/paper-rock-scissors/internal/domain/model"
)

type Session interface {
	Get(ctx context.Context, sessionId string) (*model.Session, error)
	Create(ctx context.Context, session model.Session) error
	UpdateRoundResult(ctx context.Context, sessionId string, player model.RoundResult) (*model.Session, error)
	Delete(ctx context.Context, sessionId string) (*model.Session, error)
}
