package v1

import (
	"context"
	"math/rand"
	"strings"

	"github.com/ksaritek/paper-rock-scissors/internal/observability"
	pb "github.com/ksaritek/paper-rock-scissors/proto/gen/go/game/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var looneyTunes = [10]string{"Bugs Bunny", "Daffy Duck", "Porky Pig", "Tweety Bird", "Elmer Fudd", "Yosemite Sam", "Foghorn Leghorn", "Sylvester the Cat", "Road Runner", "Wile E. Coyote"}

func (s *controller) StartSession(ctx context.Context, in *pb.StartSessionRequest) (*pb.StartSessionResponse, error) {
	ctx, span := observability.Tracer().Start(ctx, "start-session")
	defer span.End()

	playerId := in.GetPlayerId()

	if strings.TrimSpace(playerId) == "" {
		playerId = randomName()
	}

	session, err := s.srv.StartSession(ctx, playerId)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to start session")
	}

	spanAttributes(span, session.Id, playerId)

	return &pb.StartSessionResponse{Session: &pb.Session{
		Id:       session.Id,
		PlayerId: playerId,
		Wins:     0,
		Losses:   0,
		Draws:    0,
	}}, nil
}

func randomName() string {
	return looneyTunes[rand.Intn(len(looneyTunes))]
}
