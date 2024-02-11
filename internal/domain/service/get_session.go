package service

import (
	"context"

	"github.com/friendsofgo/errors"
	"github.com/ksaritek/paper-rock-scissors/internal/domain/model"
	"github.com/ksaritek/paper-rock-scissors/internal/observability"
)

func (s *sessionService) GetSession(ctx context.Context, sessionId string) (_ *model.Session, retErr error) {
	ctx, span := observability.Tracer().Start(ctx, "service_get_session")
	defer span.End()
	defer observability.RecordErr(span, retErr)

	session, err := s.sessionRepo.Get(ctx, sessionId)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to get session")
	}

	return session, nil
}
