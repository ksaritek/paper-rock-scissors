package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/ksaritek/paper-rock-scissors/internal/domain/model"
	"github.com/ksaritek/paper-rock-scissors/internal/observability"
)

func (s *sessionService) StartSession(ctx context.Context, playerId string) (_ *model.Session, retErr error) {
	ctx, span := observability.Tracer().Start(ctx, "service_start_session")
	defer span.End()
	defer observability.RecordErr(span, retErr)

	session := &model.Session{
		Id:       generateSessionID(),
		PlayerId: playerId,
		Wins:     0,
		Losses:   0,
		Draws:    0,
	}

	err := s.sessionRepo.Create(ctx, *session)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to create session")
	}

	return session, nil
}

func generateSessionID() string {
	timestamp := time.Now().UnixNano()
	hash := sha256.Sum256([]byte(strconv.FormatInt(timestamp, 10)))
	return hex.EncodeToString(hash[:])
}
