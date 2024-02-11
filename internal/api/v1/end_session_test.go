package v1

import (
	"context"
	"testing"
	"time"

	"github.com/ksaritek/paper-rock-scissors/internal/domain/model"
	"github.com/ksaritek/paper-rock-scissors/internal/domain/repository/memory"
	"github.com/ksaritek/paper-rock-scissors/internal/domain/service"
	"github.com/ksaritek/paper-rock-scissors/internal/logger"
	pb "github.com/ksaritek/paper-rock-scissors/proto/gen/go/game/v1"
	"github.com/stretchr/testify/assert"
)

func TestController_EndSession(t *testing.T) {
	ctx := context.Background()

	repo := memory.NewMemoryStore(ctx, logger.NewLogger("debug"), time.Hour)
	l := logger.NewLogger("debug")
	service := service.NewSessionService(context.Background(), l, func() string { return "123456" }, func() model.Choice { return model.PAPER }, repo)
	ctrl := &controller{
		logger: l,
		srv:    service,
	}
	playerId := "player1"
	ctrl.StartSession(ctx, &pb.StartSessionRequest{
		PlayerId: &playerId,
	})

	req := &pb.EndSessionRequest{
		Session: &pb.Session{
			Id:       "123456",
			PlayerId: "player1",
		},
	}

	resp, err := ctrl.EndSession(ctx, req)

	assert.NoError(t, err)
	assert.Equal(t, "123456", resp.Session.Id)
	assert.Equal(t, "player1", resp.Session.PlayerId)
	assert.Equal(t, int32(0), resp.Session.Wins)
	assert.Equal(t, int32(0), resp.Session.Losses)
	assert.Equal(t, int32(0), resp.Session.Draws)
}
