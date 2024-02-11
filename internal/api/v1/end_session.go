package v1

import (
	"context"

	"github.com/ksaritek/paper-rock-scissors/internal/observability"
	pb "github.com/ksaritek/paper-rock-scissors/proto/gen/go/game/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *controller) EndSession(ctx context.Context, in *pb.EndSessionRequest) (_ *pb.EndSessionResponse, retErr error) {
	ctx, span := observability.Tracer().Start(ctx, "controller_end_session")
	defer span.End()
	defer observability.RecordErr(span, retErr)

	sessionId := in.GetSession().GetId()
	player := in.GetSession().GetPlayerId()

	if err := validateSessionID(sessionId); err != nil {
		s.logger.Error("invalid session id", zap.Error(err))
		return nil, status.Error(codes.Aborted, "invalid session id")
	}

	spanAttributes(span, sessionId, player)

	session, err := s.srv.DeleteSession(ctx, sessionId)

	if err != nil {
		s.logger.Error("failed to delete session",
			zap.String("player", player),
			zap.String("session-id", sessionId),
			zap.Error(err))

		return nil, status.Error(codes.Internal, "failed to delete session")
	}

	return &pb.EndSessionResponse{Session: &pb.Session{
		Id:       session.Id,
		PlayerId: session.PlayerId,
		Wins:     int32(session.Wins),
		Losses:   int32(session.Losses),
		Draws:    int32(session.Draws),
	}}, nil
}
