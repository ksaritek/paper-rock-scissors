package v1

import (
	"context"
	"math/rand"

	"github.com/ksaritek/paper-rock-scissors/internal/domain/model"
	"github.com/ksaritek/paper-rock-scissors/internal/observability"
	pb "github.com/ksaritek/paper-rock-scissors/proto/gen/go/game/v1"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *controller) PlayGame(ctx context.Context, in *pb.PlayGameRequest) (*pb.PlayGameResponse, error) {
	ctx, span := observability.Tracer().Start(ctx, "play-game")
	defer span.End()

	playerMove := in.GetPlayerMove()
	sessionId := in.GetSessionId()

	if err := validateSessionID(sessionId); err != nil {
		s.logger.Error("invalid session id", zap.Error(err))
		return nil, err
	}

	if err := validateMove(playerMove); err != nil {
		s.logger.Error("invalid player move", zap.Error(err))
		return nil, err
	}

	span.SetAttributes(attribute.String("player_move", playerMove.String()))
	span.SetAttributes(attribute.String("session_id", sessionId))

	choice := toModelChoice(playerMove)
	move := model.Move{
		Id:     sessionId,
		Choice: choice,
	}

	session, computerMove, err := s.srv.Move(ctx, move, getRandomMove)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to play game")
	}

	return &pb.PlayGameResponse{
		ComputerMove: fromModelChoice(computerMove),
		Session: &pb.Session{
			Id:       session.Id,
			PlayerId: session.PlayerId,
			Wins:     int32(session.Wins),
			Losses:   int32(session.Losses),
			Draws:    int32(session.Draws),
		},
	}, nil
}

func getRandomMove() model.Choice {
	moves := []model.Choice{model.ROCK, model.PAPER, model.SCISSORS}
	randomIndex := rand.Intn(len(moves))
	return moves[randomIndex]
}
