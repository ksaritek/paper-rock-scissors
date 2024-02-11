package service

import (
	"context"

	"github.com/friendsofgo/errors"
	"github.com/ksaritek/paper-rock-scissors/internal/domain/model"
	"github.com/ksaritek/paper-rock-scissors/internal/observability"
)

func (s *sessionService) DeleteSession(ctx context.Context, sessionId string) (_ *model.Session, retErr error) {
	ctx, span := observability.Tracer().Start(ctx, "service_delete_session")
	defer span.End()
	defer observability.RecordErr(span, retErr)

	session, err := s.sessionRepo.Delete(ctx, sessionId)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to delete session")
	}

	return session, nil
}
