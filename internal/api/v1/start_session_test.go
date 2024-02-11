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

func TestStartSession(t *testing.T) {
	ctx := context.Background()

	repo := memory.NewMemoryStore(ctx, logger.NewLogger("debug"), time.Hour)
	l := logger.NewLogger("debug")
	service := service.NewSessionService(context.Background(), l, func() string { return "123" }, func() model.Choice { return model.PAPER }, repo)
	controller := &controller{
		logger: l,
		srv:    service,
	}

	req := &pb.StartSessionRequest{}
	resp, err := controller.StartSession(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, resp.Session)
	assert.Equal(t, "123", resp.Session.Id)
	assert.NotEmpty(t, resp.Session.PlayerId)
	assert.Equal(t, int32(0), resp.Session.Wins)
	assert.Equal(t, int32(0), resp.Session.Losses)
	assert.Equal(t, int32(0), resp.Session.Draws)

	playerId := "player1"
	req = &pb.StartSessionRequest{
		PlayerId: &playerId,
	}
	resp, err = controller.StartSession(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, resp.Session)
	assert.Equal(t, "123", resp.Session.Id)
	assert.Equal(t, "player1", resp.Session.PlayerId)
	assert.Equal(t, int32(0), resp.Session.Wins)
	assert.Equal(t, int32(0), resp.Session.Losses)
	assert.Equal(t, int32(0), resp.Session.Draws)
}
