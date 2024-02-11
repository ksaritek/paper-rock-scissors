package service

import (
	"context"

	"github.com/friendsofgo/errors"
	"github.com/ksaritek/paper-rock-scissors/internal/domain/model"
	"github.com/ksaritek/paper-rock-scissors/internal/observability"
	"go.uber.org/zap"
)

func (s *sessionService) StartSession(ctx context.Context, playerId string) (_ *model.Session, retErr error) {
	ctx, span := observability.Tracer().Start(ctx, "service_start_session")
	defer span.End()
	defer observability.RecordErr(span, retErr)

	session := &model.Session{
		Id:       s.generateSessionId(),
		PlayerId: playerId,
		Wins:     0,
		Losses:   0,
		Draws:    0,
	}

	err := s.sessionRepo.Create(ctx, *session)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to create session")
	}

	s.logger.Debug("session created", zap.String("session_id", session.Id), zap.String("player_id", playerId))

	return session, nil
}
