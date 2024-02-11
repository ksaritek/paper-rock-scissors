package v1

import (
	"context"
	"strings"

	"github.com/ksaritek/paper-rock-scissors/internal/domain/model"
	"github.com/ksaritek/paper-rock-scissors/internal/domain/service"
	pb "github.com/ksaritek/paper-rock-scissors/proto/gen/go/game/v1"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type controller struct {
	pb.UnimplementedGameServiceServer
	srv    service.Session
	logger *zap.Logger
}

func NewController(ctx context.Context, l *zap.Logger, srv service.Session) *controller {
	return &controller{
		srv:    srv,
		logger: l.With(zap.String("version", "v1"), zap.String("controller", "game")),
	}
}

func spanAttributes(span trace.Span, sessionId string, playerId string) {
	span.SetAttributes(attribute.String("player_id", playerId),
		attribute.String("session_id", sessionId))
}

func validateSessionID(sessionID string) error {
	if strings.TrimSpace(sessionID) == "" {
		return status.Errorf(codes.InvalidArgument, "session id is missing")
	}

	return nil
}

func validateMove(move pb.Move) error {
	if move == pb.Move_MOVE_UNSPECIFIED {
		return status.Errorf(codes.InvalidArgument, "player move is missing")
	}

	return nil
}

func toModelChoice(move pb.Move) model.Choice {
	switch move {
	case pb.Move_MOVE_ROCK:
		return model.ROCK
	case pb.Move_MOVE_PAPER:
		return model.PAPER
	case pb.Move_MOVE_SCISSORS:
		return model.SCISSORS
	default:
		return model.INVALID_CHOICE
	}
}
func fromModelChoice(choice model.Choice) pb.Move {
	switch choice {
	case model.ROCK:
		return pb.Move_MOVE_ROCK
	case model.PAPER:
		return pb.Move_MOVE_PAPER
	case model.SCISSORS:
		return pb.Move_MOVE_SCISSORS
	default:
		return pb.Move_MOVE_UNSPECIFIED
	}
}
